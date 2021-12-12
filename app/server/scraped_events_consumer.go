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
	AvaiableSeats   *int32           `json:"avaiable_seats"`
}

func (request *createLiveEventRequest) toCreateLiveEventParams(liveHouseId uuid.UUID) db.CreateLiveEventParams {
	availableSeats := int32(100)
	if request.AvaiableSeats != nil {
		availableSeats = *request.AvaiableSeats
	}
	return db.CreateLiveEventParams{
		LiveHouseID:     liveHouseId,
		Title:           request.Title,
		Url:             request.Url,
		Description:     request.Description,
		PriceInfo:       request.PriceInfo,
		StageOneOpenAt:  request.StageOneOpenAt,
		StageOneStartAt: request.StageOneStartAt,
		StageTwoOpenAt:  request.StageTwoOpenAt,
		StageTwoStartAt: request.StageTwoStartAt,
		AvailableSeats:  availableSeats,
	}
}

func (consumer *ScrapedEventsConsumer) Start() {
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
			if err := consumer.createLiveEvent(ctx, req); err != nil {
				log.Fatal(err)
				continue
			}
		}
	}()
}

func (consumer *ScrapedEventsConsumer) createLiveEvent(ctx context.Context, req createLiveEventRequest) error {
	liveHouseId, err := consumer.getLiveHouseIdBySlug(ctx, req.LiveHouseSlug)
	if err != nil {
		return err
	}
	create_params := req.toCreateLiveEventParams(liveHouseId)
	if _, err := consumer.store.CreateLiveEvent(ctx, create_params); err != nil {
		return err
	}
	return nil
}

var liveHouseIdBySlug map[string]uuid.UUID

func (consumer *ScrapedEventsConsumer) getLiveHouseIdBySlug(ctx context.Context, liveHouseSlug string) (uuid.UUID, error) {
	if liveHouseSlug == "" {
		return uuid.Nil, errors.New("liveHouseSlug is empty")
	}

	if liveHouseIdBySlug == nil {
		liveHouseIdAndSlugs, err := consumer.store.GetAllLiveHousesIdAndSlugs(ctx)
		if err != nil {
			return uuid.Nil, err
		}
		liveHouseIdBySlug = make(map[string]uuid.UUID)
		for _, liveHouseIdAndSlug := range liveHouseIdAndSlugs {
			if liveHouseIdAndSlug.Slug.Valid {
				liveHouseIdBySlug[liveHouseIdAndSlug.Slug.String] = liveHouseIdAndSlug.ID
			}
		}
	}
	liveHouseId, ok := liveHouseIdBySlug[liveHouseSlug]
	if ok {
		return liveHouseId, nil
	} else {
		return uuid.Nil, fmt.Errorf("liveHouseId of liveHouseSlug (%s) is not found", liveHouseSlug)
	}
}
