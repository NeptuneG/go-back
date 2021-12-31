package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NeptuneG/go-back/gen/go/services/live/proto"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
	"github.com/NeptuneG/go-back/pkg/types"
	db "github.com/NeptuneG/go-back/services/live/db/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://dev@db/live_development?sslmode=disable"
)

type LiveService struct {
	proto.UnimplementedLiveServiceServer
	store  *db.Store
	dbConn *sql.DB
}

func New() *LiveService {
	dbConn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("failed to open database connection", logField.Error(err))
		panic(err)
	}
	return &LiveService{
		store: db.NewStore(dbConn),
	}
}

func (s *LiveService) Close() {
	if err := s.store.Close(); err != nil {
		log.Error("failed to close database connection", logField.Error(err))
		panic(err)
	}
	if err := s.dbConn.Close(); err != nil {
		log.Fatal("failed to close database connection", logField.Error(err))
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

var liveHousesBySlug map[string]db.LiveHouse

func (s *LiveService) getLiveHouseBySlug(ctx context.Context, liveHouseSlug string) (*db.LiveHouse, error) {
	if liveHouseSlug == "" {
		return nil, errors.New("liveHouseSlug is empty")
	}

	if liveHousesBySlug == nil {
		liveHousesBySlug = make(map[string]db.LiveHouse)
	}
	if liveHouse, ok := liveHousesBySlug[liveHouseSlug]; ok {
		return &liveHouse, nil
	} else {
		liveHouse, err := s.store.GetLiveHouseBySlug(ctx, types.NewNullString(liveHouseSlug))
		if err != nil {
			return nil, err
		}
		liveHousesBySlug[liveHouseSlug] = liveHouse
		return &liveHouse, nil
	}
}

func (s *LiveService) ListLiveHouses(ctx context.Context, req *proto.ListLiveHousesRequest) (*proto.ListLiveHousesResponse, error) {
	liveHouses, err := s.store.GetAllLiveHousesDefault(ctx)
	if err != nil {
		return nil, err
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
	liveEvents, err := s.store.GetAllLiveEvents(ctx)
	if err != nil {
		return nil, err
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
	liveEvent, err := s.store.GetLiveEventById(ctx, uuid)
	if err != nil {
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
	return s.store.ReserveSeatTx(ctx, req)
}

func (s *LiveService) RollbackSeatReservation(ctx context.Context, req *proto.RollbackSeatReservationRequest) (*emptypb.Empty, error) {
	return s.store.RollbackSeatReservationTx(ctx, req)
}
