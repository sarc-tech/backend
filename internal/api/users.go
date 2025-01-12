package api

import (
	"context"

	"github.com/sarc-tech/backend/internal/oas"
)

type UsersHandler interface {
}

// CheckSms implements oas.Handler.
func (h Handler) CheckSms(ctx context.Context, params oas.CheckSmsParams) (oas.CheckSmsRes, error) {
	panic("unimplemented")
}

// GetUser implements oas.Handler.
func (h Handler) GetUser(ctx context.Context) (oas.GetUserRes, error) {
	panic("unimplemented")
}

// Logout implements oas.Handler.
func (h Handler) Logout(ctx context.Context) (oas.LogoutRes, error) {
	panic("unimplemented")
}

// SendSms implements oas.Handler.
func (h Handler) SendSms(ctx context.Context, params oas.SendSmsParams) (oas.SendSmsRes, error) {
	panic("unimplemented")
}
