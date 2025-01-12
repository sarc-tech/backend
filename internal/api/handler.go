package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/sarc-tech/backend/internal/oas"
)

// Compile-time check for Handler.
var _ oas.Handler = (*Handler)(nil)

type Handler struct {
	DB                       *sqlx.DB
	oas.UnimplementedHandler // automatically implement all methods
}
