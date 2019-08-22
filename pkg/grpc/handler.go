package grpc

import (
	"context"
	"errors"
	"github.com/fb-investments/location/pkg/endpoint"
	"github.com/fb-investments/location/pkg/grpc/pb"
	"github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeListCountriesHandler creates the handler logic
func makeListCountriesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListCountriesEndpoint, decodeListCountriesRequest, encodeListCountriesResponse, options...)
}

// decodeListCountriesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ListCountries request.
// TODO implement the decoder
func decodeListCountriesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Decoder is not impelemented")
}

// encodeListCountriesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeListCountriesResponse(_ context.Context, r interface{}) (interface{}, error) {

	return nil, nil
}
func (g *grpcServer) ListCountries(ctx context1.Context, req *pb.ListCountriesRequest) (*pb.ListCountriesReply, error) {
	_, rep, err := g.listCountries.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListCountriesReply), nil
}

// makeCreateCountriesHandler creates the handler logic
func makeCreateCountriesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateCountriesEndpoint, decodeCreateCountriesRequest, encodeCreateCountriesResponse, options...)
}

// decodeCreateCountriesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateCountries request.
// TODO implement the decoder
func decodeCreateCountriesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Decoder is not impelemented")
}

// encodeCreateCountriesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateCountriesResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Encoder is not impelemented")
}
func (g *grpcServer) CreateCountries(ctx context1.Context, req *pb.CreateCountriesRequest) (*pb.CreateCountriesReply, error) {
	_, rep, err := g.createCountries.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateCountriesReply), nil
}

// makeListStatesHandler creates the handler logic
func makeListStatesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListStatesEndpoint, decodeListStatesRequest, encodeListStatesResponse, options...)
}

// decodeListStatesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ListStates request.
// TODO implement the decoder
func decodeListStatesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Decoder is not impelemented")
}

// encodeListStatesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeListStatesResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Encoder is not impelemented")
}
func (g *grpcServer) ListStates(ctx context1.Context, req *pb.ListStatesRequest) (*pb.ListStatesReply, error) {
	_, rep, err := g.listStates.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListStatesReply), nil
}

// makeCreateStatesHandler creates the handler logic
func makeCreateStatesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateStatesEndpoint, decodeCreateStatesRequest, encodeCreateStatesResponse, options...)
}

// decodeCreateStatesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateStates request.
// TODO implement the decoder
func decodeCreateStatesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Decoder is not impelemented")
}

// encodeCreateStatesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateStatesResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Encoder is not impelemented")
}
func (g *grpcServer) CreateStates(ctx context1.Context, req *pb.CreateStatesRequest) (*pb.CreateStatesReply, error) {
	_, rep, err := g.createStates.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateStatesReply), nil
}

// makeListCityHandler creates the handler logic
func makeListCityHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListCityEndpoint, decodeListCityRequest, encodeListCityResponse, options...)
}

// decodeListCityResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ListCity request.
// TODO implement the decoder
func decodeListCityRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Decoder is not impelemented")
}

// encodeListCityResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeListCityResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Encoder is not impelemented")
}
func (g *grpcServer) ListCity(ctx context1.Context, req *pb.ListCityRequest) (*pb.ListCityReply, error) {
	_, rep, err := g.listCity.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListCityReply), nil
}

// makeCreateCityHandler creates the handler logic
func makeCreateCityHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateCityEndpoint, decodeCreateCityRequest, encodeCreateCityResponse, options...)
}

// decodeCreateCityResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateCity request.
// TODO implement the decoder
func decodeCreateCityRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Decoder is not impelemented")
}

// encodeCreateCityResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateCityResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Location' Encoder is not impelemented")
}
func (g *grpcServer) CreateCity(ctx context1.Context, req *pb.CreateCityRequest) (*pb.CreateCityReply, error) {
	_, rep, err := g.createCity.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateCityReply), nil
}
