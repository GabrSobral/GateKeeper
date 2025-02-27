package repositories

import (
	"context"
	"log/slog"

	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceHandler[TRequest any] interface {
	Handler(ctx context.Context, request TRequest) error
}

type ServiceHandlerRs[TRequest any, TResponse any] interface {
	Handler(ctx context.Context, request TRequest) (TResponse, error)
}

type Params[TRequest any, TService any] struct {
	DbPool  *pgxpool.Pool
	New     func(q *pgstore.Queries) ServiceHandler[TRequest]
	Request TRequest
}

type ParamsRs[TRequest any, TResponse any, TService any] struct {
	DbPool  *pgxpool.Pool
	New     func(q *pgstore.Queries) ServiceHandlerRs[TRequest, TResponse]
	Request TRequest
}

var ErrNoRows = pgx.ErrNoRows

func WithTransaction[Request any, TService any](ctx context.Context, params Params[Request, TService]) error {
	conn, err := params.DbPool.Acquire(ctx) // get the current connection from pool

	if err != nil {
		panic(err)
	}

	defer conn.Release() // release the connection back to the pool

	tx, err := conn.Begin(ctx)

	if err != nil {
		panic(err)
	}

	queries := pgstore.New(tx)
	instance := params.New(queries)
	fn := instance

	if err := fn.Handler(ctx, params.Request); err != nil && err != ErrNoRows {
		tx.Rollback(ctx)
		slog.Error("Transaction error, rolling back...", err.Error())

		panic(err)
	}

	return tx.Commit(ctx)
}

func WithTransactionRs[TRequest any, TResponse any, TService any](ctx context.Context, params ParamsRs[TRequest, TResponse, TService]) (TResponse, error) {
	conn, err := params.DbPool.Acquire(ctx) // get the current connection from pool

	if err != nil {
		panic(err)
	}

	defer conn.Release() // release the connection back to the pool

	tx, err := conn.Begin(ctx)

	if err != nil {
		panic(err)
	}

	queries := pgstore.New(tx)
	instance := params.New(queries)
	fn := instance

	response, err := fn.Handler(ctx, params.Request)

	if err != nil && err != ErrNoRows {
		tx.Rollback(ctx)
		slog.Error("Transaction error, rolling back...", err.Error())

		panic(err)
	}

	return response, tx.Commit(ctx)
}
