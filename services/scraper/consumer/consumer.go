package consumer

import (
	"context"
	"encoding/json"
	"time"

	live "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	"github.com/NeptuneG/go-back/pkg/types"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScrapedEventsConsumer struct {
	liveClient live.LiveServiceClient
	logger     *zap.Logger
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
	Seats           *int32           `json:"seats"`
	AvailableSeats  *int32           `json:"available_seats"`
}

const (
	redisAddr   = "redis-mq:6379"
	msgQueueKey = "screped_live_events"
)

func New(liveClient live.LiveServiceClient, logger *zap.Logger) *ScrapedEventsConsumer {
	return &ScrapedEventsConsumer{
		liveClient: liveClient,
		logger:     logger,
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
				consumer.logger.Error("failed to pop message from queue", zap.Error(err))
				continue
			}
			if err := json.Unmarshal([]byte(message[1]), &reqMsg); err != nil {
				consumer.logger.Error("failed to unmarshal message", zap.Error(err))
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
				Seats:           seats(reqMsg.Seats),
				AvailableSeats:  availableSeats(reqMsg.AvailableSeats),
			}); err != nil {
				consumer.logger.Error("failed to create live event", zap.Error(err))
				continue
			}
		}
	}()
}

func seats(raw *int32) int32 {
	if raw == nil {
		return 100
	} else {
		return *raw
	}
}

func availableSeats(raw *int32) int32 {
	if raw == nil {
		return 100
	} else {
		return *raw
	}
}
