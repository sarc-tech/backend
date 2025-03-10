package models

import "database/sql"

type User struct {
	ID         int            `json:"id"`
	YandexID   string         `db:"yandex_id" json:"yandex_id"`
	Surname    string         `json:"surname"`
	Name       string         `json:"name"`
	Patronymic sql.NullString `json:"patronymic"`
	CallSign   sql.NullString `db:"call_sign" json:"call_sign"`
	Gender     sql.NullString `db:"gender" json:"gender"`
	Birthdate  sql.NullString `db:"birthdate" json:"birthdate"`
	Phone      sql.NullString `db:"phone" json:"phone"`
	Vk         sql.NullString `db:"vk" json:"vk"`
	Telegram   sql.NullString `db:"telegram" json:"telegram"`
	Email      sql.NullString `db:"email" json:"email"`
	Verified   bool           `db:"approval" json:"verified"`
	CreatedAt  string         `db:"created_at" json:"created_at"`
	UpdatedAt  string         `db:"updated_at" json:"updated_at"`
}

type Session struct {
	ID     string `json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
}
