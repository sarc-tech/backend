package incident

import (
	"context"
	"fmt"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/repositories"
)

type Model struct {
	ID   int    `db:"id"`
	name string `db:"name"`
}

type repository struct {
	db repositories.Repository[Model]
}

func NewRepository(db repositories.Repository[Model]) *repository {
	return &repository{db: db}
}

func (r *repository) List(ctx context.Context, opts ...repositories.QueryOption) ([]*models.Statues, error) {
	options := &repositories.QueryOptions{}
	for _, opt := range opts {
		opt(options)
	}

	query := `SELECT * FROM statues`

	where, args := options.BuildConditions()
	query += where

	statues, err := r.db.Select(ctx, options.AddLimitAndOffset(query), args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get incidents: %w", err)
	}

	return r.dtos(statues), nil
}

func (r *repository) Get(ctx context.Context, opts ...repositories.QueryOption) (*models.Statues, error) {
	options := &repositories.QueryOptions{}
	for _, opt := range opts {
		opt(options)
	}

	query := `SELECT * FROM statues`

	where, args := options.BuildConditions()
	query += where

	statues, err := r.db.Get(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get incident: %w", err)
	}

	return r.dto(statues), nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM statues WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete statues: %w", err)
	}

	return nil
}

func (r *repository) Add(ctx context.Context, m *models.Statues) error {
	statues := r.orm(m)
	result, err := r.db.Get(ctx, `
		INSERT INTO statues (name)
		VALUES ($1) RETURNING id
	`, statues.name)
	if err != nil {
		return fmt.Errorf("failed to add incident: %w", err)
	}

	m.ID = result.ID

	return nil
}

func (r *repository) Update(ctx context.Context, m *models.Statues) error {
	statues := r.orm(m)
	_, err := r.db.Exec(ctx, `
		UPDATE statues
		SET name = $1
		WHERE id = $2
	`, statues.name, statues.ID)
	if err != nil {
		return fmt.Errorf("failed to update incident: %w", err)
	}

	return nil
}

func (r *repository) BeginTransaction(ctx context.Context) (context.Context, func(error) error, error) {
	return r.db.BeginTransaction(ctx)
}
