package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NeptuneG/go-back/gen/go/services/live/proto"
	"github.com/NeptuneG/go-back/pkg/types"
	db "github.com/NeptuneG/go-back/services/live/db/sqlc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type LiveService struct {
	proto.UnimplementedLiveServiceServer
	store *db.Store
}

func New(dbConn *sql.DB) *LiveService {
	return &LiveService{
		store: db.NewStore(dbConn),
	}
}

func (liveService *LiveService) CreateLiveHouse(ctx context.Context, req *proto.CreateLiveHouseRequest) (*proto.CreateLiveHouseResponse, error) {
	liveHouse, err := liveService.store.CreateLiveHouse(ctx, db.CreateLiveHouseParams{
		Name:    req.Name,
		Address: types.NewNullString(req.Address),
		Slug:    types.NewNullString(req.Slug),
	})
	if err != nil {
		return nil, err
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

func (liveService *LiveService) CreateLiveEvent(ctx context.Context, req *proto.CreateLiveEventRequest) (*proto.CreateLiveEventResponse, error) {
	liveHouse, err := liveService.getLiveHouseBySlug(ctx, req.LiveHouseSlug)
	if err != nil {
		return nil, err
	}
	liveEvent, err := liveService.store.CreateLiveEvent(ctx, db.CreateLiveEventParams{
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
		return nil, err
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

func (liveService *LiveService) getLiveHouseBySlug(ctx context.Context, liveHouseSlug string) (*db.LiveHouse, error) {
	if liveHouseSlug == "" {
		return nil, errors.New("liveHouseSlug is empty")
	}

	if liveHousesBySlug == nil {
		liveHousesBySlug = make(map[string]db.LiveHouse)
	}
	if liveHouse, ok := liveHousesBySlug[liveHouseSlug]; ok {
		return &liveHouse, nil
	} else {
		liveHouse, err := liveService.store.GetLiveHouseBySlug(ctx, types.NewNullString(liveHouseSlug))
		if err != nil {
			return nil, err
		}
		liveHousesBySlug[liveHouseSlug] = liveHouse
		return &liveHouse, nil
	}
}

func (liveService *LiveService) ListLiveHouses(ctx context.Context, req *proto.ListLiveHousesRequest) (*proto.ListLiveHousesResponse, error) {
	liveHouses, err := liveService.store.GetAllLiveHousesDefault(ctx)
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

func (liveService *LiveService) ListLiveEvents(ctx context.Context, req *proto.ListLiveEventsRequest) (*proto.ListLiveEventsResponse, error) {
	liveEvents, err := liveService.store.GetAllLiveEvents(ctx)
	if err != nil {
		return nil, err
	}
	liveEventsResp := make([]*proto.LiveEvent, 0, len(liveEvents))
	for _, liveEvent := range liveEvents {
		liveHouse, err := liveService.getLiveHouseBySlug(ctx, liveEvent.LiveHouseSlug.String)
		if err != nil {
			return nil, err
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
