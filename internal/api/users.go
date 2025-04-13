package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/oas"
	"net/http"
	"strconv"
)

type UsersHandler interface {
}

// GetUser implements oas.Handler.
func (h Handler) GetUser(ctx context.Context) (oas.GetUserRes, error) {
	user, err := h.repo.GetUser(ctx.Value("TOKEN").(string))
	if err != nil {
		return nil, err
	}

	return &oas.User{ID: strconv.Itoa(user.ID),
		Name:     user.Name,
		YandexId: user.YandexID,
		Surname:  oas.NewOptString(user.Surname),
		Gender:   oas.NewOptString(user.Gender.String),
		Phone:    oas.NewOptString(user.Phone.String),
		Email:    oas.NewOptString(user.Email.String),
		Role:     "Admin",
		Approval: true}, nil
}

func (h Handler) GetUsers(ctx context.Context) (oas.GetUsersRes, error) {
	users, err := h.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	req := make([]oas.User, 0, 50)
	for i, user := range users {
		req[i] = oas.User{
			ID:       strconv.Itoa(user.ID),
			Name:     user.Name,
			YandexId: user.YandexID,
			Surname:  oas.NewOptString(user.Surname),
			Gender:   oas.NewOptString(user.Gender.String),
			Phone:    oas.NewOptString(user.Phone.String),
			Email:    oas.NewOptString(user.Email.String),
			Role:     "Admin",
			Approval: true,
		}
	}

	return &oas.UsersResponse{Data: req}, nil
}

func (h Handler) GetUserById(ctx context.Context, params oas.GetUserByIdParams) (oas.GetUserByIdRes, error) {

	user, err := h.repo.GetUserById(params.UserId)
	if err != nil {
		return nil, err
	}

	return &oas.User{ID: strconv.Itoa(user.ID),
		Name:     user.Name,
		YandexId: user.YandexID,
		Surname:  oas.NewOptString(user.Surname),
		Gender:   oas.NewOptString(user.Gender.String),
		Phone:    oas.NewOptString(user.Phone.String),
		Email:    oas.NewOptString(user.Email.String),
		Role:     "Admin",
		Approval: true}, nil
}

func (h Handler) UpdateUser(ctx context.Context, req *oas.User) (oas.UpdateUserRes, error) {
	var user models.User
	err := h.Db.Get(&user, "SELECT * FROM users WHERE id = $1", req.ID)
	if err != nil {
		return nil, err
	}

	return &oas.User{ID: strconv.Itoa(user.ID),
		Name:     user.Name,
		YandexId: user.YandexID,
		Surname:  oas.NewOptString(user.Surname),
		Gender:   oas.NewOptString(user.Gender.String),
		Phone:    oas.NewOptString(user.Phone.String),
		Email:    oas.NewOptString(user.Email.String),
		Role:     "Admin",
		Approval: true}, nil
}

// Logout implements oas.Handler.
func (h Handler) Logout(ctx context.Context) (oas.LogoutRes, error) {
	return &oas.LogoutOK{}, nil
}

func (h Handler) CheckUser(ctx context.Context, params oas.CheckUserParams) (oas.CheckUserRes, error) {

	type Phone struct {
		Number string `json:"number"`
	}

	type User struct {
		ID        string `json:"id"`
		Login     string `json:"login"`
		Psuid     string `json:"psuid"`
		Birthday  string `json:"birthday"`
		Sex       string `json:"sex"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"default_email"`
		Default   Phone  `json:"default_phone"`
	}

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

	var req User
	err = json.NewDecoder(response.Body).Decode(&req)
	if err != nil {
		return &oas.CheckUserBadRequest{}, nil
	}
	fmt.Println(req)
	// Проверим существует ли у нас пользователь
	checkUser := models.User{}
	err = h.Db.Get(&checkUser, "Select yandex_id from users where yandex_id =$1", req.ID)
	if err == sql.ErrNoRows {
		// Пользователь не обнаружен в системе, добавим его в базу данных
		var sex sql.NullString
		if req.Sex == "male" || req.Sex == "female" {
			sex = sql.NullString{String: req.Sex, Valid: true}
		} else {
			sex = sql.NullString{String: "", Valid: false}
		}

		_, err = h.Db.Exec("Insert into users (yandex_id,surname,name,gender, phone, email, approval) values ($1,$2,$3,$4,$5,$6,true)",
			req.ID, req.LastName, req.FirstName, sex, req.Default.Number, req.Email)
		if err != nil {
			return &oas.CheckUserBadRequest{}, err
		}
	}

	// Получаем пользователя из базы данных
	err = h.Db.Get(&checkUser, "Select * from users where yandex_id =$1", req.ID)
	if err != nil {
		return &oas.CheckUserBadRequest{}, err
	}
	// Заполним сессию пользователя
	_, err = h.Db.Exec("INSERT INTO sessions (user_id,id) values ($1,$2)", checkUser.ID, params.Token)
	if err != nil {
		return &oas.CheckUserBadRequest{}, err
	}
	// Возвращаем пользователя
	rez := oas.User{ID: strconv.Itoa(checkUser.ID), YandexId: checkUser.YandexID,
		Surname: oas.NewOptString(checkUser.Surname), Name: checkUser.Name, Gender: oas.NewOptString(checkUser.Gender.String), Phone: oas.NewOptString(checkUser.Phone.String),
		Email: oas.NewOptString(checkUser.Email.String),
		Role:  "Admin", Approval: true}
	return &rez, nil
}
