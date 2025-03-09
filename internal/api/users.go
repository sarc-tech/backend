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

// CheckSms implements oas.Handler.
func (h Handler) CheckSms(ctx context.Context, params oas.CheckSmsParams) (oas.CheckSmsRes, error) {
	return &oas.Token{"55555"}, nil
}

// GetUser implements oas.Handler.
func (h Handler) GetUser(ctx context.Context) (oas.GetUserRes, error) {
	return &oas.User{ID: "1", Name: "Иванов Иван Иванович", Role: "Admin"}, nil
}

// Logout implements oas.Handler.
func (h Handler) Logout(ctx context.Context) (oas.LogoutRes, error) {
	return &oas.LogoutOK{}, nil
}

// SendSms implements oas.Handler.
func (h Handler) SendSms(ctx context.Context, params oas.SendSmsParams) (oas.SendSmsRes, error) {
	return &oas.SendSmsOK{}, nil
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
	err = h.DB.Get(&checkUser, "Select yandex_id from users where yandex_id =$1", req.ID)
	if err == sql.ErrNoRows {
		// Пользователь не обнаружен в системе, добавим его в базу данных
		_, err = h.DB.Exec("Insert into users (yandex_id,surname,name,gender, phone, email ) values ($1,$2,$3,$4,$5,$6)",
			req.ID, req.LastName, req.FirstName, req.Sex, req.Default.Number, req.Email)
		if err != nil {
			return &oas.CheckUserBadRequest{}, err
		}
	}

	// Получаем пользователя из базы данных
	err = h.DB.Get(&checkUser, "Select * from users where yandex_id =$1", req.ID)
	if err != nil {
		return &oas.CheckUserBadRequest{}, err
	}
	// Заполним сессию пользователя
	_, err = h.DB.Exec("INSERT INTO sessions (user_id,id) values ($1,$2)", checkUser.ID, params.Token)
	if err != nil {
		return &oas.CheckUserBadRequest{}, err
	}
	// Возвращаем пользователя
	rez := oas.User{ID: strconv.Itoa(checkUser.ID), YandexId: checkUser.YandexID,
		Surname: oas.NewOptString(checkUser.Surname), Name: checkUser.Name, Gender: oas.NewOptString(checkUser.Gender), Phone: oas.NewOptString(checkUser.Phone.String),
		Email: oas.NewOptString(checkUser.Email.String),
		Role:  "Admin", Approval: true}
	return &rez, nil
}
