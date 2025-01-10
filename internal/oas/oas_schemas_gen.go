// Code generated by ogen, DO NOT EDIT.

package oas

type AddIncidentsApplicationJSON Incident

func (*AddIncidentsApplicationJSON) addIncidentsReq() {}

type AddIncidentsApplicationXWwwFormUrlencoded Incident

func (*AddIncidentsApplicationXWwwFormUrlencoded) addIncidentsReq() {}

// AddIncidentsBadRequest is response for AddIncidents operation.
type AddIncidentsBadRequest struct{}

func (*AddIncidentsBadRequest) addIncidentsRes() {}

// AddIncidentsUnprocessableEntity is response for AddIncidents operation.
type AddIncidentsUnprocessableEntity struct{}

func (*AddIncidentsUnprocessableEntity) addIncidentsRes() {}

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

// Ref: #/components/schemas/Incident
type Incident struct {
	// Incident id.
	ID     string `json:"id"`
	Region string `json:"region"`
	Fio    string `json:"fio"`
	Status string `json:"status"`
	Date   string `json:"date"`
}

// GetID returns the value of ID.
func (s *Incident) GetID() string {
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

// GetStatus returns the value of Status.
func (s *Incident) GetStatus() string {
	return s.Status
}

// GetDate returns the value of Date.
func (s *Incident) GetDate() string {
	return s.Date
}

// SetID sets the value of ID.
func (s *Incident) SetID(val string) {
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

// SetStatus sets the value of Status.
func (s *Incident) SetStatus(val string) {
	s.Status = val
}

// SetDate sets the value of Date.
func (s *Incident) SetDate(val string) {
	s.Date = val
}

func (*Incident) addIncidentsRes()    {}
func (*Incident) getIncidentByIdRes() {}
func (*Incident) updateIncidentsRes() {}

// Ref: #/components/schemas/IncidentsResponse
type IncidentsResponse struct {
	TrackingId OptString  `json:"trackingId"`
	Status     OptString  `json:"status"`
	Data       []Incident `json:"data"`
}

// GetTrackingId returns the value of TrackingId.
func (s *IncidentsResponse) GetTrackingId() OptString {
	return s.TrackingId
}

// GetStatus returns the value of Status.
func (s *IncidentsResponse) GetStatus() OptString {
	return s.Status
}

// GetData returns the value of Data.
func (s *IncidentsResponse) GetData() []Incident {
	return s.Data
}

// SetTrackingId sets the value of TrackingId.
func (s *IncidentsResponse) SetTrackingId(val OptString) {
	s.TrackingId = val
}

// SetStatus sets the value of Status.
func (s *IncidentsResponse) SetStatus(val OptString) {
	s.Status = val
}

// SetData sets the value of Data.
func (s *IncidentsResponse) SetData(val []Incident) {
	s.Data = val
}

func (*IncidentsResponse) getIncidentsRes() {}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type UpdateIncidentsApplicationJSON Incident

func (*UpdateIncidentsApplicationJSON) updateIncidentsReq() {}

type UpdateIncidentsApplicationXWwwFormUrlencoded Incident

func (*UpdateIncidentsApplicationXWwwFormUrlencoded) updateIncidentsReq() {}

// UpdateIncidentsBadRequest is response for UpdateIncidents operation.
type UpdateIncidentsBadRequest struct{}

func (*UpdateIncidentsBadRequest) updateIncidentsRes() {}

// UpdateIncidentsNotFound is response for UpdateIncidents operation.
type UpdateIncidentsNotFound struct{}

func (*UpdateIncidentsNotFound) updateIncidentsRes() {}

// UpdateIncidentsUnprocessableEntity is response for UpdateIncidents operation.
type UpdateIncidentsUnprocessableEntity struct{}

func (*UpdateIncidentsUnprocessableEntity) updateIncidentsRes() {}
