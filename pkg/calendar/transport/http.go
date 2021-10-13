package transport

import (
	"calendarWorkshop/internal/util"
	"calendarWorkshop/pkg/calendar/endpoints"
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/status", httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))
	m.Handle("/index", httptransport.NewServer(
		ep.IndexEventEndpoint,
		decodeHTTPIndexEventRequest,
		encodeResponse,
	))
	m.Handle("/store", httptransport.NewServer(
		ep.StoreEventEndpoint,
		decodeHTTPStoreEventRequest,
		encodeResponse,
	))
	m.Handle("/show", httptransport.NewServer(
		ep.ShowEventEndpoint,
		decodeHTTPShowEventRequest,
		encodeResponse,
	))
	m.Handle("/update", httptransport.NewServer(
		ep.UpdateEventEndpoint,
		decodeHTTPUpdateEventRequest,
		encodeResponse,
	))
	m.Handle("/delete", httptransport.NewServer(
		ep.DeleteEventEndpoint,
		decodeHTTPDeleteEventRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPIndexEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.IndexEventRequest
	if r.ContentLength == 0 {
		return req, nil
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
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
	var req endpoints.ShowEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPUpdateEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPDeleteEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
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
