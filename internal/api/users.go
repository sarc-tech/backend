package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/oas"
)

type UsersHandler interface {
}

// GetUser implements oas.Handler.
func (h Handler) GetUser(ctx context.Context) (oas.GetUserRes, error) {
	user, err := h.repo.GetUser(ctx.Value("TOKEN").(string))
	if err != nil {
		return nil, fmt.Errorf("failed to get user")
	}

	return &oas.UserResponse{Data: *user.MapUserToApi()}, nil
}

func (h Handler) GetUsers(ctx context.Context) (oas.GetUsersRes, error) {
	users, err := h.repo.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get users")
	}

	res := make([]oas.User, len(users), len(users))
	for i, user := range users {
		res[i] = *user.MapUserToApi()
	}

	return &oas.UsersResponse{Data: res}, nil
}

func (h Handler) GetUserById(ctx context.Context, params oas.GetUserByIdParams) (oas.GetUserByIdRes, error) {
	user, err := h.repo.GetUserById(params.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id")
	}

	return &oas.UserResponse{Data: *user.MapUserToApi()}, nil
}

func (h Handler) UpdateUser(ctx context.Context, req *oas.User) (oas.UpdateUserRes, error) {
	user, err := h.repo.UpdateUser(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update user")
	}

	return &oas.UserResponse{Data: *user.MapUserToApi()}, nil
}

// Logout implements oas.Handler.
func (h Handler) Logout(ctx context.Context) (oas.LogoutRes, error) {
	err := h.repo.DeleteSession(ctx.Value("TOKEN").(string))
	if err != nil {
		return nil, fmt.Errorf("failed to delete session")
	}
	return &oas.LogoutOK{}, nil
}

func (h Handler) CheckUser(ctx context.Context, params oas.CheckUserParams) (oas.CheckUserRes, error) {
	// Проверяем пользователя в базе данных
	request, err := http.NewRequest(http.MethodGet, "https://login.yandex.ru/info?format=json", nil)
	if err != nil {
		return &oas.CheckUserBadRequest{}, err
	}
	request.Header.Set("Authorization", "OAuth "+params.Token)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return &oas.CheckUserBadRequest{}, err
	}
	if response.StatusCode != http.StatusOK {
		return &oas.CheckUserBadRequest{}, nil
	}

	defer response.Body.Close()

	userFromYandex, err := models.MapUserFromYandex(response.Body, params.Token)
	if err != nil {
		return &oas.CheckUserBadRequest{}, nil
	}
	// Проверим существует ли у нас пользователь
	userFromDb, err := h.repo.GetUserByYandexId(userFromYandex.YandexID)
	if err == sql.ErrNoRows {
		// Пользователь не обнаружен в системе, добавим его в базу данных
		userFromDb, err = h.repo.CreateUserFromModel(userFromYandex)
		if err != nil {
			return &oas.CheckUserBadRequest{}, err
		}
	}
	// Заполним сессию пользователя
	err = h.repo.CreateSession(userFromDb)
	if err != nil {
		return &oas.CheckUserBadRequest{}, err
	}
	// Возвращаем пользователя
	return &oas.UserResponse{Data: *userFromDb.MapUserToApi()}, nil
}
