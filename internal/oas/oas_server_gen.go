// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	IncidentsHandler
	StatusesHandler
}

// IncidentsHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Incidents
type IncidentsHandler interface {
	// AddIncident implements addIncident operation.
	//
	// Add a new incidents.
	//
	// POST /incidents
	AddIncident(ctx context.Context, req *Incident) (AddIncidentRes, error)
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
	// UpdateIncident implements updateIncident operation.
	//
	// Update an existing incidents by Id.
	//
	// PUT /incidents
	UpdateIncident(ctx context.Context, req *Incident) (UpdateIncidentRes, error)
}

// StatusesHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Statuses
type StatusesHandler interface {
	// GetStatuses implements getStatuses operation.
	//
	// Получение списка статусов.
	//
	// GET /statuses
	GetStatuses(ctx context.Context) (GetStatusesRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
