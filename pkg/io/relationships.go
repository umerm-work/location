package io


// Relationship structs
type CountryStates struct {
	UID
	HasState []State `json:"hasState" dgraph:"hasState"`
}
type StateCities struct {
	UID
	HasCities []City `json:"hasCities" dgraph:"hasCities"`
}
