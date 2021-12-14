package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NeptuneG/go-back/pkg/types"
	db "github.com/NeptuneG/go-back/services/live/db/sqlc"
	"github.com/NeptuneG/go-back/services/live/proto"
	"github.com/google/uuid"
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

func (liveService *LiveService) CreateLiveEvent(ctx context.Context, req *proto.CreateLiveEventRequest) (*proto.CreateLiveEventResponse, error) {
	liveHouseId, err := liveService.getLiveHouseIdBySlug(ctx, req.LiveHouseSlug)
	if err != nil {
		return nil, err
	}
	liveEvent, err := liveService.store.CreateLiveEvent(ctx, db.CreateLiveEventParams{
		LiveHouseID:     liveHouseId,
		Title:           req.Title,
		Url:             req.Url,
		Description:     types.NewNullString(req.Description),
		PriceInfo:       types.NewNullString(req.PriceInfo),
		StageOneOpenAt:  types.NewNullTime(req.StageOneOpenAt.AsTime()),
		StageOneStartAt: req.StageOneStartAt.AsTime(),
		StageTwoOpenAt:  types.NewNullTime(req.StageTwoOpenAt.AsTime()),
		StageTwoStartAt: types.NewNullTime(req.StageTwoStartAt.AsTime()),
		AvailableSeats:  req.AvailableSeats,
	})
	if err != nil {
		return nil, err
	}
	return &proto.CreateLiveEventResponse{
		LiveEvent: &proto.LiveEvent{
			Id:              liveEvent.ID.String(),
			LiveHouseId:     liveEvent.LiveHouseID.String(),
			Title:           liveEvent.Title,
			Url:             liveEvent.Url,
			Description:     liveEvent.Description.String,
			PriceInfo:       liveEvent.PriceInfo.String,
			StageOneOpenAt:  timestamppb.New(liveEvent.StageOneOpenAt.Time),
			StageOneStartAt: timestamppb.New(liveEvent.StageOneStartAt),
			StageTwoOpenAt:  timestamppb.New(liveEvent.StageTwoOpenAt.Time),
			StageTwoStartAt: timestamppb.New(liveEvent.StageTwoStartAt.Time),
			AvailableSeats:  liveEvent.AvailableSeats,
		},
	}, nil
}

var liveHouseIdBySlug map[string]uuid.UUID

func (liveService *LiveService) getLiveHouseIdBySlug(ctx context.Context, liveHouseSlug string) (uuid.UUID, error) {
	if liveHouseSlug == "" {
		return uuid.Nil, errors.New("liveHouseSlug is empty")
	}

	if liveHouseIdBySlug == nil {
		liveHouseIdAndSlugs, err := liveService.store.GetAllLiveHousesIdAndSlugs(ctx)
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
