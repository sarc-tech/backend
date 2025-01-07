package repositories

import "context"

type Repository[T any] interface {
	Select(ctx context.Context, query string, args ...any) ([]*T, error)
	Get(ctx context.Context, query string, args ...any) (*T, error)
	Exec(ctx context.Context, query string, args ...any) (int64, error)
	BeginTransaction(ctx context.Context) (context.Context, func(error) error, error)
}
