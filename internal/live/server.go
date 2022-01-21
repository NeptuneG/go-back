package live

import (
	"context"
	"errors"
	"fmt"
	"time"

	proto "github.com/NeptuneG/go-back/api/proto/live"
	db "github.com/NeptuneG/go-back/internal/live/db/sqlc"
	"github.com/NeptuneG/go-back/internal/pkg/cache"
	"github.com/NeptuneG/go-back/internal/pkg/db/types"
	"github.com/NeptuneG/go-back/internal/pkg/log"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	count = 0
)

type LiveService struct {
	proto.UnimplementedLiveServiceServer
	store *db.Store
}

func New(ctx context.Context) *LiveService {
	return &LiveService{
		store: db.NewStore(),
	}
}

func (s *LiveService) Close() {
	if err := s.store.Close(); err != nil {
		log.Error("failed to close database connection", log.Field.Error(err))
		panic(err)
	}
}

func (s *LiveService) CreateLiveHouse(ctx context.Context, req *proto.CreateLiveHouseRequest) (*proto.CreateLiveHouseResponse, error) {
	liveHouse, err := s.store.CreateLiveHouse(ctx, db.CreateLiveHouseParams{
		Name:    req.Name,
		Address: types.NewNullString(req.Address),
		Slug:    types.NewNullString(req.Slug),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create live house")
	}
	return &proto.CreateLiveHouseResponse{
		LiveHouse: &proto.LiveHouse{
			Id:      liveHouse.ID.String(),
			Name:    liveHouse.Name,
			Address: liveHouse.Address.String,
			Slug:    liveHouse.Slug.String,
		},
	}, nil
}

func (s *LiveService) CreateLiveEvent(ctx context.Context, req *proto.CreateLiveEventRequest) (*proto.CreateLiveEventResponse, error) {
	liveHouse, err := s.getLiveHouseBySlug(ctx, req.LiveHouseSlug)
	if err != nil {
		return nil, status.Error(codes.NotFound, "live house not found")
	}
	liveEvent, err := s.store.CreateLiveEvent(ctx, db.CreateLiveEventParams{
		LiveHouseID:     liveHouse.ID,
		Title:           req.Title,
		Url:             req.Url,
		Description:     types.NewNullString(req.Description),
		PriceInfo:       types.NewNullString(req.PriceInfo),
		StageOneOpenAt:  types.NewNullTime(req.StageOneOpenAt.AsTime()),
		StageOneStartAt: req.StageOneStartAt.AsTime(),
		StageTwoOpenAt:  types.NewNullTime(req.StageTwoOpenAt.AsTime()),
		StageTwoStartAt: types.NewNullTime(req.StageTwoStartAt.AsTime()),
		Seats:           req.Seats,
		AvailableSeats:  req.AvailableSeats,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create live event")
	}
	return &proto.CreateLiveEventResponse{
		LiveEvent: &proto.LiveEvent{
			Id: liveEvent.ID.String(),
			LiveHouse: &proto.LiveHouse{
				Id:   liveHouse.ID.String(),
				Name: liveHouse.Name,
				Slug: liveHouse.Slug.String,
			},
			Title:           liveEvent.Title,
			Url:             liveEvent.Url,
			Description:     liveEvent.Description.String,
			PriceInfo:       liveEvent.PriceInfo.String,
			StageOneOpenAt:  timestamppb.New(liveEvent.StageOneOpenAt.Time),
			StageOneStartAt: timestamppb.New(liveEvent.StageOneStartAt),
			StageTwoOpenAt:  timestamppb.New(liveEvent.StageTwoOpenAt.Time),
			StageTwoStartAt: timestamppb.New(liveEvent.StageTwoStartAt.Time),
			Seats:           liveEvent.Seats,
			AvailableSeats:  liveEvent.AvailableSeats,
		},
	}, nil
}

func (s *LiveService) getLiveHouseBySlug(ctx context.Context, liveHouseSlug string) (*db.LiveHouse, error) {
	if liveHouseSlug == "" {
		return nil, errors.New("liveHouseSlug is empty")
	}

	var liveHouse db.LiveHouse
	if err := cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("live-house:%s", liveHouseSlug),
		Value: &liveHouse,
		TTL:   5 * time.Minute,
		Do: func(item *cache.Item) (interface{}, error) {
			liveHouse, err := s.store.GetLiveHouseBySlug(ctx, types.NewNullString(liveHouseSlug))
			if err != nil {
				log.Error("failed to get live house by slug", log.Field.Error(err), log.Field.String("liveHouseSlug", liveHouseSlug))
				return nil, err
			}
			return liveHouse, nil
		},
	}); err != nil {
		return nil, err
	}
	return &liveHouse, nil
}

func (s *LiveService) ListLiveHouses(ctx context.Context, req *proto.ListLiveHousesRequest) (*proto.ListLiveHousesResponse, error) {
	var liveHouses []db.GetAllLiveHousesDefaultRow
	if err := cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   "live-houses",
		Value: &liveHouses,
		TTL:   5 * time.Minute,
		Do: func(item *cache.Item) (interface{}, error) {
			liveHouses, err := s.store.GetAllLiveHousesDefault(ctx)
			if err != nil {
				log.Error("failed to get all live houses", log.Field.Error(err))
				return nil, err
			}
			return liveHouses, nil
		},
	}); err != nil {
		return nil, status.Error(codes.Internal, "failed to list live houses")
	}
	liveHousesResp := make([]*proto.LiveHouse, 0, len(liveHouses))
	for _, liveHouse := range liveHouses {
		liveHousesResp = append(liveHousesResp, &proto.LiveHouse{
			Id:      liveHouse.ID.String(),
			Name:    liveHouse.Name,
			Address: liveHouse.Address.String,
			Slug:    liveHouse.Slug.String,
		})
	}
	return &proto.ListLiveHousesResponse{
		LiveHouses: liveHousesResp,
	}, nil
}

func (s *LiveService) ListLiveEvents(ctx context.Context, req *proto.ListLiveEventsRequest) (*proto.ListLiveEventsResponse, error) {
	var liveEvents []db.GetAllLiveEventsRow
	if err := cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   "live-events",
		Value: &liveEvents,
		TTL:   5 * time.Minute,
		Do: func(item *cache.Item) (interface{}, error) {
			liveEvents, err := s.store.GetAllLiveEvents(ctx)
			if err != nil {
				log.Error("failed to get all live events", log.Field.Error(err))
				return nil, err
			}
			return liveEvents, nil
		},
	}); err != nil {
		return nil, status.Error(codes.Internal, "failed to list live events")
	}
	liveEventsResp := make([]*proto.LiveEvent, 0, len(liveEvents))
	for _, liveEvent := range liveEvents {
		liveHouse, err := s.getLiveHouseBySlug(ctx, liveEvent.LiveHouseSlug.String)
		if err != nil {
			return nil, status.Error(codes.NotFound, "live house not found")
		}
		liveEventsResp = append(liveEventsResp, &proto.LiveEvent{
			Id: liveEvent.ID.String(),
			LiveHouse: &proto.LiveHouse{
				Id:      liveHouse.ID.String(),
				Name:    liveHouse.Name,
				Address: liveHouse.Address.String,
				Slug:    liveHouse.Slug.String,
			},
			Title:           liveEvent.Title,
			Url:             liveEvent.Url,
			Description:     liveEvent.Description.String,
			PriceInfo:       liveEvent.PriceInfo.String,
			StageOneOpenAt:  timestamppb.New(liveEvent.StageOneOpenAt.Time),
			StageOneStartAt: timestamppb.New(liveEvent.StageOneStartAt),
			StageTwoOpenAt:  timestamppb.New(liveEvent.StageTwoOpenAt.Time),
			StageTwoStartAt: timestamppb.New(liveEvent.StageTwoStartAt.Time),
			Seats:           liveEvent.Seats,
			AvailableSeats:  liveEvent.AvailableSeats,
		})
	}
	return &proto.ListLiveEventsResponse{
		LiveEvents: liveEventsResp,
	}, nil
}

func (s *LiveService) GetLiveEvent(ctx context.Context, req *proto.GetLiveEventRequest) (*proto.GetLiveEventResponse, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to parse id")
	}
	// TODO: BloomFilter to avoid hitting DB by unexisting id
	var liveEvent db.GetLiveEventByIdRow
	if err := cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("live-event:%s", req.Id),
		Value: &liveEvent,
		TTL:   5 * time.Minute,
		Do: func(item *cache.Item) (interface{}, error) {
			liveEvent, err := s.store.GetLiveEventById(ctx, uuid)
			if err != nil {
				log.Error("failed to get live event by id", log.Field.Error(err), log.Field.String("id", req.Id))
				return nil, err
			}
			return liveEvent, nil
		},
	}); err != nil {
		return nil, status.Error(codes.NotFound, "live event not found")
	}
	return &proto.GetLiveEventResponse{
		LiveEvent: &proto.LiveEvent{
			Id: liveEvent.ID.String(),
			LiveHouse: &proto.LiveHouse{
				Id:   liveEvent.LiveHouseID.String(),
				Name: liveEvent.LiveHouseName,
				Slug: liveEvent.LiveHouseSlug.String,
			},
			Title:           liveEvent.Title,
			Url:             liveEvent.Url,
			Description:     liveEvent.Description.String,
			PriceInfo:       liveEvent.PriceInfo.String,
			StageOneOpenAt:  timestamppb.New(liveEvent.StageOneOpenAt.Time),
			StageOneStartAt: timestamppb.New(liveEvent.StageOneStartAt),
			StageTwoOpenAt:  timestamppb.New(liveEvent.StageTwoOpenAt.Time),
			StageTwoStartAt: timestamppb.New(liveEvent.StageTwoStartAt.Time),
			Seats:           liveEvent.Seats,
			AvailableSeats:  liveEvent.AvailableSeats,
		},
	}, nil
}

func (s *LiveService) ReserveSeat(ctx context.Context, req *proto.ReserveSeatRequest) (*proto.ReserveSeatResponse, error) {
	// force a retry
	count++
	log.Debug("mock failure for retry", log.Field.Int("count", count))
	if count%3 != 0 {
		return nil, status.Error(codes.Internal, "just failed")
	}

	// mock delay
	if false {
		log.Debug("mock delay")
		time.Sleep(10 * time.Second)
		log.Debug("mock delay done")
	}

	liveEventID, err := uuid.Parse(req.LiveEventId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to parse live event id")
	}
	if liveEvent, err := s.store.ReserveSeatTx(ctx, liveEventID); err != nil {
		return nil, status.Error(codes.Internal, "failed to reserve seat")
	} else {
		log.Info("reserved seat",
			log.Field.String("live_event_id", liveEvent.ID.String()),
			log.Field.Int32("live_event_available_seats", liveEvent.AvailableSeats),
		)
		return &proto.ReserveSeatResponse{
			LiveEvent: &proto.LiveEvent{
				Id: liveEvent.ID.String(),
				LiveHouse: &proto.LiveHouse{
					Id:   liveEvent.LiveHouseID.String(),
					Name: liveEvent.LiveHouseName,
					Slug: liveEvent.LiveHouseSlug.String,
				},
				Title:           liveEvent.Title,
				Url:             liveEvent.Url,
				Description:     liveEvent.Description.String,
				PriceInfo:       liveEvent.PriceInfo.String,
				StageOneOpenAt:  timestamppb.New(liveEvent.StageOneOpenAt.Time),
				StageOneStartAt: timestamppb.New(liveEvent.StageOneStartAt),
				StageTwoOpenAt:  timestamppb.New(liveEvent.StageTwoOpenAt.Time),
				StageTwoStartAt: timestamppb.New(liveEvent.StageTwoStartAt.Time),
				Seats:           liveEvent.Seats,
				AvailableSeats:  liveEvent.AvailableSeats,
			},
		}, nil
	}
}

func (s *LiveService) RollbackSeatReservation(ctx context.Context, req *proto.RollbackSeatReservationRequest) (*emptypb.Empty, error) {
	liveEventID, err := uuid.Parse(req.LiveEventId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to parse live event id")
	}
	if liveEvent, err := s.store.RollbackSeatReservationTx(ctx, liveEventID); err != nil {
		return nil, status.Error(codes.Internal, "failed to rollback seat reservation")
	} else {
		log.Info("rollbacked seat reservation",
			log.Field.String("live_event_id", liveEvent.ID.String()),
			log.Field.Int32("live_event_available_seats", liveEvent.AvailableSeats),
		)
		return &emptypb.Empty{}, nil
	}
}
