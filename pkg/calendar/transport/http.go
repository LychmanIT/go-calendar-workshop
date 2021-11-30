package transport

import (
	authsvc "calendarWorkshop/api/v1/pb/auth"
	"calendarWorkshop/internal/util"
	"calendarWorkshop/pkg/calendar/endpoints"
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

var (
	ErrBadRouting = errors.New("bad routing")
	authClient    authsvc.AuthClient
)

func NewHTTPHandler(ep endpoints.Set, ac authsvc.AuthClient) http.Handler {
	authClient = ac

	r := mux.NewRouter()

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		ep.LoginEndpoint,
		decodeHTTPLoginRequest,
		encodeHTTPLoginResponse,
	))

	r.Methods("POST").Path("/logout").Handler(httptransport.NewServer(
		ep.LogoutEndpoint,
		decodeHTTPLogoutRequest,
		encodeHTTPLogoutResponse,
	))

	r.Methods("POST").Path("/change-timezone").Handler(httptransport.NewServer(
		ep.ChangeTimezoneEndpoint,
		decodeHTTPChangeTimezoneRequest,
		encodeHTTPChangeTimezoneResponse,
	))

	r.Methods("GET").Path("/api/status").Handler(httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/api/index").Handler(httptransport.NewServer(
		ep.IndexEventEndpoint,
		decodeHTTPIndexEventRequest,
		encodeHTTPIndexResponse,
	))

	r.Methods("POST").Path("/api/store").Handler(httptransport.NewServer(
		ep.StoreEventEndpoint,
		decodeHTTPStoreEventRequest,
		encodeHTTPStoreResponse,
	))

	r.Methods("GET").Path("/api/show/{id}").Handler(httptransport.NewServer(
		ep.ShowEventEndpoint,
		decodeHTTPShowEventRequest,
		encodeHTTPShowResponse,
	))

	r.Methods("PUT").Path("/api/update/{id}").Handler(httptransport.NewServer(
		ep.UpdateEventEndpoint,
		decodeHTTPUpdateEventRequest,
		encodeHTTPUpdateResponse,
	))

	r.Methods("DELETE").Path("/api/delete/{id}").Handler(httptransport.NewServer(
		ep.DeleteEventEndpoint,
		decodeHTTPDeleteEventRequest,
		encodeHTTPDeleteResponse,
	))

	r.Handle("/metrics", promhttp.Handler())

	return r
}

func decodeHTTPLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPLogoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LogoutRequest

	if err := verifyToken(ctx, r); err != nil {
		req.Err = err.Error()
		return req, nil
	}

	req.Token = strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	return req, nil
}

func decodeHTTPChangeTimezoneRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ChangeTimezoneRequest

	if err := verifyToken(ctx, r); err != nil {
		req.Err = err.Error()
		return req, nil
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Token = strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	return req, nil
}

func decodeHTTPIndexEventRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.IndexEventRequest

	if err := verifyToken(ctx, r); err != nil {
		req.Err = err.Error()
		return req, nil
	}

	form := r.ParseForm()
	if form != nil {
		return nil, form
	}
	req.Filters = util.MapToFilters(r.Form)
	return req, nil
}

func decodeHTTPStoreEventRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.StoreEventRequest

	if err := verifyToken(ctx, r); err != nil {
		req.Err = err.Error()
		return req, nil
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPShowEventRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ShowEventRequest

	if err := verifyToken(ctx, r); err != nil {
		req.Err = err.Error()
		return req, nil
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoints.ShowEventRequest{EventID: id}, nil
}

func decodeHTTPUpdateEventRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateEventRequest

	if err := verifyToken(ctx, r); err != nil {
		req.Err = err.Error()
		return req, nil
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return endpoints.UpdateEventRequest{
		EventID: id,
		Event:   req.Event,
	}, nil
}

func decodeHTTPDeleteEventRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteEventRequest

	if err := verifyToken(ctx, r); err != nil {
		req.Err = err.Error()
		return req, nil
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoints.DeleteEventRequest{EventID: id}, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	return req, nil
}

func encodeHTTPLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	authResp, ok := response.(endpoints.LoginResponse)
	if authResp.Err != "" || !ok {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
		return nil
	}

	w.Header().Set("X-TOKEN-GEN", authResp.Token)
	return json.NewEncoder(w).Encode(response)
}

func encodeHTTPLogoutResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authResp, ok := response.(endpoints.LogoutResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}

	switch authResp.Err {
	case "":
		return json.NewEncoder(w).Encode(response)
	case util.ErrUnauthorized.Error():
		w.WriteHeader(http.StatusForbidden)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
}

func encodeHTTPChangeTimezoneResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authResp, ok := response.(endpoints.ChangeTimezoneResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}

	switch authResp.Err {
	case "":
		return json.NewEncoder(w).Encode(response)
	case util.ErrUnauthorized.Error():
		w.WriteHeader(http.StatusForbidden)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
}

func encodeHTTPIndexResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authResp, ok := response.(endpoints.IndexEventResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}

	switch authResp.Err {
	case "":
		return json.NewEncoder(w).Encode(response)
	case util.ErrUnauthorized.Error():
		w.WriteHeader(http.StatusForbidden)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
}

func encodeHTTPStoreResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authResp, ok := response.(endpoints.StoreEventResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}

	switch authResp.Err {
	case "":
		return json.NewEncoder(w).Encode(response)
	case util.ErrUnauthorized.Error():
		w.WriteHeader(http.StatusForbidden)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
}

func encodeHTTPShowResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authResp, ok := response.(endpoints.ShowEventResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}

	switch authResp.Err {
	case "":
		return json.NewEncoder(w).Encode(response)
	case util.ErrEventNotFound.Error():
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	case util.ErrUnauthorized.Error():
		w.WriteHeader(http.StatusForbidden)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
}

func encodeHTTPUpdateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authResp, ok := response.(endpoints.UpdateEventResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
	if authResp.Status == 404 {
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": util.ErrEventNotFound.Error(),
		})
	}

	switch authResp.Err {
	case "":
		return json.NewEncoder(w).Encode(response)
	case util.ErrUnauthorized.Error():
		w.WriteHeader(http.StatusForbidden)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
}

func encodeHTTPDeleteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	authResp, ok := response.(endpoints.DeleteEventResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
	if authResp.Status == 404 {
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": util.ErrEventNotFound.Error(),
		})
	}

	switch authResp.Err {
	case "":
		return json.NewEncoder(w).Encode(response)
	case util.ErrEventNotFound.Error():
		w.WriteHeader(http.StatusNotFound)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	case util.ErrUnauthorized.Error():
		w.WriteHeader(http.StatusForbidden)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(map[string]interface{}{
			"error": authResp.Err,
		})
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	e, ok := response.(error)
	if ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrUnknownArgument:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	case util.ErrUserNotFound:
		w.WriteHeader(http.StatusUnauthorized)
	case util.ErrAlreadyLoggedIn:
		w.WriteHeader(http.StatusForbidden)
	case util.ErrUnauthorized:
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func getToken(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		return splitToken[1], nil
	}
	return "", util.ErrUnauthorized
}

func verifyToken(ctx context.Context, r *http.Request) error {
	token, err := getToken(r)
	if err != nil {
		return err
	}
	resp, err := authClient.AuthVerifyToken(ctx, &authsvc.AuthVerifyTokenRequest{Token: token})
	if err != nil {
		return err
	}
	if resp.Err != "" {
		return util.ErrUnauthorized
	}
	return nil
}
