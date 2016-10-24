package service

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

func makeGeocodeEndpoint(svc GeocodeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(geocodeRequest)
		v, err := svc.Geocode(req.S)
		if err != nil {
			return geocodeResponse{v, err.Error()}, nil
		}
		return geocodeResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc GeocodeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func decodeGeocodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request geocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type geocodeRequest struct {
	S string `json:"placename"`
}

type geocodeResponse struct {
	V   string `json:"geocode_result"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}
