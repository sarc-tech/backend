package controller

import (
	"fmt"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/oas"
)

func (r *Controller) GetIncidentById(id int) (*models.Incident, error) {
	incident := &models.Incident{}
	err := r.Db.Get(
		incident,
		`SELECT * FROM incident_incident WHERE id = $1;`,
		id,
	)
	return incident, err
}

func (r *Controller) CreateIncident(incidentToCreate *oas.Incident) (*models.Incident, error) {
	incidentModel := models.MapIncidentFromAPi(incidentToCreate)
	rows, err := r.Db.NamedQuery(
		`INSERT INTO incident_incident` + incidentModel.GetSqlInsertMapString() + `RETURNING ` + incidentModel.GetSqlReturningString(), 
		incidentModel,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	incident := models.Incident{}
	if rows.Next() {
        if err := rows.StructScan(&incident); err != nil {
            return &models.Incident{}, err
        }
        return &incident, nil
    }
	return &incident, fmt.Errorf("insert failed, no rows returned")
}

func (r *Controller) DeleteIncident(id int) (err error) {
	_, err = r.Db.Exec(`DELETE FROM incident_incident WHERE id = $1;`, id)
	return
}

func (r *Controller) GetIncidents() ([]*models.Incident, error) {
	var incidents []*models.Incident
	err := r.Db.Select(
		&incidents, 
		`SELECT * FROM incident_incident LIMIT 50;`, 
	)
	return incidents, err
}

func (r *Controller) UpdateIncident(incidentToUpdate *oas.Incident) (*models.Incident, error) {
	incidentModel := models.MapIncidentFromAPi(incidentToUpdate)
	rows, err := r.Db.NamedQuery(
		`UPDATE incident_incident SET` + incidentModel.GetSqlUpdateMapString() + `WHERE id = :id RETURNING ` + incidentModel.GetSqlReturningString(),
		incidentModel,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	incident := models.Incident{}
	if rows.Next() {
        if err := rows.StructScan(&incident); err != nil {
            return &models.Incident{}, err
        }
        return &incident, nil
    }
	return &incident, fmt.Errorf("update failed, no rows returned")
}