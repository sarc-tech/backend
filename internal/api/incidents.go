package api

import (
	"context"
	"fmt"

	"github.com/sarc-tech/backend/internal/oas"
	"github.com/sarc-tech/backend/internal/utils"
)

// AddIncidents implements oas.Handler.
func (h Handler) AddIncident(ctx context.Context, req *oas.Incident) (oas.AddIncidentRes, error) {
	inc, err := h.controller.CreateIncident(req)
	if err != nil {
		return nil, err
	}
	return &oas.IncidentResponse{TrackingId: utils.GetTraceID(ctx), Data: *inc.MapIncidentToAPi()}, nil
}

// DeleteIncident implements oas.Handler.
func (h Handler) DeleteIncident(ctx context.Context, params oas.DeleteIncidentParams) (oas.DeleteIncidentRes, error) {
	err := h.controller.DeleteIncident(params.IncidentId)
	if err != nil {
		return &oas.DeleteIncidentBadRequest{}, fmt.Errorf("failed to delete incident by id")
	}
	return &oas.DeleteIncidentOK{}, nil
}

// GetIncidentById implements oas.Handler.
func (h Handler) GetIncidentById(ctx context.Context, params oas.GetIncidentByIdParams) (oas.GetIncidentByIdRes, error) {
	inc, err := h.controller.GetIncidentById(params.IncidentId)
	if err != nil {
		return &oas.GetIncidentByIdBadRequest{}, fmt.Errorf("failed to get incident by id")
	}
	return &oas.IncidentResponse{TrackingId: utils.GetTraceID(ctx), Data: *inc.MapIncidentToAPi()}, nil
}

// GetIncidents implements oas.Handler.
func (h Handler) GetIncidents(ctx context.Context) (oas.GetIncidentsRes, error) {
	incidents, err := h.controller.GetIncidents()
	if err != nil {
		return nil, fmt.Errorf("failed to get incidents")
	}

	res := make([]oas.Incident, len(incidents), len(incidents))
	for i, inc := range incidents {
		res[i] = *inc.MapIncidentToAPi()
	}

	return &oas.IncidentsResponse{TrackingId: utils.GetTraceID(ctx), Data: res}, nil
}

// UpdateIncidents implements oas.Handler.
func (h Handler) UpdateIncident(ctx context.Context, req *oas.Incident) (oas.UpdateIncidentRes, error) {
	inc, err := h.controller.UpdateIncident(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update incident")
	}

	return &oas.IncidentResponse{TrackingId: utils.GetTraceID(ctx), Data: *inc.MapIncidentToAPi()}, nil
}
