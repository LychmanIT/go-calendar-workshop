package endpoints

import (
	"calendarWorkshop/models"
)

type LoginRequest struct {
	Credentials models.Auth `json:"credentials"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type LogoutRequest struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type LogoutResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type ChangeTimezoneRequest struct {
	Token    string `json:"token"`
	Timezone string `json:"timezone"`
	Err      string `json:"err,omitempty"`
}

type ChangeTimezoneResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type IndexEventRequest struct {
	Filters []models.Filter `json:"filters,omitempty"`
	Err     string          `json:"err,omitempty"`
}

type IndexEventResponse struct {
	Events []models.Event `json:"events"`
	Err    string         `json:"err,omitempty"`
}

type StoreEventRequest struct {
	Event *models.Event `json:"event"`
	Err   string        `json:"err,omitempty"`
}

type StoreEventResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type ShowEventRequest struct {
	EventID string `json:"id"`
	Err     string `json:"err,omitempty"`
}

type ShowEventResponse struct {
	Event *models.Event `json:"event"`
	Err   string        `json:"err,omitempty"`
}

type UpdateEventRequest struct {
	EventID string        `json:"id"`
	Event   *models.Event `json:"event"`
	Err     string        `json:"err,omitempty"`
}

type UpdateEventResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type DeleteEventRequest struct {
	EventID string `json:"id"`
	Err     string `json:"err,omitempty"`
}

type DeleteEventResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}
