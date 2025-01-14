package api

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sarc-tech/backend/internal/oas"
)

// Compile-time check for Handler.
var _ oas.Handler = (*Handler)(nil)

type Handler struct {
	DB                       *sqlx.DB
	oas.UnimplementedHandler // automatically implement all methods
}

type HandlerSecurity struct {
}

// HandleBearerAuth implements oas.SecurityHandler.
func (h HandlerSecurity) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {
	if t.Token == "55555" {
		return ctx, nil
	}
	return ctx, fmt.Errorf("unauthorized")
}
