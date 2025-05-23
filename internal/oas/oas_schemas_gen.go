// Code generated by ogen, DO NOT EDIT.

package oas

// AddIncidentBadRequest is response for AddIncident operation.
type AddIncidentBadRequest struct{}

func (*AddIncidentBadRequest) addIncidentRes() {}

// AddIncidentUnprocessableEntity is response for AddIncident operation.
type AddIncidentUnprocessableEntity struct{}

func (*AddIncidentUnprocessableEntity) addIncidentRes() {}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// DeleteIncidentBadRequest is response for DeleteIncident operation.
type DeleteIncidentBadRequest struct{}

func (*DeleteIncidentBadRequest) deleteIncidentRes() {}

// DeleteIncidentOK is response for DeleteIncident operation.
type DeleteIncidentOK struct{}

func (*DeleteIncidentOK) deleteIncidentRes() {}

// GetIncidentByIdBadRequest is response for GetIncidentById operation.
type GetIncidentByIdBadRequest struct{}

func (*GetIncidentByIdBadRequest) getIncidentByIdRes() {}

// GetIncidentByIdNotFound is response for GetIncidentById operation.
type GetIncidentByIdNotFound struct{}

func (*GetIncidentByIdNotFound) getIncidentByIdRes() {}

// GetIncidentsBadRequest is response for GetIncidents operation.
type GetIncidentsBadRequest struct{}

func (*GetIncidentsBadRequest) getIncidentsRes() {}

// GetStatusesBadRequest is response for GetStatuses operation.
type GetStatusesBadRequest struct{}

func (*GetStatusesBadRequest) getStatusesRes() {}

// Ref: #/components/schemas/Incident
type Incident struct {
	// Incident id.
	ID       OptInt `json:"id"`
	Region   string `json:"region"`
	Fio      string `json:"fio"`
	StatusId string `json:"statusId"`
	Date     string `json:"date"`
}

// GetID returns the value of ID.
func (s *Incident) GetID() OptInt {
	return s.ID
}

// GetRegion returns the value of Region.
func (s *Incident) GetRegion() string {
	return s.Region
}

// GetFio returns the value of Fio.
func (s *Incident) GetFio() string {
	return s.Fio
}

// GetStatusId returns the value of StatusId.
func (s *Incident) GetStatusId() string {
	return s.StatusId
}

// GetDate returns the value of Date.
func (s *Incident) GetDate() string {
	return s.Date
}

// SetID sets the value of ID.
func (s *Incident) SetID(val OptInt) {
	s.ID = val
}

// SetRegion sets the value of Region.
func (s *Incident) SetRegion(val string) {
	s.Region = val
}

// SetFio sets the value of Fio.
func (s *Incident) SetFio(val string) {
	s.Fio = val
}

// SetStatusId sets the value of StatusId.
func (s *Incident) SetStatusId(val string) {
	s.StatusId = val
}

// SetDate sets the value of Date.
func (s *Incident) SetDate(val string) {
	s.Date = val
}

// Ref: #/components/schemas/IncidentResponse
type IncidentResponse struct {
	TrackingId string   `json:"trackingId"`
	Status     string   `json:"status"`
	Data       Incident `json:"data"`
}

// GetTrackingId returns the value of TrackingId.
func (s *IncidentResponse) GetTrackingId() string {
	return s.TrackingId
}

// GetStatus returns the value of Status.
func (s *IncidentResponse) GetStatus() string {
	return s.Status
}

// GetData returns the value of Data.
func (s *IncidentResponse) GetData() Incident {
	return s.Data
}

// SetTrackingId sets the value of TrackingId.
func (s *IncidentResponse) SetTrackingId(val string) {
	s.TrackingId = val
}

// SetStatus sets the value of Status.
func (s *IncidentResponse) SetStatus(val string) {
	s.Status = val
}

// SetData sets the value of Data.
func (s *IncidentResponse) SetData(val Incident) {
	s.Data = val
}

func (*IncidentResponse) addIncidentRes()     {}
func (*IncidentResponse) getIncidentByIdRes() {}
func (*IncidentResponse) updateIncidentRes()  {}

// Ref: #/components/schemas/IncidentsResponse
type IncidentsResponse struct {
	TrackingId string     `json:"trackingId"`
	Status     string     `json:"status"`
	Data       []Incident `json:"data"`
}

// GetTrackingId returns the value of TrackingId.
func (s *IncidentsResponse) GetTrackingId() string {
	return s.TrackingId
}

// GetStatus returns the value of Status.
func (s *IncidentsResponse) GetStatus() string {
	return s.Status
}

// GetData returns the value of Data.
func (s *IncidentsResponse) GetData() []Incident {
	return s.Data
}

// SetTrackingId sets the value of TrackingId.
func (s *IncidentsResponse) SetTrackingId(val string) {
	s.TrackingId = val
}

// SetStatus sets the value of Status.
func (s *IncidentsResponse) SetStatus(val string) {
	s.Status = val
}

// SetData sets the value of Data.
func (s *IncidentsResponse) SetData(val []Incident) {
	s.Data = val
}

func (*IncidentsResponse) getIncidentsRes() {}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Status
type Status struct {
	// Status id.
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *Status) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *Status) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *Status) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Status) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/StatusesResponse
type StatusesResponse struct {
	TrackingId string   `json:"trackingId"`
	Status     string   `json:"status"`
	Data       []Status `json:"data"`
}

// GetTrackingId returns the value of TrackingId.
func (s *StatusesResponse) GetTrackingId() string {
	return s.TrackingId
}

// GetStatus returns the value of Status.
func (s *StatusesResponse) GetStatus() string {
	return s.Status
}

// GetData returns the value of Data.
func (s *StatusesResponse) GetData() []Status {
	return s.Data
}

// SetTrackingId sets the value of TrackingId.
func (s *StatusesResponse) SetTrackingId(val string) {
	s.TrackingId = val
}

// SetStatus sets the value of Status.
func (s *StatusesResponse) SetStatus(val string) {
	s.Status = val
}

// SetData sets the value of Data.
func (s *StatusesResponse) SetData(val []Status) {
	s.Data = val
}

func (*StatusesResponse) getStatusesRes() {}

// UpdateIncidentBadRequest is response for UpdateIncident operation.
type UpdateIncidentBadRequest struct{}

func (*UpdateIncidentBadRequest) updateIncidentRes() {}

// UpdateIncidentNotFound is response for UpdateIncident operation.
type UpdateIncidentNotFound struct{}

func (*UpdateIncidentNotFound) updateIncidentRes() {}

// UpdateIncidentUnprocessableEntity is response for UpdateIncident operation.
type UpdateIncidentUnprocessableEntity struct{}

func (*UpdateIncidentUnprocessableEntity) updateIncidentRes() {}
