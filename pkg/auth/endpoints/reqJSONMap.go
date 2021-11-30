package endpoints

import "calendarWorkshop/models"

type LoginRequest struct {
	Credentials models.Auth `json:"credentials"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type LogoutRequest struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
	Status int32  `json:"status"`
	Err    string `json:"err,omitempty"`
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct {
	Err string `json:"err,omitempty"`
}

type ChangeTimezoneRequest struct {
	Token    string `json:"token"`
	Timezone string `json:"timezone"`
}

type ChangeTimezoneResponse struct {
	Status int32  `json:"status"`
	Err    string `json:"err,omitempty"`
}

type GetTimezoneRequest struct {
	Token string `json:"token"`
}

type GetTimezoneResponse struct {
	Timezone string `json:"timezone"`
	Err      string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Status int    `json:"status"`
	Err    string `json:"err,omitempty"`
}
