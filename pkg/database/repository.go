package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repository[T any] interface {
	Select(ctx context.Context, query string, args ...any) ([]*T, error)
	Get(ctx context.Context, query string, args ...any) (*T, error)
	Exec(ctx context.Context, query string, args ...any) (int64, error)
	BeginTransaction(ctx context.Context) (context.Context, func(error) error, error)
}

type repository[T any] struct {
	db *sqlx.DB
}

func NewRepository[T any](db *sqlx.DB) *repository[T] {
	return &repository[T]{db: db}
}

type txKeyType struct{}

var txKey = txKeyType{}

func withTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func txFromCtx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(txKey).(*sqlx.Tx)
	return tx, ok
}

func (r *repository[T]) BeginTransaction(ctx context.Context) (context.Context, func(error) error, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return ctx, nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	ctxWithTx := withTx(ctx, tx)

	cleanup := func(txErr error) error {
		if txErr != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return fmt.Errorf("transaction rollback failed: %w", rollbackErr)
			}
			return txErr
		}
		if commitErr := tx.Commit(); commitErr != nil {
			return fmt.Errorf("transaction commit failed: %w", commitErr)
		}
		return nil
	}

	return ctxWithTx, cleanup, nil
}
func (r *repository[T]) Select(ctx context.Context, query string, args ...any) ([]*T, error) {
	var results []*T

	if tx, ok := txFromCtx(ctx); ok {
		err := tx.SelectContext(ctx, &results, query, args...)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	err := r.db.SelectContext(ctx, &results, query, args...)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Get retrieves a single row from the database and scans it into an instance of T.
func (r *repository[T]) Get(ctx context.Context, query string, args ...any) (*T, error) {
	var result T

	if tx, ok := txFromCtx(ctx); ok {
		err := tx.GetContext(ctx, &result, query, args...)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	err := r.db.GetContext(ctx, &result, query, args...)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Exec executes a query that doesn't return rows, like INSERT, UPDATE, or DELETE.
func (r *repository[T]) Exec(ctx context.Context, query string, args ...any) (int64, error) {
	if tx, ok := txFromCtx(ctx); ok {
		res, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			return 0, err
		}
		return res.RowsAffected()
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
