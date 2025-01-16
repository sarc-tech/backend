package api

import (
	"context"

	"github.com/sarc-tech/backend/internal/oas"
)

type UsersHandler interface {
}

// CheckSms implements oas.Handler.
func (h Handler) CheckSms(ctx context.Context, params oas.CheckSmsParams) (oas.CheckSmsRes, error) {
	return &oas.Token{"55555"}, nil
}

// GetUser implements oas.Handler.
func (h Handler) GetUser(ctx context.Context) (oas.GetUserRes, error) {
	return &oas.User{ID: "1", Fio: "Иванов Иван Иванович", Role: "Admin"}, nil
}

// Logout implements oas.Handler.
func (h Handler) Logout(ctx context.Context) (oas.LogoutRes, error) {
	return &oas.LogoutOK{}, nil
}

// SendSms implements oas.Handler.
func (h Handler) SendSms(ctx context.Context, params oas.SendSmsParams) (oas.SendSmsRes, error) {
	return &oas.SendSmsOK{}, nil
}
