package dgraph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fb-investments/location/pkg/io"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"time"
)

var (
	ErrRepository = errors.New("unable to handle request")
)

// Repository describes the resource methods needed for this service.
type Repository interface {
	ListCountry(ctx context.Context) (c []io.Country, err error)
	CreateCountry(ctx context.Context, data io.Country) (err error)

	ListState(ctx context.Context, countryUid string) (c []io.State, err error)
	CreateState(ctx context.Context, data io.State) (err error)

	ListCities(ctx context.Context, stateUid string) (c []io.City, err error)
	CreateCity(ctx context.Context, data io.City) (err error)

	IsLocationExist(ctx context.Context, uid, locationType string, ) (isExist bool, err error)
}

type repository struct {
	db     *dgo.Dgraph
	logger log.Logger
}

// New returns a concrete repository backed by dgraph
func New(addr *string, logger log.Logger) (Repository) {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		logger.Log("method", "main dgraph", "err", err)
	}
	dc := api.NewDgraphClient(conn)
	dbConn := dgo.NewDgraphClient(dc)

	// return  repository
	return &repository{
		db:     dbConn,
		logger: log.With(logger, "Repository", "dgraph"),
	}
}

// ChangeOrderStatus changes the order status
func (repo *repository) IsLocationExist(ctx context.Context, uid, locationType string, ) (isExist bool, err error) {

	method := "IsLocationExist"

	q := fmt.Sprintf(`
	{
		Location(func: eq(label,%q)) @filter (uid(%q)){
    		expand(_all_)
  		}  
	}
`, locationType, uid)
	//fmt.Println("-----------", q)
	rs, err := repo.db.NewTxn().Query(ctx, q)
	if err != nil {
		msg := "invalid id"
		repo.logger.Log("method", method, "Error", msg)
		//isExist = true
		return
	}
	type Root struct {
		Location []io.City `json:"Location"`
	}

	var r Root
	err = json.Unmarshal(rs.Json, &r)
	if err != nil {
		repo.logger.Log("method", method, "Error", err.Error())
		isExist = false
		return
	}
	//fmt.Printf("------------%+v\n", q)
	repo.logger.Log("method", method, "Location Data Len", len(r.Location))
	if len(r.Location) > 0 {
		msg := "data found"
		repo.logger.Log("method", method, "Error", msg)
		isExist = true
		//err = fmt.Errorf(msg)
		return
	}
	return
}

// ChangeOrderStatus changes the order status
func (repo *repository) ListCountry(ctx context.Context) (c []io.Country, err error) {
	method := "list"
	q := `
    	query Countries($id: string){
			Countries(func: eq(label,$id)) {
				    uid
    				countryName
			}
    	}
	`
	//resp, err := repo.db.NewReadOnlyTxn().Query(ctx, q)
	variables := map[string]string{"$id": io.Lable_Country}
	resp, err := repo.db.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		repo.logger.Log("method", method, "Type", "Query", "Error", err, "label", io.Lable_Country, "query", q)
		return
	}

	type Root struct {
		Countries []io.Country `json:"Countries"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		repo.logger.Log("method", method, "unmarshalling", "Find", "Error", err)
		return
	}

	c = r.Countries
	return
}

func (repo *repository) CreateCountry(ctx context.Context, data io.Country) (err error) {

	data.Label = io.Lable_Country
	data.CountryCreatedDate = time.Now()
	pb, err := json.Marshal(data)
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb
	_, err = repo.db.NewTxn().Mutate(ctx, mu)
	if err != nil {
		repo.logger.Log("method", "Create", "Error", err)
		return
	}
	return
}
func (repo *repository) ListState(ctx context.Context, countryUid string) (c []io.State, err error) {
	method := "list"
	q := `
    	query States($id: string,$label: string){
			States(func: eq(countryUid,$id))  @filter(eq(label,$label)){
				    uid
                    stateName
			}
    	}
	`
	//resp, err := repo.db.NewReadOnlyTxn().Query(ctx, q)
	variables := map[string]string{"$id": countryUid, "$label": io.Lable_State}
	resp, err := repo.db.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		repo.logger.Log("method", method, "Type", "Query", "Error", err, "Country UID", countryUid, "query", q)
		return
	}

	type Root struct {
		States []io.State `json:"States"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		repo.logger.Log("method", method, "Unmarshalling", "Find", "Error", err)
		return
	}

	c = r.States
	return
}

func (repo *repository) CreateState(ctx context.Context, state io.State) (err error) {

	stateData := io.CountryStates{}

	state.Label = io.Lable_State
	stateData.Uid = state.CountryUid
	stateData.HasState = make([]io.State, 1)
	stateData.HasState[0] = state
	pb, err := json.Marshal(stateData)
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb
	_, err = repo.db.NewTxn().Mutate(ctx, mu)
	if err != nil {
		repo.logger.Log("method", "CreateState", "Error", err)
		return
	}
	return nil
}

func (repo *repository) ListCities(ctx context.Context, stateUid string) (c []io.City, err error) {
	method := "list"
	q := `
    	query Cities($id: string,$label: string){
			Cities(func: eq(stateUid,$id)) @filter(eq(label,$label)){
				    uid
    				cityName
			}
    	}
	`
	//resp, err := repo.db.NewReadOnlyTxn().Query(ctx, q)
	variables := map[string]string{"$id": stateUid, "$label": io.Lable_City}
	resp, err := repo.db.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		repo.logger.Log("method", method, "Type", "Query", "Error", err, "State UID", stateUid, "query", q)
		return
	}

	type Root struct {
		Cities []io.City `json:"Cities"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		repo.logger.Log("method", method, "unmarshalling", "Find", "Error", err)
		return
	}

	c = r.Cities
	return
}

func (repo *repository) CreateCity(ctx context.Context, city io.City) (err error) {

	cityData := io.StateCities{}
	city.BelongsTo.Uid = city.CountryUid
	city.Label = io.Lable_City
	cityData.Uid = city.StateUid
	cityData.HasCities = make([]io.City, 1)
	cityData.HasCities[0] = city
	pb, err := json.Marshal(cityData)
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb
	_, err = repo.db.NewTxn().Mutate(ctx, mu)
	if err != nil {
		repo.logger.Log("method", "CreateCity", "Error", err)
		return
	}
	return nil
}

// Close implements DB.Close
func (repo *repository) Close() error {
	return repo.Close()
}
