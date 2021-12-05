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

type ScrapedConsumer struct {
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

func (scraped_consumer *ScrapedConsumer) Start() {
	ctx := context.Background()
	redisOptions := &redis.Options{
		Addr: "redis:6379",
	}
	redisClient := redis.NewClient(redisOptions)

	go func() {
		for {
			cmd := redisClient.BRPop(ctx, 0, "screped_live_events")
			var live_event LiveEvent

			cmd_val, err := cmd.Result()
			if err != nil {
				log.Fatal(err)
				continue
			}
			if err := json.Unmarshal([]byte(cmd_val[1]), &live_event); err != nil {
				log.Fatal(err)
				continue
			}
			if err := scraped_consumer.createLiveEvent(ctx, live_event); err != nil {
				log.Fatal(err)
				continue
			}
		}
	}()
}

func (scraped_consumer *ScrapedConsumer) createLiveEvent(ctx context.Context, live_event LiveEvent) error {
	live_house_id, err := scraped_consumer.getLiveHouseIdBySlug(ctx, live_event.LiveHouseSlug)
	if err != nil {
		return err
	}
	create_params := db.CreateLiveEventParams{
		LiveHouseID:     live_house_id,
		Title:           live_event.Title,
		Url:             live_event.Url,
		Description:     live_event.Description,
		PriceInfo:       live_event.PriceInfo,
		StageOneOpenAt:  live_event.StageOneOpenAt,
		StageOneStartAt: live_event.StageOneStartAt,
		StageTwoOpenAt:  live_event.StageTwoOpenAt,
		StageTwoStartAt: live_event.StageTwoStartAt,
	}
	if _, err := scraped_consumer.store.CreateLiveEvent(ctx, create_params); err != nil {
		return err
	}
	return nil
}

var live_house_id_by_slug map[types.NullString]uuid.UUID

func (scraped_consumer *ScrapedConsumer) getLiveHouseIdBySlug(ctx context.Context, live_house_slug types.NullString) (uuid.UUID, error) {
	if live_house_id_by_slug == nil {
		live_house_id_and_slugs, err := scraped_consumer.store.GetAllLiveHousesIdAndSlugs(ctx)
		if err != nil {
			return uuid.Nil, err
		}
		live_house_id_by_slug = make(map[types.NullString]uuid.UUID)
		for _, live_house_id_and_slug := range live_house_id_and_slugs {
			live_house_id_by_slug[live_house_id_and_slug.Slug] = live_house_id_and_slug.ID
		}
	}
	return live_house_id_by_slug[live_house_slug], nil
}
