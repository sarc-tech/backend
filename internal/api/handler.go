package api

import (
	"context"
	"fmt"

	"github.com/sarc-tech/backend/internal/repo"

	"github.com/sarc-tech/backend/internal/oas"
)

// Compile-time check for Handler.
var _ oas.Handler = (*Handler)(nil)

type Handler struct {
	repo                     *repo.RepoHandler
	oas.UnimplementedHandler // automatically implement all methods
}

func NewHandler(repo *repo.RepoHandler) *Handler {
	return &Handler{
		repo: repo,
	}
}

// HandleUserLogin implements oas.Handler.
type HandlerSecurity struct {
	repo                     *repo.RepoHandler
}

func NewHandlerSecurity(repo *repo.RepoHandler) *HandlerSecurity {
	return &HandlerSecurity{
		repo: repo,
	}
}

// HandleBearerAuth implements oas.SecurityHandler.
func (h HandlerSecurity) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {
	_, err := h.repo.GetSession(t.Token)
	if err != nil {
		return ctx, fmt.Errorf("unauthorized")
	}
	ctx = context.WithValue(ctx, "TOKEN", t.Token)
	return ctx, nil
}
