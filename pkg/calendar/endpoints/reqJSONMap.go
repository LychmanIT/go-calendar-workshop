package endpoints

import "calendarWorkshop/internal/domain/calendar"

type IndexEventRequest struct {
	Filters []calendar.Filter `json:"filters,omitempty"`
}

type IndexEventResponse struct {
	Events []calendar.Event `json:"events"`
	Err    string           `json:"err,omitempty"`
}

type StoreEventRequest struct {
	Event *calendar.Event `json:"event"`
}

type StoreEventResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}

type ShowEventRequest struct {
	EventID string `json:"id"`
}

type ShowEventResponse struct {
	Event *calendar.Event `json:"event"`
	Err   string          `json:"err,omitempty"`
}

type UpdateEventRequest struct {
	EventID string          `json:"id"`
	Event   *calendar.Event `json:"event"`
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
