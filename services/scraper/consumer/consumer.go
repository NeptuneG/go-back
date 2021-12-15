package consumer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/NeptuneG/go-back/pkg/types"
	live "github.com/NeptuneG/go-back/services/live/proto"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScrapedEventsConsumer struct {
	liveClient live.LiveServiceClient
}
type createLiveEventMessage struct {
	LiveHouseSlug   string           `json:"live_house_slug"`
	Title           string           `json:"title" binding:"required"`
	Url             string           `json:"url" binding:"required"`
	Description     types.NullString `json:"description"`
	PriceInfo       types.NullString `json:"price_info"`
	StageOneOpenAt  types.NullTime   `json:"stage_one_open_at"`
	StageOneStartAt time.Time        `json:"stage_one_start_at" binding:"required"`
	StageTwoOpenAt  types.NullTime   `json:"stage_two_open_at"`
	StageTwoStartAt types.NullTime   `json:"stage_two_start_at"`
	AvailableSeats  *int32           `json:"available_seats"`
}

const (
	redisAddr   = "redis:6379"
	msgQueueKey = "screped_live_events"
)

func New(liveClient live.LiveServiceClient) *ScrapedEventsConsumer {
	return &ScrapedEventsConsumer{
		liveClient: liveClient,
	}
}

func (consumer *ScrapedEventsConsumer) Start(ctx context.Context) {
	redisOptions := &redis.Options{
		Addr: redisAddr,
	}
	redisClient := redis.NewClient(redisOptions)

	go func() {
		for {
			raw := redisClient.BRPop(ctx, 0, msgQueueKey)
			var reqMsg createLiveEventMessage

			message, err := raw.Result()
			if err != nil {
				log.Fatal(err)
				continue
			}
			if err := json.Unmarshal([]byte(message[1]), &reqMsg); err != nil {
				log.Fatal(err)
				continue
			}
			if _, err := consumer.liveClient.CreateLiveEvent(ctx, &live.CreateLiveEventRequest{
				LiveHouseSlug:   reqMsg.LiveHouseSlug,
				Title:           reqMsg.Title,
				Url:             reqMsg.Url,
				Description:     reqMsg.Description.String,
				PriceInfo:       reqMsg.PriceInfo.String,
				StageOneOpenAt:  timestamppb.New(reqMsg.StageOneOpenAt.Time),
				StageOneStartAt: timestamppb.New(reqMsg.StageOneStartAt),
				StageTwoOpenAt:  timestamppb.New(reqMsg.StageTwoOpenAt.Time),
				StageTwoStartAt: timestamppb.New(reqMsg.StageTwoStartAt.Time),
				AvailableSeats:  availableSeats(reqMsg.AvailableSeats),
			}); err != nil {
				log.Fatal(err)
				continue
			}
		}
	}()
}

func availableSeats(raw *int32) int32 {
	if raw == nil {
		return 100
	} else {
		return *raw
	}
}
