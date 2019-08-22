package service

import (
	"context"

	"github.com/fb-investments/location/pkg/db/dgraph"
	"github.com/fb-investments/location/pkg/io"
	"github.com/go-kit/kit/log"
)

// LocationService describes the service.
type LocationService interface {
	// Add your methods here
	ListCountries(ctx context.Context) (res io.Response)
	CreateCountries(ctx context.Context, c io.Country) (res io.Response)
	ListStates(ctx context.Context, countryUid io.UID) (res io.Response)
	CreateStates(ctx context.Context, c io.State) (res io.Response)
	ListCity(ctx context.Context, stateUid io.UID) (res io.Response)
	CreateCity(ctx context.Context, c io.City) (res io.Response)
}

type basicLocationService struct {
	DbRepo dgraph.Repository
	logger log.Logger
}

// NewBasicLocationService returns a naive, stateless implementation of LocationService.
func NewBasicLocationService(db dgraph.Repository, lo log.Logger) LocationService {
	return &basicLocationService{
		DbRepo: db,
		logger: lo,
	}
}

func (b *basicLocationService) ListCountries(ctx context.Context) (res io.Response) {
	method := "ListCountries"
	var countryList []io.Country
	if countryList, res.Error.InternalMessage = b.DbRepo.ListCountry(ctx); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error", res.Error.InternalMessage.Error())
		res = io.FailureMessage(res.Error, "failed to list the country")
		return
	}
	data := make(map[string][]io.Country)
	data["countries"] = countryList
	res = io.SuccessMessage(data, "success")
	return
}
func (b *basicLocationService) CreateCountries(ctx context.Context, c io.Country) (res io.Response) {

	method := "CreateCountries"
	if res.Error.InternalMessage = b.ValidateCreateCountry(c); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error type", "validation", "error", res.Error.InternalMessage.Error())
		res.Error.UserMessage = res.Error.InternalMessage.Error()
		res = io.FailureMessage(res.Error, "failed to validate the country")
		return
	}

	if res.Error.InternalMessage = b.DbRepo.CreateCountry(ctx, c); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error", res.Error.InternalMessage.Error())
		res.Error.UserMessage = res.Error.InternalMessage.Error()
		res = io.FailureMessage(res.Error, "failed to create the country")
		return
	}
	res = io.SuccessMessage(nil, "country have been successfully created")
	return
}

func (b *basicLocationService) ListStates(ctx context.Context, countryUid io.UID) (res io.Response) {
	method := "ListStates"
	var stateList []io.State
	if stateList, res.Error.InternalMessage = b.DbRepo.ListState(ctx, countryUid.Uid); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error", res.Error.InternalMessage)
		res = io.FailureMessage(res.Error, "failed to list the state")
		return
	}
	data := make(map[string][]io.State)
	data["states"] = stateList
	res = io.SuccessMessage(data, "success")
	return
}
func (b *basicLocationService) CreateStates(ctx context.Context, c io.State) (res io.Response) {
	method := "CreateStates"
	if res.Error.InternalMessage = b.ValidateCreateState(c); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error type", "validation", "error", res.Error.InternalMessage.Error())
		res = io.FailureMessage(res.Error, "failed to validate the state")
		return
	}
	var isExist bool
	if isExist, res.Error.InternalMessage = b.DbRepo.IsLocationExist(ctx, c.CountryUid, io.Lable_Country); !isExist || res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error type", "db", "error", res.Error.InternalMessage)
		res = io.FailureMessage(res.Error, "country doesn't exist")
		return
	}
	if res.Error.InternalMessage = b.DbRepo.CreateState(ctx, c); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error", res.Error.InternalMessage)
		res = io.FailureMessage(res.Error, "failed to create the state")
		return
	}
	res = io.SuccessMessage(nil, "state have been successfully created")
	return
}
func (b *basicLocationService) ListCity(ctx context.Context, stateUid io.UID) (res io.Response) {

	method := "ListCity"
	var cityList []io.City
	if cityList, res.Error.InternalMessage = b.DbRepo.ListCities(ctx, stateUid.Uid); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error", res.Error.InternalMessage.Error())
		res = io.FailureMessage(res.Error, "failed to list the state")
		return
	}
	data := make(map[string][]io.City)
	data["cities"] = cityList
	res = io.SuccessMessage(data, "success")
	return
}
func (b *basicLocationService) CreateCity(ctx context.Context, c io.City) (res io.Response) {

	method := "CreateCity"
	if res.Error.InternalMessage = b.ValidateCreateCity(c); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error type", "validation", "error", res.Error.InternalMessage.Error())
		res = io.FailureMessage(res.Error, "failed to validate the city")
		return
	}
	var isExist bool
	if isExist, res.Error.InternalMessage = b.DbRepo.IsLocationExist(ctx, c.StateUid, io.Lable_State); !isExist || res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error type", "db", "error", res.Error.InternalMessage.Error())
		res = io.FailureMessage(res.Error, "state doesn't exist")
		return
	}
	if res.Error.InternalMessage = b.DbRepo.CreateCity(ctx, c); res.Error.InternalMessage != nil {
		b.logger.Log("method", method, "error", res.Error.InternalMessage.Error())
		res = io.FailureMessage(res.Error, "failed to create the city")
		return
	}
	res = io.SuccessMessage(nil, "city have been successfully created")
	return
}

// New returns a LocationService with all of the expected middleware wired in.
func New(middleware []Middleware, db dgraph.Repository, lo log.Logger) LocationService {
	var svc LocationService = NewBasicLocationService(db, lo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
