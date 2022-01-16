package consumer

import (
	"context"
	"encoding/json"
	"os"
	"time"

	live "github.com/NeptuneG/go-back/api/proto/live"
	liveSvc "github.com/NeptuneG/go-back/internal/live"
	"github.com/NeptuneG/go-back/internal/pkg/cache"
	"github.com/NeptuneG/go-back/internal/pkg/db/types"
	"github.com/NeptuneG/go-back/internal/pkg/log"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScrapedEventsConsumer struct {
	liveClient  live.LiveServiceClient
	redisClient *redis.Client
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
	msgQueueKey = "screped_live_events"
)

var (
	redisMqHost = os.Getenv("REDIS_MQ_SERVICE_HOST") + ":" + os.Getenv("REDIS_MQ_SERVICE_PORT")
)

func New() *ScrapedEventsConsumer {
	redisClient := make(chan *redis.Client)
	liveClient := make(chan live.LiveServiceClient)

	go func() {
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		}

		client, err := liveSvc.NewClient(opts...)
		if err != nil {
			log.Error("failed to create live client", log.Field.Error(err))
			panic(err)
		}

		liveClient <- client
	}()

	go func() {
		redisOptions := &redis.Options{Addr: redisMqHost}
		redisClient <- redis.NewClient(redisOptions)
	}()

	return &ScrapedEventsConsumer{
		liveClient:  <-liveClient,
		redisClient: <-redisClient,
	}
}

func (consumer *ScrapedEventsConsumer) Close() {
}

func (consumer *ScrapedEventsConsumer) Start() {
	ctx := context.Background()
	for {
		raw := consumer.redisClient.BRPop(ctx, 0, msgQueueKey)
		var reqMsg createLiveEventMessage

		message, err := raw.Result()
		if err != nil {
			log.Error("failed to pop message from queue", log.Field.Error(err))
			continue
		}
		if err := json.Unmarshal([]byte(message[1]), &reqMsg); err != nil {
			log.Error("failed to unmarshal message", log.Field.Error(err))
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
			log.Error("failed to create live event", log.Field.Error(err))
			continue
		}
		if err := cache.Delete(ctx, "live-events"); err != nil {
			log.Error("failed to delete cache", log.Field.Error(err))
		}
	}
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
