package api

import (
	"context"
	"fmt"
	"log"

	"github.com/sarc-tech/backend/internal/oas"
)

type StatusesHandler interface {
}

// AddStatus implements oas.Handler.
func (h Handler) AddStatus(ctx context.Context, req oas.AddStatusReq) (oas.AddStatusRes, error) {
	request, ok := req.(*oas.AddStatusApplicationJSON)
	if !ok {
		return &oas.AddStatusBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	q := `
	INSERT INTO statuses (name)
	VALUES ($1) RETURNING id
	`
	h.Db.MustExec(q, request.Name)

	//	id, err := rez.LastInsertId()

	//	if err != nil {
	//		return &oas.AddIncidentsBadRequest{}, fmt.Errorf("failed to get last insert id")
	//	}

	return &oas.Status{
		//		ID:       string(id),
		Name: request.Name,
	}, nil
}

// DeleteStatus implements oas.Handler.
func (h Handler) DeleteStatus(ctx context.Context, params oas.DeleteStatusParams) (oas.DeleteStatusRes, error) {
	q := `
	DELETE FROM statuses
	where id=$1
	`
	h.Db.MustExec(q, params.StatusId)

	return &oas.DeleteStatusOK{}, nil
}

// GetStatusById implements oas.Handler.
func (h Handler) GetStatusById(ctx context.Context, params oas.GetStatusByIdParams) (oas.GetStatusByIdRes, error) {
	rez := oas.Status{}
	err := h.Db.Get(&rez, "SELECT * FROM statuses WHERE id=$1", params.StatusId)
	if err != nil {
		return &oas.GetStatusByIdBadRequest{}, fmt.Errorf("failed to get incident by id")
	}
	return &rez, nil
}

// GetStatuses implements oas.Handler.
func (h Handler) GetStatuses(ctx context.Context) (oas.GetStatusesRes, error) {
	req := make([]oas.Status, 0, 50)

	incident := oas.Status{}
	rows, err := h.Db.Queryx("SELECT * FROM statuses")
	if err != nil {
		return &oas.GetStatusesBadRequest{}, fmt.Errorf("failed to get incidents")
	}
	for rows.Next() {
		err := rows.StructScan(&incident)
		if err != nil {
			log.Fatalln(err)
		}
		req = append(req, incident)
	}

	return &oas.StatususResponse{TrackingId: "123", Status: "Ok", Data: req}, nil
}

// UpdateStatus implements oas.Handler.
func (h Handler) UpdateStatus(ctx context.Context, req oas.UpdateStatusReq) (oas.UpdateStatusRes, error) {
	request, ok := req.(*oas.UpdateStatusApplicationJSON)
	if !ok {
		return &oas.UpdateStatusBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	q := `
	UPDATE statuses SET id=$1, name=$2
	where id=$1
	`
	rez := h.Db.MustExec(q, request.ID, request.Name)

	_, err := rez.RowsAffected()

	if err != nil {
		return &oas.UpdateStatusBadRequest{}, fmt.Errorf("failed to get last insert id")
	}

	return &oas.Status{
		ID:   request.ID,
		Name: request.Name,
	}, nil
}
