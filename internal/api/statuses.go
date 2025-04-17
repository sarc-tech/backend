package api

import (
	"context"
	"fmt"

	"github.com/sarc-tech/backend/internal/oas"
	"github.com/sarc-tech/backend/internal/utils"
)

// GetStatuses implements oas.Handler.
func (h Handler) GetStatuses(ctx context.Context) (oas.GetStatusesRes, error) {
	statuses, err := h.repo.GetStatuses()
	if err != nil {
		return nil, fmt.Errorf("failed to get statuses")
	}

	res := make([]oas.Status, len(statuses), len(statuses))
	for i, status := range statuses {
		res[i] = *status.MapStatusToAPi()
	}

	return &oas.StatusesResponse{TrackingId: utils.GetTraceID(ctx), Data: res}, nil
}
