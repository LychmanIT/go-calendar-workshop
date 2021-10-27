package transport

import (
	"calendarWorkshop/internal/domain/calendar"
	"calendarWorkshop/internal/util"
	"calendarWorkshop/pkg/calendar/endpoints"
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/index").Handler(httptransport.NewServer(
		ep.IndexEventEndpoint,
		decodeHTTPIndexEventRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/store").Handler(httptransport.NewServer(
		ep.StoreEventEndpoint,
		decodeHTTPStoreEventRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/show/{id}").Handler(httptransport.NewServer(
		ep.ShowEventEndpoint,
		decodeHTTPShowEventRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/update/{id}").Handler(httptransport.NewServer(
		ep.UpdateEventEndpoint,
		decodeHTTPUpdateEventRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/delete/{id}").Handler(httptransport.NewServer(
		ep.DeleteEventEndpoint,
		decodeHTTPDeleteEventRequest,
		encodeResponse,
	))

	return r
}

func decodeHTTPIndexEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.IndexEventRequest
	form := r.ParseForm()
	if form != nil {
		return nil, form
	}
	req.Filters = calendar.MapToFilters(r.Form)
	return req, nil
}

func decodeHTTPStoreEventRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.StoreEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPShowEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoints.ShowEventRequest{EventID: id}, nil
}

func decodeHTTPUpdateEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateEventRequest
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

func decodeHTTPDeleteEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
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
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
