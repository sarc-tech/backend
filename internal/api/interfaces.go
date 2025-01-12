package api

import (
	"context"

	"github.com/sarc-tech/backend/internal/models"
)

type IncidentUsecase interface {
	List(ctx context.Context, limit, offset int) ([]*models.Incident, error)
	Get(ctx context.Context, id int) (*models.Incident, error)
	Delete(ctx context.Context, id int) error
	Add(ctx context.Context, m *models.Incident) error
	Update(ctx context.Context, m *models.Incident) error
}

type StatusesHandler interface {
}

type UsersHandler interface {
}
