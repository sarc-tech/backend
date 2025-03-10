package api

import (
	"context"
	"fmt"
	"github.com/sarc-tech/backend/internal/models"

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
	DB *sqlx.DB
}

// HandleBearerAuth implements oas.SecurityHandler.
func (h HandlerSecurity) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {
	session := models.Session{}
	err := h.DB.Get(&session, "SELECT * FROM sessions WHERE id = $1", t.Token)
	if err != nil {
		return ctx, fmt.Errorf("unauthorized")
	}
	ctx = context.WithValue(ctx, "TOKEN", t.Token)
	return ctx, nil
}
