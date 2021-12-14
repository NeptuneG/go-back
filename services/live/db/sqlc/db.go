// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createLiveEventStmt, err = db.PrepareContext(ctx, createLiveEvent); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLiveEvent: %w", err)
	}
	if q.createLiveHouseStmt, err = db.PrepareContext(ctx, createLiveHouse); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLiveHouse: %w", err)
	}
	if q.getAllLiveEventsStmt, err = db.PrepareContext(ctx, getAllLiveEvents); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllLiveEvents: %w", err)
	}
	if q.getAllLiveEventsByLiveHouseSlugStmt, err = db.PrepareContext(ctx, getAllLiveEventsByLiveHouseSlug); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllLiveEventsByLiveHouseSlug: %w", err)
	}
	if q.getAllLiveHousesStmt, err = db.PrepareContext(ctx, getAllLiveHouses); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllLiveHouses: %w", err)
	}
	if q.getAllLiveHousesDefaultStmt, err = db.PrepareContext(ctx, getAllLiveHousesDefault); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllLiveHousesDefault: %w", err)
	}
	if q.getAllLiveHousesIdAndSlugsStmt, err = db.PrepareContext(ctx, getAllLiveHousesIdAndSlugs); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllLiveHousesIdAndSlugs: %w", err)
	}
	if q.getLiveEventByIdStmt, err = db.PrepareContext(ctx, getLiveEventById); err != nil {
		return nil, fmt.Errorf("error preparing query GetLiveEventById: %w", err)
	}
	if q.getLiveEventsByLiveHouseStmt, err = db.PrepareContext(ctx, getLiveEventsByLiveHouse); err != nil {
		return nil, fmt.Errorf("error preparing query GetLiveEventsByLiveHouse: %w", err)
	}
	if q.getLiveHouseByIdStmt, err = db.PrepareContext(ctx, getLiveHouseById); err != nil {
		return nil, fmt.Errorf("error preparing query GetLiveHouseById: %w", err)
	}
	if q.updateLiveEventAvailableSeatsByIdStmt, err = db.PrepareContext(ctx, updateLiveEventAvailableSeatsById); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateLiveEventAvailableSeatsById: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createLiveEventStmt != nil {
		if cerr := q.createLiveEventStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLiveEventStmt: %w", cerr)
		}
	}
	if q.createLiveHouseStmt != nil {
		if cerr := q.createLiveHouseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLiveHouseStmt: %w", cerr)
		}
	}
	if q.getAllLiveEventsStmt != nil {
		if cerr := q.getAllLiveEventsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllLiveEventsStmt: %w", cerr)
		}
	}
	if q.getAllLiveEventsByLiveHouseSlugStmt != nil {
		if cerr := q.getAllLiveEventsByLiveHouseSlugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllLiveEventsByLiveHouseSlugStmt: %w", cerr)
		}
	}
	if q.getAllLiveHousesStmt != nil {
		if cerr := q.getAllLiveHousesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllLiveHousesStmt: %w", cerr)
		}
	}
	if q.getAllLiveHousesDefaultStmt != nil {
		if cerr := q.getAllLiveHousesDefaultStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllLiveHousesDefaultStmt: %w", cerr)
		}
	}
	if q.getAllLiveHousesIdAndSlugsStmt != nil {
		if cerr := q.getAllLiveHousesIdAndSlugsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllLiveHousesIdAndSlugsStmt: %w", cerr)
		}
	}
	if q.getLiveEventByIdStmt != nil {
		if cerr := q.getLiveEventByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLiveEventByIdStmt: %w", cerr)
		}
	}
	if q.getLiveEventsByLiveHouseStmt != nil {
		if cerr := q.getLiveEventsByLiveHouseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLiveEventsByLiveHouseStmt: %w", cerr)
		}
	}
	if q.getLiveHouseByIdStmt != nil {
		if cerr := q.getLiveHouseByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLiveHouseByIdStmt: %w", cerr)
		}
	}
	if q.updateLiveEventAvailableSeatsByIdStmt != nil {
		if cerr := q.updateLiveEventAvailableSeatsByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateLiveEventAvailableSeatsByIdStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                    DBTX
	tx                                    *sql.Tx
	createLiveEventStmt                   *sql.Stmt
	createLiveHouseStmt                   *sql.Stmt
	getAllLiveEventsStmt                  *sql.Stmt
	getAllLiveEventsByLiveHouseSlugStmt   *sql.Stmt
	getAllLiveHousesStmt                  *sql.Stmt
	getAllLiveHousesDefaultStmt           *sql.Stmt
	getAllLiveHousesIdAndSlugsStmt        *sql.Stmt
	getLiveEventByIdStmt                  *sql.Stmt
	getLiveEventsByLiveHouseStmt          *sql.Stmt
	getLiveHouseByIdStmt                  *sql.Stmt
	updateLiveEventAvailableSeatsByIdStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                    tx,
		tx:                                    tx,
		createLiveEventStmt:                   q.createLiveEventStmt,
		createLiveHouseStmt:                   q.createLiveHouseStmt,
		getAllLiveEventsStmt:                  q.getAllLiveEventsStmt,
		getAllLiveEventsByLiveHouseSlugStmt:   q.getAllLiveEventsByLiveHouseSlugStmt,
		getAllLiveHousesStmt:                  q.getAllLiveHousesStmt,
		getAllLiveHousesDefaultStmt:           q.getAllLiveHousesDefaultStmt,
		getAllLiveHousesIdAndSlugsStmt:        q.getAllLiveHousesIdAndSlugsStmt,
		getLiveEventByIdStmt:                  q.getLiveEventByIdStmt,
		getLiveEventsByLiveHouseStmt:          q.getLiveEventsByLiveHouseStmt,
		getLiveHouseByIdStmt:                  q.getLiveHouseByIdStmt,
		updateLiveEventAvailableSeatsByIdStmt: q.updateLiveEventAvailableSeatsByIdStmt,
	}
}
