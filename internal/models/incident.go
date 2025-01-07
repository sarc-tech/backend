package models

type Incident struct {
	ID         int
	IsTraining bool
	Region     string
	FIO        string
	Status     string
	Date       string
}
