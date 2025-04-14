package models

import (
	"strconv"

	"github.com/sarc-tech/backend/internal/oas"
)

type Incident struct {
	ID       	int  	`db:"id"`
	Region     	string 	`db:"region"`
	FIO     	string 	`db:"fio"`
	StatusId 	int  	`db:"status_id"`
	Date   		string 	`db:"date"`
}

func (i *Incident) MapIncidentToAPi() *oas.Incident {
	return &oas.Incident{
		ID:        oas.NewOptInt(i.ID),
		Region:    i.Region,
		Fio:       i.FIO,
		StatusId:  strconv.Itoa(i.StatusId),
		Date:      i.Date,
	}
}

func MapIncidentFromAPi(apiIncident *oas.Incident) (incident *Incident) {
	status, _ := strconv.Atoi(apiIncident.StatusId)
	incident = &Incident{
		Region:    apiIncident.Region,
		FIO:       apiIncident.Fio,
		StatusId:  status,
		Date:      apiIncident.Date,
	}
	if v, ok := apiIncident.ID.Get(); ok {
		incident.ID = v
	}
	return
}

func (i *Incident) GetSqlUpdateMapString() string {
	return `
		region = :region,
		fio = :fio,
		status_id = :status_id, 
		date = :date
	`
}

func (u *Incident) GetSqlInsertMapString() string {
	return `
		(region, fio, status_id, date) VALUES 
		(:region, :fio, :status_id, :date)
	`
}

func (u *Incident) GetSqlReturningString() string {
	return `id, region, fio, status_id, date`
}