package endpoints

import "calendarWorkshop/models"

type GetUserRequest struct {
	Auth *models.Auth `json:"auth"`
}

type GetUserResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type AllEventsRequest struct {
	Filters []models.Filter `json:"filters,omitempty"`
}

type AllEventsResponse struct {
	Events []models.Event `json:"events"`
	Err    string         `json:"err,omitempty"`
}

type AddEventRequest struct {
	Event *models.Event `json:"event"`
}

type AddEventResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type ShowEventRequest struct {
	EventID string `json:"id"`
}

type ShowEventResponse struct {
	Event *models.Event `json:"event"`
	Err   string        `json:"err,omitempty"`
}

type UpdateEventRequest struct {
	EventID string        `json:"id"`
	Event   *models.Event `json:"event"`
}

type UpdateEventResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type DeleteEventRequest struct {
	EventID string `json:"id"`
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
