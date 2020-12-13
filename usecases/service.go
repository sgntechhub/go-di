package usecases

import (
	"context"
)

type Service interface {
	Ping(ctx context.Context) string
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Ping(ctx context.Context) string {
	return "pong"
}
