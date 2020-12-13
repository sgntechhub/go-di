package rest

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"go_id/services/usecases"
)

type Endpoint struct {
	Ping                 endpoint.Endpoint
	GetInfoHome          endpoint.Endpoint
	GetCateRecommends    endpoint.Endpoint
	GetCateRecommendsWeb endpoint.Endpoint
	PostRecommendHistory endpoint.Endpoint
}

func NewEndpoint(svc usecases.Service) Endpoint {
	ping := makePingEndpoint(svc)
	return Endpoint{
		Ping: ping,
	}
}

func makePingEndpoint(svc usecases.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Ping(ctx), nil
	}
}
