package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/NeptuneG/go-back/db/types"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type ScrapedEventsConsumer struct {
	store *db.Store
}

type createLiveEventRequest struct {
	LiveHouseSlug   string           `json:"live_house_slug"`
	Title           string           `json:"title" binding:"required"`
	Url             string           `json:"url" binding:"required"`
	Description     types.NullString `json:"description"`
	PriceInfo       types.NullString `json:"price_info"`
	StageOneOpenAt  types.NullTime   `json:"stage_one_open_at"`
	StageOneStartAt time.Time        `json:"stage_one_start_at" binding:"required"`
	StageTwoOpenAt  types.NullTime   `json:"stage_two_open_at"`
	StageTwoStartAt types.NullTime   `json:"stage_two_start_at"`
	AvaiableSeats   types.NullInt32  `json:"avaiable_seats"`
}

func (scraped_events_consumer *ScrapedEventsConsumer) Start() {
	ctx := context.Background()
	redisOptions := &redis.Options{
		Addr: "redis:6379",
	}
	redisClient := redis.NewClient(redisOptions)

	go func() {
		for {
			cmd := redisClient.BRPop(ctx, 0, "screped_live_events")
			var req createLiveEventRequest

			cmd_val, err := cmd.Result()
			if err != nil {
				log.Fatal(err)
				continue
			}
			if err := json.Unmarshal([]byte(cmd_val[1]), &req); err != nil {
				log.Fatal(err)
				continue
			}
			if err := scraped_events_consumer.createLiveEvent(ctx, req); err != nil {
				log.Fatal(err)
				continue
			}
		}
	}()
}

func (scraped_events_consumer *ScrapedEventsConsumer) createLiveEvent(ctx context.Context, req createLiveEventRequest) error {
	live_house_id, err := scraped_events_consumer.getLiveHouseIdBySlug(ctx, req.LiveHouseSlug)
	if err != nil {
		return err
	}
	create_params := db.CreateLiveEventParams{
		LiveHouseID:     live_house_id,
		Title:           req.Title,
		Url:             req.Url,
		Description:     req.Description,
		PriceInfo:       req.PriceInfo,
		StageOneOpenAt:  req.StageOneOpenAt,
		StageOneStartAt: req.StageOneStartAt,
		StageTwoOpenAt:  req.StageTwoOpenAt,
		StageTwoStartAt: req.StageTwoStartAt,
		AvailableSeats:  req.AvaiableSeats,
	}
	if _, err := scraped_events_consumer.store.CreateLiveEvent(ctx, create_params); err != nil {
		return err
	}
	return nil
}

var live_house_id_by_slug map[string]uuid.UUID

func (scraped_events_consumer *ScrapedEventsConsumer) getLiveHouseIdBySlug(ctx context.Context, live_house_slug string) (uuid.UUID, error) {
	if live_house_slug == "" {
		return uuid.Nil, errors.New("live_house_slug is empty")
	}

	if live_house_id_by_slug == nil {
		live_house_id_and_slugs, err := scraped_events_consumer.store.GetAllLiveHousesIdAndSlugs(ctx)
		if err != nil {
			return uuid.Nil, err
		}
		live_house_id_by_slug = make(map[string]uuid.UUID)
		for _, live_house_id_and_slug := range live_house_id_and_slugs {
			if live_house_id_and_slug.Slug.Valid {
				live_house_id_by_slug[live_house_id_and_slug.Slug.String] = live_house_id_and_slug.ID
			}
		}
	}
	live_house_id, ok := live_house_id_by_slug[live_house_slug]
	if ok {
		return live_house_id, nil
	} else {
		return uuid.Nil, fmt.Errorf("live_house_id of live_house_slug (%s) is not found", live_house_slug)
	}
}
