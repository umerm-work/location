package http

import (
	"context"
	"encoding/json"
	endpoint "github.com/fb-investments/aone-location-service/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)
const (
	apiVersion = `/api/v1/`
	// Routes
	listCountriesRoute   = apiVersion + `location/countries`
	createCountriesRoute = apiVersion + `location/country`
	listStatesRoute      = apiVersion + `location/states`
	createStateRoute     = apiVersion + `location/state`
	listCitiesRoute      = apiVersion + `location/cities`
	createCitiesRoute    = apiVersion + `location/city`
)

// makeListCountriesHandler creates the handler logic
func makeListCountriesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path(listCountriesRoute).Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListCountriesEndpoint, decodeListCountriesRequest, encodeListCountriesResponse, options...)))
}

// decodeListCountriesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListCountriesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListCountriesRequest{}
	//err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
	//return req, err
}

// encodeListCountriesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListCountriesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateCountriesHandler creates the handler logic
func makeCreateCountriesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path(createCountriesRoute).Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateCountriesEndpoint, decodeCreateCountriesRequest, encodeCreateCountriesResponse, options...)))
}

// decodeCreateCountriesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateCountriesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateCountriesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateCountriesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateCountriesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListStatesHandler creates the handler logic
func makeListStatesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path(listStatesRoute).Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListStatesEndpoint, decodeListStatesRequest, encodeListStatesResponse, options...)))
}

// decodeListStatesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListStatesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListStatesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListStatesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListStatesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateStatesHandler creates the handler logic
func makeCreateStatesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path(createStateRoute).Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateStatesEndpoint, decodeCreateStatesRequest, encodeCreateStatesResponse, options...)))
}

// decodeCreateStatesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateStatesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateStatesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateStatesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateStatesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeListCityHandler creates the handler logic
func makeListCityHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path(listCitiesRoute).Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ListCityEndpoint, decodeListCityRequest, encodeListCityResponse, options...)))
}

// decodeListCityRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeListCityRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ListCityRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeListCityResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeListCityResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateCityHandler creates the handler logic
func makeCreateCityHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path(createCitiesRoute).Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateCityEndpoint, decodeCreateCityRequest, encodeCreateCityResponse, options...)))
}

// decodeCreateCityRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateCityRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateCityRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateCityResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateCityResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
