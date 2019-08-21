package io

import (
	"time"
)

type UID struct {
	Uid string `json:"uid,omitempty" dgraph:"uid,omitempty"`
}

type Country struct {
	UID
	CountryName        string    `json:"countryName,omitempty"  dgraph:"countryName,omitempty"  validate:"required"`
	CountryCode        string    `json:"countryCode,omitempty"  dgraph:"countryCode,omitempty" validate:"required"`
	Label              string    `json:"label,omitempty"  dgraph:"label,omitempty"`
	CountryPhoneCode   string    `json:"countryPhoneCode,omitempty"  dgraph:"countryPhoneCode,omitempty" validate:"required"`
	CurrencyCode       string    `json:"currencyCode,omitempty"  dgraph:"currencyCode,omitempty" validate:"required"`
	CountryUpdatedDate time.Time `json:"countryUpdatedDate,omitempty"  dgraph:"countryUpdatedDate,omitempty"`
	CountryCreatedDate time.Time `json:"countryCreatedDate,omitempty"  dgraph:"countryCreatedDate,omitempty"`
}
type State struct {
	UID
	CountryUid       string    `json:"countryUid,omitempty" dgraph:"countryUid,omitempty"`
	StateName        string    `json:"stateName,omitempty"  dgraph:"stateName,omitempty"  validate:"required"`
	StateCode        string    `json:"stateCode,omitempty"  dgraph:"stateCode,omitempty" validate:"required"`
	Label            string    `json:"label,omitempty"  dgraph:"label,omitempty"`
	StateUpdatedDate time.Time `json:"stateUpdatedDate,omitempty"  dgraph:"stateUpdatedDate,omitempty"`
	StateCreatedDate time.Time `json:"stateCreatedDate,omitempty"  dgraph:"stateCreatedDate,omitempty"`
}

type City struct {
	UID
	CountryUid      string    `json:"countryUid,omitempty" dgraph:"countryUid,omitempty"`
	StateUid        string    `json:"stateUid,omitempty" dgraph:"stateUid,omitempty"`
	CityName        string    `json:"cityName,omitempty"  dgraph:"cityName,omitempty"  validate:"required"`
	BelongsTo       UID       `json:"belongsTo,omitempty" `
	Label           string    `json:"label,omitempty"  dgraph:"label,omitempty"`
	CityUpdatedDate time.Time `json:"cityUpdatedDate,omitempty"  dgraph:"cityUpdatedDate,omitempty"`
	CityCreatedDate time.Time `json:"cityCreatedDate,omitempty"  dgraph:"cityCreatedDate,omitempty"`
}

// Response structs
type ErrorResponse struct {
	UserMessage     string `json:"userMessage"`
	InternalMessage error  `json:"internalMessage"`
}

type Response struct {
	Data    interface{}   `json:"data"`
	Message string        `json:"message"`
	Success bool          `json:"success" `
	Error   ErrorResponse `json:"error"`
}
