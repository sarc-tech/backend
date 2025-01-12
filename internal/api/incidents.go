package api

import (
	"context"
	"fmt"
	"log"

	"github.com/sarc-tech/backend/internal/oas"
)

// AddIncidents implements oas.Handler.
func (h Handler) AddIncidents(ctx context.Context, req oas.AddIncidentsReq) (oas.AddIncidentsRes, error) {
	request, ok := req.(*oas.AddIncidentsApplicationJSON)
	if !ok {
		return &oas.AddIncidentsBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	q := `
	INSERT INTO incidents (region, fio, statusId, date)
	VALUES ($1, $2, $3, $4) RETURNING id
	`
	h.DB.MustExec(q, request.Region, request.Fio, request.StatusId, request.Date)

	//	id, err := rez.LastInsertId()

	//	if err != nil {
	//		return &oas.AddIncidentsBadRequest{}, fmt.Errorf("failed to get last insert id")
	//	}

	return &oas.Incident{
		//		ID:       string(id),
		Region:   request.Region,
		Fio:      request.Fio,
		StatusId: request.StatusId,
		Date:     request.Date,
	}, nil
}

// DeleteIncident implements oas.Handler.
func (h Handler) DeleteIncident(ctx context.Context, params oas.DeleteIncidentParams) (oas.DeleteIncidentRes, error) {
	q := `
	DELETE FROM incidents
	where id=$1
	`
	h.DB.MustExec(q, params.IncidentId)

	return &oas.DeleteIncidentOK{}, nil
}

// GetIncidentById implements oas.Handler.
func (h Handler) GetIncidentById(ctx context.Context, params oas.GetIncidentByIdParams) (oas.GetIncidentByIdRes, error) {
	rez := oas.Incident{}
	err := h.DB.Get(&rez, "SELECT * FROM incidents WHERE id=$1", params.IncidentId)
	if err != nil {
		return &oas.GetIncidentByIdBadRequest{}, fmt.Errorf("failed to get incident by id")
	}
	return &rez, nil
}

// GetIncidents implements oas.Handler.
func (h Handler) GetIncidents(ctx context.Context) (oas.GetIncidentsRes, error) {
	req := make([]oas.Incident, 0, 50)

	// Получим все инциденты
	incident := oas.Incident{}
	rows, err := h.DB.Queryx("SELECT * FROM incidents")
	if err != nil {
		return &oas.GetIncidentsBadRequest{}, fmt.Errorf("failed to get incidents")
	}
	for rows.Next() {
		err := rows.StructScan(&incident)
		if err != nil {
			log.Fatalln(err)
		}
		req = append(req, incident)
	}

	// Получим все статусы
	req_status := make([]oas.Status, 0, 50)

	status := oas.Status{}
	rows, err = h.DB.Queryx("SELECT * FROM statuses")
	if err != nil {
		return &oas.GetIncidentsBadRequest{}, fmt.Errorf("failed to get incidents")
	}
	for rows.Next() {
		err := rows.StructScan(&status)
		if err != nil {
			log.Fatalln(err)
		}
		req_status = append(req_status, status)
	}

	return &oas.IncidentsResponse{TrackingId: "123", Status: "Ok", Data: req, Statuses: req_status}, nil
}

// UpdateIncidents implements oas.Handler.
func (h Handler) UpdateIncidents(ctx context.Context, req oas.UpdateIncidentsReq) (oas.UpdateIncidentsRes, error) {
	request, ok := req.(*oas.UpdateIncidentsApplicationJSON)
	if !ok {
		return &oas.UpdateIncidentsBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	q := `
	UPDATE incidents SET id=$1, region=$2, fio=$3, statusid=$4, date=$5
	where id=$1
	`
	rez := h.DB.MustExec(q, request.ID, request.Region, request.Fio, request.StatusId, request.Date)

	_, err := rez.RowsAffected()

	if err != nil {
		return &oas.UpdateIncidentsBadRequest{}, fmt.Errorf("failed to get last insert id")
	}

	return &oas.Incident{
		ID:       request.ID,
		Region:   request.Region,
		Fio:      request.Fio,
		StatusId: request.StatusId,
		Date:     request.Date,
	}, nil
}
