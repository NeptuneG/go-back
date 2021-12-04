package server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/NeptuneG/go-back/db/types"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type PubSub struct {
	store *db.Store
}

type LiveEvent struct {
	LiveHouseSlug   types.NullString `json:"live_house_slug"`
	Title           string           `json:"title"`
	Url             string           `json:"url"`
	Description     types.NullString `json:"description"`
	PriceInfo       types.NullString `json:"price_info"`
	StageOneOpenAt  types.NullTime   `json:"stage_one_open_at"`
	StageOneStartAt time.Time        `json:"stage_one_start_at"`
	StageTwoOpenAt  types.NullTime   `json:"stage_two_open_at"`
	StageTwoStartAt types.NullTime   `json:"stage_two_start_at"`
}

func (pubsub *PubSub) StartPubSub() {
	ctx := context.Background()
	redisOptions := &redis.Options{
		Addr: "redis:6379",
	}
	redisClient := redis.NewClient(redisOptions)
	subcriber := redisClient.Subscribe(ctx, "live_events")
	var live_events []LiveEvent

	go func() {
		for {
			msg, err := subcriber.ReceiveMessage(ctx)
			if err != nil {
				log.Fatal(err)
				continue
			}
			if err := json.Unmarshal([]byte(msg.Payload), &live_events); err != nil {
				log.Fatal(err)
				continue
			}
			err = pubsub.createLiveEvent(context.Background(), live_events)
			if err != nil {
				log.Fatal(err)
				continue
			} else {
				log.Println("created live events. count: ", len(live_events))
			}
		}
	}()
}

func (pubsub *PubSub) createLiveEvent(ctx context.Context, live_events []LiveEvent) error {
	live_house_id_and_slugs, err := pubsub.store.GetAllLiveHousesIdAndSlugs(ctx)
	if err != nil {
		return err
	}

	live_house_id_by_slug := make(map[types.NullString]uuid.UUID)
	for _, live_house_id_and_slug := range live_house_id_and_slugs {
		live_house_id_by_slug[live_house_id_and_slug.Slug] = live_house_id_and_slug.ID
	}
	for _, live_event := range live_events {
		create_params := db.CreateLiveEventParams{
			LiveHouseID:     live_house_id_by_slug[live_event.LiveHouseSlug],
			Title:           live_event.Title,
			Url:             live_event.Url,
			Description:     live_event.Description,
			PriceInfo:       live_event.PriceInfo,
			StageOneOpenAt:  live_event.StageOneOpenAt,
			StageOneStartAt: live_event.StageOneStartAt,
			StageTwoOpenAt:  live_event.StageTwoOpenAt,
			StageTwoStartAt: live_event.StageTwoStartAt,
		}
		if _, err := pubsub.store.CreateLiveEvent(ctx, create_params); err != nil {
			return err
		}
	}
	return nil
}
