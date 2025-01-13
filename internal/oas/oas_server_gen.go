// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	IncidentsHandler
	StatusesHandler
	UsersHandler
}

// IncidentsHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Incidents
type IncidentsHandler interface {
	// AddIncidents implements addIncidents operation.
	//
	// Add a new incidents.
	//
	// POST /incidents
	AddIncidents(ctx context.Context, req AddIncidentsReq) (AddIncidentsRes, error)
	// DeleteIncident implements deleteIncident operation.
	//
	// Delete an incidents.
	//
	// DELETE /incidents/{incidentId}
	DeleteIncident(ctx context.Context, params DeleteIncidentParams) (DeleteIncidentRes, error)
	// GetIncidentById implements getIncidentById operation.
	//
	// Returns a single incidents.
	//
	// GET /incidents/{incidentId}
	GetIncidentById(ctx context.Context, params GetIncidentByIdParams) (GetIncidentByIdRes, error)
	// GetIncidents implements getIncidents operation.
	//
	// List of incidents.
	//
	// GET /incidents
	GetIncidents(ctx context.Context) (GetIncidentsRes, error)
	// UpdateIncidents implements updateIncidents operation.
	//
	// Update an existing incidents by Id.
	//
	// PUT /incidents
	UpdateIncidents(ctx context.Context, req UpdateIncidentsReq) (UpdateIncidentsRes, error)
}

// StatusesHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Statuses
type StatusesHandler interface {
	// AddStatus implements addStatus operation.
	//
	// Add a new status.
	//
	// POST /statuses
	AddStatus(ctx context.Context, req AddStatusReq) (AddStatusRes, error)
	// DeleteStatus implements deleteStatus operation.
	//
	// Delete an status.
	//
	// DELETE /statuses/{statusId}
	DeleteStatus(ctx context.Context, params DeleteStatusParams) (DeleteStatusRes, error)
	// GetStatusById implements getStatusById operation.
	//
	// Returns a single incidents.
	//
	// GET /statuses/{statusId}
	GetStatusById(ctx context.Context, params GetStatusByIdParams) (GetStatusByIdRes, error)
	// GetStatuses implements getStatuses operation.
	//
	// List of statuses.
	//
	// GET /statuses
	GetStatuses(ctx context.Context) (GetStatusesRes, error)
	// UpdateStatus implements updateStatus operation.
	//
	// Update an existing status by Id.
	//
	// PUT /statuses
	UpdateStatus(ctx context.Context, req UpdateStatusReq) (UpdateStatusRes, error)
}

// UsersHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Users
type UsersHandler interface {
	// CheckSms implements CheckSms operation.
	//
	// Returns a token.
	//
	// POST /checksms
	CheckSms(ctx context.Context, params CheckSmsParams) (CheckSmsRes, error)
	// GetUser implements getUser operation.
	//
	// Returns a user.
	//
	// GET /user
	GetUser(ctx context.Context) (GetUserRes, error)
	// Logout implements Logout operation.
	//
	// Удаляет сессию пользователя.
	//
	// POST /logout
	Logout(ctx context.Context) (LogoutRes, error)
	// SendSms implements SendSms operation.
	//
	// Returns a token.
	//
	// POST /sendsms/{phone}
	SendSms(ctx context.Context, params SendSmsParams) (SendSmsRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
