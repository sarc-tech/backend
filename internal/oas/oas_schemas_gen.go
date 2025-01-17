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

type AddStatusApplicationJSON Status

func (*AddStatusApplicationJSON) addStatusReq() {}

type AddStatusApplicationXWwwFormUrlencoded Status

func (*AddStatusApplicationXWwwFormUrlencoded) addStatusReq() {}

// AddStatusBadRequest is response for AddStatus operation.
type AddStatusBadRequest struct{}

func (*AddStatusBadRequest) addStatusRes() {}

// AddStatusUnprocessableEntity is response for AddStatus operation.
type AddStatusUnprocessableEntity struct{}

func (*AddStatusUnprocessableEntity) addStatusRes() {}

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

// CheckSmsBadRequest is response for CheckSms operation.
type CheckSmsBadRequest struct{}

func (*CheckSmsBadRequest) checkSmsRes() {}

// CheckSmsNotFound is response for CheckSms operation.
type CheckSmsNotFound struct{}

func (*CheckSmsNotFound) checkSmsRes() {}

// DeleteIncidentBadRequest is response for DeleteIncident operation.
type DeleteIncidentBadRequest struct{}

func (*DeleteIncidentBadRequest) deleteIncidentRes() {}

// DeleteIncidentOK is response for DeleteIncident operation.
type DeleteIncidentOK struct{}

func (*DeleteIncidentOK) deleteIncidentRes() {}

// DeleteStatusBadRequest is response for DeleteStatus operation.
type DeleteStatusBadRequest struct{}

func (*DeleteStatusBadRequest) deleteStatusRes() {}

// DeleteStatusOK is response for DeleteStatus operation.
type DeleteStatusOK struct{}

func (*DeleteStatusOK) deleteStatusRes() {}

// GetIncidentByIdBadRequest is response for GetIncidentById operation.
type GetIncidentByIdBadRequest struct{}

func (*GetIncidentByIdBadRequest) getIncidentByIdRes() {}

// GetIncidentByIdNotFound is response for GetIncidentById operation.
type GetIncidentByIdNotFound struct{}

func (*GetIncidentByIdNotFound) getIncidentByIdRes() {}

// GetIncidentsBadRequest is response for GetIncidents operation.
type GetIncidentsBadRequest struct{}

func (*GetIncidentsBadRequest) getIncidentsRes() {}

// GetStatusByIdBadRequest is response for GetStatusById operation.
type GetStatusByIdBadRequest struct{}

func (*GetStatusByIdBadRequest) getStatusByIdRes() {}

// GetStatusByIdNotFound is response for GetStatusById operation.
type GetStatusByIdNotFound struct{}

func (*GetStatusByIdNotFound) getStatusByIdRes() {}

// GetStatusesBadRequest is response for GetStatuses operation.
type GetStatusesBadRequest struct{}

func (*GetStatusesBadRequest) getStatusesRes() {}

// GetUserBadRequest is response for GetUser operation.
type GetUserBadRequest struct{}

func (*GetUserBadRequest) getUserRes() {}

// GetUserNotFound is response for GetUser operation.
type GetUserNotFound struct{}

func (*GetUserNotFound) getUserRes() {}

// Ref: #/components/schemas/Incident
type Incident struct {
	// Incident id.
	ID       string `json:"id"`
	Region   string `json:"region"`
	Fio      string `json:"fio"`
	StatusId string `json:"statusId"`
	Date     string `json:"date"`
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

// GetStatusId returns the value of StatusId.
func (s *Incident) GetStatusId() string {
	return s.StatusId
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

// SetStatusId sets the value of StatusId.
func (s *Incident) SetStatusId(val string) {
	s.StatusId = val
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
	TrackingId string     `json:"trackingId"`
	Status     string     `json:"status"`
	Data       []Incident `json:"data"`
	Statuses   []Status   `json:"statuses"`
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

// GetStatuses returns the value of Statuses.
func (s *IncidentsResponse) GetStatuses() []Status {
	return s.Statuses
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

// SetStatuses sets the value of Statuses.
func (s *IncidentsResponse) SetStatuses(val []Status) {
	s.Statuses = val
}

func (*IncidentsResponse) getIncidentsRes() {}

// LogoutBadRequest is response for Logout operation.
type LogoutBadRequest struct{}

func (*LogoutBadRequest) logoutRes() {}

// LogoutNotFound is response for Logout operation.
type LogoutNotFound struct{}

func (*LogoutNotFound) logoutRes() {}

// LogoutOK is response for Logout operation.
type LogoutOK struct{}

func (*LogoutOK) logoutRes() {}

// SendSmsBadRequest is response for SendSms operation.
type SendSmsBadRequest struct{}

func (*SendSmsBadRequest) sendSmsRes() {}

// SendSmsNotFound is response for SendSms operation.
type SendSmsNotFound struct{}

func (*SendSmsNotFound) sendSmsRes() {}

// SendSmsOK is response for SendSms operation.
type SendSmsOK struct{}

func (*SendSmsOK) sendSmsRes() {}

// Ref: #/components/schemas/Status
type Status struct {
	// Status id.
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *Status) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *Status) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *Status) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Status) SetName(val string) {
	s.Name = val
}

func (*Status) addStatusRes()     {}
func (*Status) getStatusByIdRes() {}
func (*Status) updateStatusRes()  {}

// Ref: #/components/schemas/StatususResponse
type StatususResponse struct {
	TrackingId string   `json:"trackingId"`
	Status     string   `json:"status"`
	Data       []Status `json:"data"`
}

// GetTrackingId returns the value of TrackingId.
func (s *StatususResponse) GetTrackingId() string {
	return s.TrackingId
}

// GetStatus returns the value of Status.
func (s *StatususResponse) GetStatus() string {
	return s.Status
}

// GetData returns the value of Data.
func (s *StatususResponse) GetData() []Status {
	return s.Data
}

// SetTrackingId sets the value of TrackingId.
func (s *StatususResponse) SetTrackingId(val string) {
	s.TrackingId = val
}

// SetStatus sets the value of Status.
func (s *StatususResponse) SetStatus(val string) {
	s.Status = val
}

// SetData sets the value of Data.
func (s *StatususResponse) SetData(val []Status) {
	s.Data = val
}

func (*StatususResponse) getStatusesRes() {}

// Ref: #/components/schemas/Token
type Token struct {
	// Token.
	Token string `json:"token"`
}

// GetToken returns the value of Token.
func (s *Token) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Token) SetToken(val string) {
	s.Token = val
}

func (*Token) checkSmsRes() {}

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

type UpdateStatusApplicationJSON Status

func (*UpdateStatusApplicationJSON) updateStatusReq() {}

type UpdateStatusApplicationXWwwFormUrlencoded Status

func (*UpdateStatusApplicationXWwwFormUrlencoded) updateStatusReq() {}

// UpdateStatusBadRequest is response for UpdateStatus operation.
type UpdateStatusBadRequest struct{}

func (*UpdateStatusBadRequest) updateStatusRes() {}

// UpdateStatusNotFound is response for UpdateStatus operation.
type UpdateStatusNotFound struct{}

func (*UpdateStatusNotFound) updateStatusRes() {}

// UpdateStatusUnprocessableEntity is response for UpdateStatus operation.
type UpdateStatusUnprocessableEntity struct{}

func (*UpdateStatusUnprocessableEntity) updateStatusRes() {}

// Ref: #/components/schemas/User
type User struct {
	// Incident id.
	ID   string `json:"id"`
	Fio  string `json:"fio"`
	Role string `json:"role"`
}

// GetID returns the value of ID.
func (s *User) GetID() string {
	return s.ID
}

// GetFio returns the value of Fio.
func (s *User) GetFio() string {
	return s.Fio
}

// GetRole returns the value of Role.
func (s *User) GetRole() string {
	return s.Role
}

// SetID sets the value of ID.
func (s *User) SetID(val string) {
	s.ID = val
}

// SetFio sets the value of Fio.
func (s *User) SetFio(val string) {
	s.Fio = val
}

// SetRole sets the value of Role.
func (s *User) SetRole(val string) {
	s.Role = val
}

func (*User) getUserRes() {}
