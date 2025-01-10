package api

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sarc-tech/backend/internal/models"
	"github.com/sarc-tech/backend/internal/oas"
)

// Compile-time check for Handler.
var _ oas.Handler = (*handler)(nil)

type handler struct {
	incidents IncidentUsecase
}

func NewHandler(incidents IncidentUsecase) *handler {
	return &handler{
		incidents: incidents,
	}
}

func (h *handler) AddIncidents(ctx context.Context, req oas.AddIncidentsReq) (oas.AddIncidentsRes, error) {
	request, ok := req.(*oas.AddIncidentsApplicationJSON)
	if !ok {
		return &oas.AddIncidentsBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	incident := &models.Incident{
		IsTraining: false, //! Нет поля в oas
		Region:     request.Region,
		FIO:        request.Fio,
		Status:     request.Status,
		Date:       request.Date,
	}

	if err := h.incidents.Add(ctx, incident); err != nil {
		return nil, fmt.Errorf("failed to add incident: %w", err)
	}

	return &oas.Incident{
		ID:     fmt.Sprintf("%d", incident.ID),
		Region: incident.Region,
		Fio:    incident.FIO,
		Status: incident.Status,
		Date:   incident.Date,
	}, nil //? Не понятно что вернуть
}

func (h *handler) DeleteIncident(ctx context.Context, params oas.DeleteIncidentParams) error {
	id, err := strconv.Atoi(params.IncidentId)
	if err != nil {
		return fmt.Errorf("invalid request type: %T", params)
	}
	err = h.incidents.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get incident: %w", err)
	}
	return nil
}

func (h *handler) GetIncidentById(
	ctx context.Context,
	params oas.GetIncidentByIdParams,
) (oas.GetIncidentByIdRes, error) {
	id, err := strconv.Atoi(params.IncidentId)
	if err != nil {
		return &oas.GetIncidentByIdBadRequest{}, fmt.Errorf("invalid request type: %T", params)
	}
	incident, err := h.incidents.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get incident: %w", err)
	}
	return &oas.Incident{
		ID:     fmt.Sprintf("%d", incident.ID),
		Region: incident.Region,
		Fio:    incident.FIO,
		Status: incident.Status,
		Date:   incident.Date,
	}, nil
}

func (h *handler) GetIncidents(ctx context.Context) (oas.GetIncidentsRes, error) {

	incidents, err := h.incidents.List(ctx, 50, 0)

	if err != nil {
		return nil, fmt.Errorf("failed to add incident: %w", err)
	}

	req := make([]oas.Incident, 0, 50)
	for _, val := range incidents {
		req = append(req, oas.Incident{
			ID:     fmt.Sprintf("%d", val.ID),
			Region: val.Region,
			Fio:    val.FIO,
			Status: val.Status,
			Date:   val.Date,
		})

	}

	return &oas.IncidentsResponse{Data: req}, nil
}

func (h *handler) UpdateIncidents(ctx context.Context, req oas.UpdateIncidentsReq) (oas.UpdateIncidentsRes, error) {
	request, ok := req.(*oas.UpdateIncidentsApplicationJSON)
	if !ok {
		return &oas.UpdateIncidentsBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	incident := &models.Incident{
		IsTraining: false, //! Нет поля в oas
		Region:     request.Region,
		FIO:        request.Fio,
		Status:     request.Status,
		Date:       request.Date,
	}

	if err := h.incidents.Update(ctx, incident); err != nil {
		return nil, fmt.Errorf("failed to add incident: %w", err)
	}

	return &oas.Incident{
		ID:     fmt.Sprintf("%d", incident.ID),
		Region: incident.Region,
		Fio:    incident.FIO,
		Status: incident.Status,
		Date:   incident.Date,
	}, nil //? Не понятно что вернуть
}
