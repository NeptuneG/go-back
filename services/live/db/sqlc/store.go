package db

import (
	"context"
	"database/sql"

	"github.com/NeptuneG/go-back/gen/go/services/live/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) ReserveSeatTx(ctx context.Context, req *proto.ReserveSeatRequest) (*proto.ReserveSeatResponse, error) {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	queries := New(tx)
	liveEventID, err := uuid.Parse(req.LiveEventId)

	if err != nil {
		return nil, err
	}
	liveEvent, err := queries.GetLiveEventById(ctx, liveEventID)
	if err != nil {
		return nil, err
	}

	if err = queries.UpdateLiveEventAvailableSeatsById(ctx, UpdateLiveEventAvailableSeatsByIdParams{
		ID:             liveEventID,
		AvailableSeats: liveEvent.AvailableSeats - 1,
	}); err != nil {
		return nil, err
	}
	return &proto.ReserveSeatResponse{
		LiveEvent: &proto.LiveEvent{
			Id:              liveEvent.ID.String(),
			LiveHouse:       nil,
			Title:           liveEvent.Title,
			Url:             liveEvent.Url,
			Description:     liveEvent.Description.String,
			PriceInfo:       liveEvent.PriceInfo.String,
			StageOneOpenAt:  timestamppb.New(liveEvent.StageOneOpenAt.Time),
			StageOneStartAt: timestamppb.New(liveEvent.StageOneStartAt),
			StageTwoOpenAt:  timestamppb.New(liveEvent.StageTwoOpenAt.Time),
			StageTwoStartAt: timestamppb.New(liveEvent.StageTwoStartAt.Time),
			Seats:           liveEvent.Seats,
			AvailableSeats:  liveEvent.AvailableSeats - 1,
		},
	}, tx.Commit()
}

func (store *Store) RollbackSeatReservationTx(ctx context.Context, req *proto.RollbackSeatReservationRequest) (*emptypb.Empty, error) {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	queries := New(tx)
	liveEventID, err := uuid.Parse(req.LiveEventId)

	if err != nil {
		return nil, err
	}
	liveEvent, err := queries.GetLiveEventById(ctx, liveEventID)
	if err != nil {
		return nil, err
	}

	if err = queries.UpdateLiveEventAvailableSeatsById(ctx, UpdateLiveEventAvailableSeatsByIdParams{
		ID:             liveEventID,
		AvailableSeats: liveEvent.AvailableSeats + 1,
	}); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, tx.Commit()
}
