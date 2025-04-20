package api

import (
	"context"
	"fmt"

	"github.com/sarc-tech/backend/internal/controller"

	"github.com/sarc-tech/backend/internal/oas"
)

// Compile-time check for Handler.
var _ oas.Handler = (*Handler)(nil)

type Handler struct {
	controller                     *controller.Controller
	oas.UnimplementedHandler // automatically implement all methods
}

func NewHandler(controller *controller.Controller) *Handler {
	return &Handler{
		controller: controller,
	}
}

// HandleUserLogin implements oas.Handler.
type HandlerSecurity struct {
	controller                     *controller.Controller
}

func NewHandlerSecurity(controller *controller.Controller) *HandlerSecurity {
	return &HandlerSecurity{
		controller: controller,
	}
}

// HandleBearerAuth implements oas.SecurityHandler.
func (h HandlerSecurity) HandleBearerAuth(ctx context.Context, operationName oas.OperationName, t oas.BearerAuth) (context.Context, error) {
	valid := h.controller.GetSession(t.Token)
	if !valid {
		return ctx, fmt.Errorf("unauthorized")
	}
	ctx = context.WithValue(ctx, "TOKEN", t.Token)
	return ctx, nil
}
