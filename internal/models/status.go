package models


import (
	"github.com/sarc-tech/backend/internal/oas"
)

type Status struct {
	ID       	int  	`db:"id"`
	Name     	string 	`db:"name"`
}

func (s *Status) MapStatusToAPi() *oas.Status {
	return &oas.Status{
		ID:      s.ID,
		Name:    s.Name,
	}
}
