package models

import (
	"database/sql"
	"encoding/json"
	"io"

	"github.com/sarc-tech/backend/internal/oas"
	"github.com/sarc-tech/backend/internal/utils"
)

type UserModel struct {
	ID         int            `db:"id"`
	YandexID   string         `db:"yandex_id"`
	FirstName  sql.NullString `db:"first_name"`
	LastName   sql.NullString `db:"last_name"`
	Patronymic sql.NullString `db:"patronymic"`
	CallSign   sql.NullString `db:"call_sign"`
	Gender     sql.NullString `db:"gender"`
	Birthdate  sql.NullTime   `db:"birthdate"`
	Phone      sql.NullString `db:"phone"`
	Vk         sql.NullString `db:"vk"`
	Telegram   sql.NullString `db:"telegram"`
	Email      sql.NullString `db:"email"`
	Approval   bool           `db:"approval"`
	CreatedAt  string         `db:"created_at"`
	UpdatedAt  string         `db:"updated_at"`
}

func (u *UserModel) MapUserToApi() *oas.User {
	return &oas.User{
		ID: 	  	oas.NewOptInt(u.ID),
		YandexId: 	u.YandexID,
		FirstName:  oas.NewOptString(u.FirstName.String),
		LastName:  	oas.NewOptString(u.LastName.String),
		Gender:   	oas.NewOptString(u.Gender.String),
		Phone:    	oas.NewOptString(u.Phone.String),
		Email:    	oas.NewOptString(u.Email.String),
		Role:     	"Admin",
		Approval: 	true,
	}
}

func (u *UserModel) GetSqlUpdateMapString() string {
	return `
		yandex_id = :yandex_id,
		last_name = :last_name,
		first_name = :first_name, 
		patronymic = :patronymic,
		call_sign = :call_sign,
		gender = :gender,
		birthdate = :birthdate,
		phone = :phone,
		vk = :vk,
		telegram = :telegram,
		email = :email,
		approval = :approval
	`
}

func (u *UserModel) GetSqlInsertMapString() string {
	return `
		(yandex_id, last_name, first_name, patronymic, call_sign, gender, birthdate, phone, vk, telegram, email, approval) VALUES 
		(:yandex_id, :last_name, :first_name, :patronymic, :call_sign, :gender, :birthdate, :phone, :vk, :telegram, :email, true)
	`
}

func (u *UserModel) GetSqlReturningString() string {
	return `id, yandex_id, last_name, first_name, patronymic, call_sign, gender, birthdate, phone, vk, telegram, email, approval`
}

func MapUserFromApi(apiUser oas.User) (user *UserModel) {
	user = &UserModel{
		YandexID: apiUser.YandexId,
		LastName: utils.ToNullStringFromOptString(apiUser.LastName),
		FirstName: utils.ToNullStringFromOptString(apiUser.FirstName),
		Patronymic: utils.ToNullStringFromOptString(apiUser.Patronymic),
		CallSign: utils.ToNullStringFromOptString(apiUser.CallSign),
		Gender: utils.ToNullStringFromOptString(apiUser.Gender),
		Birthdate: utils.ToNullDateFromOptDate(apiUser.Birthdate),
		Phone: utils.ToNullStringFromOptString(apiUser.Phone),
		Vk: utils.ToNullStringFromOptString(apiUser.Vk),
		Telegram: utils.ToNullStringFromOptString(apiUser.Telegram),
		Email: utils.ToNullStringFromOptString(apiUser.Email),
		Approval: apiUser.Approval,
	}
	if v, ok := apiUser.ID.Get(); ok {
		user.ID = v
	}
	return
}

func MapUserFromYandex(rawYandexUser io.ReadCloser, token string) (user *UserModel, err error) {
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

	var yandexUser User
	err = json.NewDecoder(rawYandexUser).Decode(&yandexUser)
	if err != nil {
		return nil, err
	}

	user = &UserModel{
		YandexID:   token,
		FirstName:  utils.ToNullStringFromString(&yandexUser.FirstName),
		LastName:   utils.ToNullStringFromString(&yandexUser.LastName),
		Email:      utils.ToNullStringFromString(&yandexUser.Email),
		Birthdate:  utils.ToNullDateFromString(yandexUser.Birthday),
		Gender:     utils.ToNullStringFromString(&yandexUser.Sex),
      	Phone:      utils.ToNullStringFromString(&yandexUser.Default.Number),
	}
	return
}

type Session struct {
	ID     string `db:"id"`
	UserID int    `db:"user_id"`
}
