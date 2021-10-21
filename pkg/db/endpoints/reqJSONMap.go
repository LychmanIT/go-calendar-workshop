package endpoints

import (
	"calendarWorkshop/models"
)

type AllEventsRequest struct {
	Filters []models.Event `json:"filters,omitempty"`
}

type AllEventsResponse struct {
	Event []models.Event `json:"event"`
	Err   string         `json:"err,omitempty"`
}

type UpdateRequest struct {
	TicketID string             `json:"ticketID"`
	Document *internal.Document `json:"document"`
}

type UpdateResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}

type AddRequest struct {
	Document *internal.Document `json:"document"`
}

type AddResponse struct {
	TicketID string `json:"ticketID"`
	Err      string `json:"err"`
}

type RemoveRequest struct {
	TicketID string `json:"ticketID"`
}

type RemoveResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

//type ValidateRequest struct {
//	Document *internal.Document `json:"document"`
//}
//
//type ValidateResponse struct {
//	Valid bool   `json:"valid"`
//	Err   string `json:"err,omitempty"`
//}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}
