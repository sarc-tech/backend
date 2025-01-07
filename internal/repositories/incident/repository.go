package incident

import (
	"context"
	"fmt"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/repositories"
)

type Model struct {
	ID         int    `db:"id"`
	IsTraining bool   `db:"is_training"`
	Region     string `db:"region"`
	FIO        string `db:"fio"`
	Status     string `db:"status"`
	Date       string `db:"date"`
}

type repository struct {
	db repositories.Repository[Model]
}

func NewRepository(db repositories.Repository[Model]) *repository {
	return &repository{db: db}
}

func (r *repository) List(ctx context.Context, opts ...repositories.QueryOption) ([]*models.Incident, error) {
	options := &repositories.QueryOptions{}
	for _, opt := range opts {
		opt(options)
	}

	query := `SELECT * FROM incidents`

	where, args := options.BuildConditions()
	query += where

	incidents, err := r.db.Select(ctx, options.AddLimitAndOffset(query), args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get incidents: %w", err)
	}

	return r.dtos(incidents), nil
}

func (r *repository) Get(ctx context.Context, opts ...repositories.QueryOption) (*models.Incident, error) {
	options := &repositories.QueryOptions{}
	for _, opt := range opts {
		opt(options)
	}

	query := `SELECT * FROM incidents`

	where, args := options.BuildConditions()
	query += where

	incident, err := r.db.Get(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get incident: %w", err)
	}

	return r.dto(incident), nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM incidents WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete incident: %w", err)
	}

	return nil
}

func (r *repository) Add(ctx context.Context, m *models.Incident) error {
	incident := r.orm(m)
	result, err := r.db.Get(ctx, `
		INSERT INTO incidents (is_training, region, fio, status, date)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`, incident.IsTraining, incident.Region, incident.FIO, incident.Status, incident.Date)
	if err != nil {
		return fmt.Errorf("failed to add incident: %w", err)
	}

	m.ID = result.ID

	return nil
}

func (r *repository) Update(ctx context.Context, m *models.Incident) error {
	incident := r.orm(m)
	_, err := r.db.Exec(ctx, `
		UPDATE incidents
		SET is_training = $1, region = $2, fio = $3, status = $4, date = $5
		WHERE id = $6
	`, incident.IsTraining, incident.Region, incident.FIO, incident.Status, incident.Date, incident.ID)
	if err != nil {
		return fmt.Errorf("failed to update incident: %w", err)
	}

	return nil
}

func (r *repository) BeginTransaction(ctx context.Context) (context.Context, func(error) error, error) {
	return r.db.BeginTransaction(ctx)
}
