package repo

import (
	"fmt"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/oas"
)


func (r *RepoHandler) GetUser(token string) (*models.UserModel, error) {
	user := &models.UserModel{}
	err := r.Db.Get(
		user, 
		`SELECT DISTINCT u.* 
		FROM core_portaluser u
		JOIN core_session s ON u.id = s.user_id
		WHERE s.id = $1;`, 
		token,
	)
	return user, err
}

func (r *RepoHandler) GetUsers() (users []*models.UserModel, err error) {
	err = r.Db.Select(
		users, 
		`SELECT * FROM core_portaluser LIMIT 50;`, 
	)
	return
}

func (r *RepoHandler) GetUserById(userId int) (*models.UserModel, error) {
	user := &models.UserModel{}
	err := r.Db.Get(
		user,
		`SELECT * FROM core_portaluser WHERE id = $1;`,
		userId,
	)
	return user, err
}

func (r *RepoHandler) GetUserByYandexId(yandexId string) (*models.UserModel, error) {
	user := &models.UserModel{}
	err := r.Db.Get(
		user,
		`SELECT * FROM core_portaluser WHERE yandex_id = $1 LIMIT 1;`,
		yandexId,
	)
	return user, err
}

func (r *RepoHandler) UpdateUser(userToUpdate *oas.User) (*models.UserModel, error) {
	userModel := models.MapUserFromApi(*userToUpdate)
	rows, err := r.Db.NamedQuery(
		`UPDATE core_portaluser SET` + userModel.GetSqlUpdateMapString() + `WHERE id = :id RETURNING ` + userModel.GetSqlReturningString(),
		userModel,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	user := models.UserModel{}
	if rows.Next() {
        if err := rows.StructScan(&user); err != nil {
            return &models.UserModel{}, err
        }
        return &user, nil
    }
	return &user, fmt.Errorf("update failed, no rows returned")
}

func (r *RepoHandler) CreateUserFromModel(userToCreate *models.UserModel) (*models.UserModel, error) {
	rows, err := r.Db.NamedQuery(
		`INSERT INTO core_portaluser` + userToCreate.GetSqlInsertMapString() + `RETURNING ` + userToCreate.GetSqlReturningString(),
		userToCreate,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	user := models.UserModel{}
	if rows.Next() {
        if err := rows.StructScan(&user); err != nil {
            return &models.UserModel{}, err
        }
        return &user, nil
    }
	return &user, fmt.Errorf("insert failed, no rows returned")
}

func (r *RepoHandler) CreateSession(user *models.UserModel) (err error) {
	_, err = r.Db.Exec(
		`INSERT INTO core_session (user_id, id) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING;`, 
		user.ID, 
		user.YandexID,
	)
	r.sessionTTL.Set(user.YandexID, user.ID)
	return
}

func (r *RepoHandler) DeleteSession(token string) error {
	_, err := r.Db.Exec(
		`DELETE FROM core_session WHERE id = $1;`, 
		token,
	)
	return err
}