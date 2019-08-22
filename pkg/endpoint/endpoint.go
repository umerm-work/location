package endpoint

import (
	"context"
	"github.com/fb-investments/location/pkg/io"
	"github.com/fb-investments/location/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"fmt"
)

// ListCountriesRequest collects the request parameters for the ListCountries method.
type ListCountriesRequest struct{}

// ListCountriesResponse collects the response parameters for the ListCountries method.
type ListCountriesResponse struct {
	io.Response
}

// MakeListCountriesEndpoint returns an endpoint that invokes ListCountries on the service.
func MakeListCountriesEndpoint(s service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res := s.ListCountries(ctx)
		return ListCountriesResponse{res}, nil
	}
}

// CreateCountriesRequest collects the request parameters for the CreateCountries method.
type CreateCountriesRequest struct {
	C io.Country `json:"payload"`
}

// CreateCountriesResponse collects the response parameters for the CreateCountries method.
type CreateCountriesResponse struct {
	io.Response
}

// MakeCreateCountriesEndpoint returns an endpoint that invokes CreateCountries on the service.
func MakeCreateCountriesEndpoint(s service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("asdasda")
		req := request.(CreateCountriesRequest)
		res := s.CreateCountries(ctx, req.C)
		return CreateCountriesResponse{res}, nil
	}
}

// ListStatesRequest collects the request parameters for the ListStates method.
type ListStatesRequest struct {
	CountryUid io.UID `json:"payload"`
}

// ListStatesResponse collects the response parameters for the ListStates method.
type ListStatesResponse struct {
	io.Response
}

// MakeListStatesEndpoint returns an endpoint that invokes ListStates on the service.
func MakeListStatesEndpoint(s service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListStatesRequest)
		res := s.ListStates(ctx, req.CountryUid)
		return ListStatesResponse{res}, nil
	}
}

// CreateStatesRequest collects the request parameters for the CreateStates method.
type CreateStatesRequest struct {
	C io.State `json:"payload"`
}

// CreateStatesResponse collects the response parameters for the CreateStates method.
type CreateStatesResponse struct {
	io.Response
}

// MakeCreateStatesEndpoint returns an endpoint that invokes CreateStates on the service.
func MakeCreateStatesEndpoint(s service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateStatesRequest)
		res := s.CreateStates(ctx, req.C)
		return CreateStatesResponse{res}, nil
	}
}

// ListCityRequest collects the request parameters for the ListCity method.
type ListCityRequest struct {
	StateUid io.UID `json:"payload"`
}

// ListCityResponse collects the response parameters for the ListCity method.
type ListCityResponse struct {
	io.Response
}

// MakeListCityEndpoint returns an endpoint that invokes ListCity on the service.
func MakeListCityEndpoint(s service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListCityRequest)
		res := s.ListCity(ctx, req.StateUid)
		return ListCityResponse{ res}, nil
	}
}

// CreateCityRequest collects the request parameters for the CreateCity method.
type CreateCityRequest struct {
	C io.City `json:"payload"`
}

// CreateCityResponse collects the response parameters for the CreateCity method.
type CreateCityResponse struct {
	io.Response
}

// MakeCreateCityEndpoint returns an endpoint that invokes CreateCity on the service.
func MakeCreateCityEndpoint(s service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCityRequest)
		res := s.CreateCity(ctx, req.C)
		return CreateCityResponse{ res}, nil
	}
}

// ListCountries implements Service. Primarily useful in a client.
func (e Endpoints) ListCountries(ctx context.Context) (res io.Response) {
	request := ListCountriesRequest{}
	response, err := e.ListCountriesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Response)
}

// CreateCountries implements Service. Primarily useful in a client.
func (e Endpoints) CreateCountries(ctx context.Context, c io.Country) (res io.Response) {
	request := CreateCountriesRequest{C: c}
	response, err := e.CreateCountriesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Response)
}

// ListStates implements Service. Primarily useful in a client.
func (e Endpoints) ListStates(ctx context.Context, countryUid io.UID) (res io.Response) {
	request := ListStatesRequest{CountryUid: countryUid}
	response, err := e.ListStatesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Response)
}

// CreateStates implements Service. Primarily useful in a client.
func (e Endpoints) CreateStates(ctx context.Context, c io.State) (res io.Response) {
	request := CreateStatesRequest{C: c}
	response, err := e.CreateStatesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Response)
}

// ListCity implements Service. Primarily useful in a client.
func (e Endpoints) ListCity(ctx context.Context, stateUid io.UID) (res io.Response) {
	request := ListCityRequest{StateUid: stateUid}
	response, err := e.ListCityEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Response)
}

// CreateCity implements Service. Primarily useful in a client.
func (e Endpoints) CreateCity(ctx context.Context, c io.City) (res io.Response) {
	request := CreateCityRequest{C: c}
	response, err := e.CreateCityEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(io.Response)
}
