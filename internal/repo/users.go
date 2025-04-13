package repo

import (
	"github.com/sarc-tech/backend/internal/models"
)


func (r *RepoHandler) GetUser(token string) (user *models.User, err error) {
	err = r.Db.Get(
		user, 
		`SELECT DISTINCT u.* 
		FROM users u
		JOIN sessions s ON u.id = s.user_id
		WHERE s.id = $1;`, 
		token,
	)
	return
}

func (r *RepoHandler) GetUsers() (users []*models.User, err error) {
	err = r.Db.Select(
		users, 
		`SELECT * FROM users LIMIT 50;`, 
	)
	return
}

func (r *RepoHandler) GetUserById(userId int) (user *models.User, err error) {
	err = r.Db.Get(
		user,
		`SELECT * FROM users WHERE id = $1;`,
		userId,
	)
	return
}

func (r *RepoHandler) UpdateUser(userToUpdate *models.User) (user *models.User, err error) {
	_, err = r.Db.NamedExec(
		`UPDATE users 
		SET yandex_id = :yandex_id,
			surname = :surname,
			name = :name, 
			patronymic = :patronymic,
			call_sign = :call_sign,
			gender = :gender,
			birthdate = :birthdate,
			phone = :phone,
			vk = :vk,
			telegram = :telegram,
			email = :email,
			approval = :approval
			updated_at = NOW()
		WHERE id = :id`,
		userToUpdate,
	)
	return r.GetUserById(userToUpdate.ID)
}