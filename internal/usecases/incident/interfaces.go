package incident

import (
	"context"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/repositories"
)

type Repository interface {
	List(ctx context.Context, opts ...repositories.QueryOption) ([]*models.Incident, error)
	Get(ctx context.Context, opts ...repositories.QueryOption) (*models.Incident, error)
	Delete(ctx context.Context, id int) error
	Add(ctx context.Context, m *models.Incident) error
	Update(ctx context.Context, m *models.Incident) error
	BeginTransaction(ctx context.Context) (context.Context, func(error) error, error)
}
