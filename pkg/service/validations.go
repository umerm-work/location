package service

import (
	"github.com/fb-investments/location/pkg/io"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate = validator.New()
)

func (b *basicLocationService) ValidateCreateCountry(u io.Country) (err error) {
	err = validate.Struct(u)
	return
}

func (b *basicLocationService) ValidateCreateState(u io.State) (err error) {
	err = validate.Struct(u)
	return
}

func (b *basicLocationService) ValidateCreateCity(u io.City) (err error) {
	err = validate.Struct(u)
	return
}
