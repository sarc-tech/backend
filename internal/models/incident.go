package models

type Incident struct {
	ID         int
	IsTraining bool
	Region     string
	FIO        string
	StatusId   string
	Date       string
}

type Statues struct {
	ID   int
	Name string
}
