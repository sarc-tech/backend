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
		Region:     request.Region.Value,
		FIO:        request.Fio.Value,
		Status:     request.Status.Value,
		Date:       request.Date.Value,
	}

	if err := h.incidents.Add(ctx, incident); err != nil {
		return nil, fmt.Errorf("failed to add incident: %w", err)
	}

	return &oas.Incidents{
		ID:     oas.NewOptString(fmt.Sprintf("%d", incident.ID)),
		Region: oas.NewOptString(incident.Region),
		Fio:    oas.NewOptString(incident.FIO),
		Status: oas.NewOptString(incident.Status),
		Date:   oas.NewOptString(incident.Date),
	}, nil //? Не понятно что вернуть
}

func (h *handler) DeleteIncident(ctx context.Context, params oas.DeleteIncidentParams) error {
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
	return &oas.Incidents{
		ID:     oas.NewOptString(fmt.Sprintf("%d", incident.ID)),
		Region: oas.NewOptString(incident.Region),
		Fio:    oas.NewOptString(incident.FIO),
		Status: oas.NewOptString(incident.Status),
		Date:   oas.NewOptString(incident.Date),
	}, nil
}

func (h *handler) GetIncidents(ctx context.Context) (oas.GetIncidentsRes, error) {
	return nil, nil
}

func (h *handler) UpdateIncidents(ctx context.Context, req oas.UpdateIncidentsReq) (oas.UpdateIncidentsRes, error) {
	return nil, nil
}
