package incident

import "github.com/sarc-tech/backend/internal/models"

func (r *repository) dto(incident *Model) *models.Incident {
	return &models.Incident{
		ID:         incident.ID,
		IsTraining: incident.IsTraining,
		Region:     incident.Region,
		FIO:        incident.FIO,
		Status:     incident.Status,
		Date:       incident.Date,
	}
}

func (r *repository) dtos(incidents []*Model) []*models.Incident {
	result := make([]*models.Incident, len(incidents))
	for i, incident := range incidents {
		result[i] = r.dto(incident)
	}
	return result
}

func (r *repository) orm(incident *models.Incident) *Model {
	return &Model{
		ID:         incident.ID,
		IsTraining: incident.IsTraining,
		Region:     incident.Region,
		FIO:        incident.FIO,
		Status:     incident.Status,
		Date:       incident.Date,
	}
}
