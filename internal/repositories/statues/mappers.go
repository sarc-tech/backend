package incident

import "github.com/sarc-tech/backend/internal/models"

func (r *repository) dto(statues *Model) *models.Statues {
	return &models.Statues{
		ID:   statues.ID,
		Name: statues.name,
	}
}

func (r *repository) dtos(statues []*Model) []*models.Statues {
	result := make([]*models.Statues, len(statues))
	for i, statues := range statues {
		result[i] = r.dto(statues)
	}
	return result
}

func (r *repository) orm(statues *models.Statues) *Model {
	return &Model{
		ID:   statues.ID,
		name: statues.Name,
	}
}
