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
	statuses  StatusesHandler
	users     UsersHandler
}

// AddStatus implements oas.Handler.
func (h *handler) AddStatus(ctx context.Context, req oas.AddStatusReq) (oas.AddStatusRes, error) {
	request, ok := req.(*oas.AddStatusApplicationJSON)
	if !ok {
		return &oas.AddStatusBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	statues := &models.Statues{

		Name: request.Name,
	}

	// TODO Надо переделать вызов
	//	if err := h.statues.Add(ctx, statues); err != nil {
	//		return nil, fmt.Errorf("failed to add incident: %w", err)
	//	}

	return &oas.Status{
		ID:   fmt.Sprintf("%d", statues.ID),
		Name: statues.Name,
	}, nil //? Не понятно что вернуть
}

// CheckSms implements oas.Handler.
func (h *handler) CheckSms(ctx context.Context, params oas.CheckSmsParams) (oas.CheckSmsRes, error) {
	panic("unimplemented")
}

// DeleteStatus implements oas.Handler.
func (h *handler) DeleteStatus(ctx context.Context, params oas.DeleteStatusParams) (oas.DeleteStatusRes, error) {
	panic("unimplemented")
}

// GetStatusById implements oas.Handler.
func (h *handler) GetStatusById(ctx context.Context, params oas.GetStatusByIdParams) (oas.GetStatusByIdRes, error) {
	panic("unimplemented")
}

// GetStatuses implements oas.Handler.
func (h *handler) GetStatuses(ctx context.Context) (oas.GetStatusesRes, error) {
	panic("unimplemented")
}

// GetUser implements oas.Handler.
func (h *handler) GetUser(ctx context.Context) (oas.GetUserRes, error) {
	panic("unimplemented")
}

// Logout implements oas.Handler.
func (h *handler) Logout(ctx context.Context) (oas.LogoutRes, error) {
	panic("unimplemented")
}

// SendSms implements oas.Handler.
func (h *handler) SendSms(ctx context.Context, params oas.SendSmsParams) (oas.SendSmsRes, error) {
	panic("unimplemented")
}

// UpdateStatus implements oas.Handler.
func (h *handler) UpdateStatus(ctx context.Context, req oas.UpdateStatusReq) (oas.UpdateStatusRes, error) {
	panic("unimplemented")
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
		StatusId:   request.StatusId,
		Date:       request.Date,
	}

	if err := h.incidents.Add(ctx, incident); err != nil {
		return nil, fmt.Errorf("failed to add incident: %w", err)
	}

	return &oas.Incident{
		ID:       fmt.Sprintf("%d", incident.ID),
		Region:   incident.Region,
		Fio:      incident.FIO,
		StatusId: incident.StatusId,
		Date:     incident.Date,
	}, nil //? Не понятно что вернуть
}

func (h *handler) DeleteIncident(ctx context.Context, params oas.DeleteIncidentParams) (oas.DeleteIncidentRes, error) {
	id, err := strconv.Atoi(params.IncidentId)
	if err != nil {
		return nil, fmt.Errorf("invalid request type: %T", params)
	}
	err = h.incidents.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get incident: %w", err)
	}
	return &oas.DeleteIncidentOK{}, nil
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
		ID:       fmt.Sprintf("%d", incident.ID),
		Region:   incident.Region,
		Fio:      incident.FIO,
		StatusId: incident.StatusId,
		Date:     incident.Date,
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
			ID:       fmt.Sprintf("%d", val.ID),
			Region:   val.Region,
			Fio:      val.FIO,
			StatusId: val.StatusId,
			Date:     val.Date,
		})

	}

	return &oas.IncidentsResponse{Data: req}, nil
}

func (h *handler) UpdateIncidents(ctx context.Context, req oas.UpdateIncidentsReq) (oas.UpdateIncidentsRes, error) {
	request, ok := req.(*oas.UpdateIncidentsApplicationJSON)
	if !ok {
		return &oas.UpdateIncidentsBadRequest{}, fmt.Errorf("invalid request type: %T", req)
	}

	id, err := strconv.Atoi(request.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid request type: %T", id)
	}

	incident := &models.Incident{
		ID:         id,
		IsTraining: false, //! Нет поля в oas
		Region:     request.Region,
		FIO:        request.Fio,
		StatusId:   request.StatusId,
		Date:       request.Date,
	}

	if err := h.incidents.Update(ctx, incident); err != nil {
		return nil, fmt.Errorf("failed to add incident: %w", err)
	}

	return &oas.Incident{
		ID:       fmt.Sprintf("%d", id),
		Region:   incident.Region,
		Fio:      incident.FIO,
		StatusId: incident.StatusId,
		Date:     incident.Date,
	}, nil //? Не понятно что вернуть
}
