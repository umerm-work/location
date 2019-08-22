// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "github.com/fb-investments/location/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	ListCountriesEndpoint   endpoint.Endpoint
	CreateCountriesEndpoint endpoint.Endpoint
	ListStatesEndpoint      endpoint.Endpoint
	CreateStatesEndpoint    endpoint.Endpoint
	ListCityEndpoint        endpoint.Endpoint
	CreateCityEndpoint      endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.LocationService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateCityEndpoint:      MakeCreateCityEndpoint(s),
		CreateCountriesEndpoint: MakeCreateCountriesEndpoint(s),
		CreateStatesEndpoint:    MakeCreateStatesEndpoint(s),
		ListCityEndpoint:        MakeListCityEndpoint(s),
		ListCountriesEndpoint:   MakeListCountriesEndpoint(s),
		ListStatesEndpoint:      MakeListStatesEndpoint(s),
	}
	for _, m := range mdw["ListCountries"] {
		eps.ListCountriesEndpoint = m(eps.ListCountriesEndpoint)
	}
	for _, m := range mdw["CreateCountries"] {
		eps.CreateCountriesEndpoint = m(eps.CreateCountriesEndpoint)
	}
	for _, m := range mdw["ListStates"] {
		eps.ListStatesEndpoint = m(eps.ListStatesEndpoint)
	}
	for _, m := range mdw["CreateStates"] {
		eps.CreateStatesEndpoint = m(eps.CreateStatesEndpoint)
	}
	for _, m := range mdw["ListCity"] {
		eps.ListCityEndpoint = m(eps.ListCityEndpoint)
	}
	for _, m := range mdw["CreateCity"] {
		eps.CreateCityEndpoint = m(eps.CreateCityEndpoint)
	}
	return eps
}
