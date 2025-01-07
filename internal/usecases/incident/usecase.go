package incident

import (
	"context"
	"fmt"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/repositories"
)

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository: repository}
}

func (u *usecase) List(ctx context.Context, limit, offset int) ([]*models.Incident, error) {
	incidents, err := u.repository.List(ctx, repositories.WithLimit(limit), repositories.WithOffset(offset))
	if err != nil {
		return nil, err
	}

	return incidents, nil
}

func (u *usecase) Get(ctx context.Context, id int) (*models.Incident, error) {
	incident, err := u.repository.Get(ctx,
		repositories.WithAndFilter("id", "=", id),
	)
	if err != nil {
		return nil, err
	}

	return incident, nil
}

func (u *usecase) Delete(ctx context.Context, id int) error {
	ctx, commit, err := u.repository.BeginTransaction(ctx)
	defer commit(err)

	if err = u.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete incident: %w", err)
	}

	return nil
}

func (u *usecase) Add(ctx context.Context, m *models.Incident) error {
	ctx, commit, err := u.repository.BeginTransaction(ctx)
	defer commit(err)

	if err = u.repository.Add(ctx, m); err != nil {
		return fmt.Errorf("failed to add incident: %w", err)
	}

	return nil
}

func (u *usecase) Update(ctx context.Context, m *models.Incident) error {
	ctx, commit, err := u.repository.BeginTransaction(ctx)
	defer commit(err)

	if err = u.repository.Update(ctx, m); err != nil {
		return fmt.Errorf("failed to update incident: %w", err)
	}

	return nil
}
