package service

import (
	"github.com/jwt-example/internal/cache"
	"github.com/jwt-example/internal/config"
)

// Service holds all the different services that are used when handling request.
type Service struct {
	Cache      *cache.Cache
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (*Service, error) {
	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		Cache:      c,
	}, nil
}