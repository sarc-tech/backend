package controller

import (
	"github.com/sarc-tech/backend/internal/models"
)

func (r *Controller) GetStatuses() ([]*models.Status, error) {
	statuses := []*models.Status{}
	err := r.Db.Select(
		&statuses, 
		`SELECT * FROM dirs_status LIMIT 50;`, 
	)
	return statuses, err
}