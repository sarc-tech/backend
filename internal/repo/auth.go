package repo

import (
	"github.com/sarc-tech/backend/internal/models"
)

func (r *RepoHandler) GetSession(token string) (models.Session, error) {
	session := models.Session{}
	err := r.Db.Get(&session, "SELECT * FROM sessions WHERE id = $1", token)
	return session, err
}