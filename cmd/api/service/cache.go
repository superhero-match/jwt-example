package service

import (
	"github.com/jwt-example/internal/cache/model"
)

func (srv *Service) CreateAuth(userID uint64, td model.TokenDetails) error {
	return  srv.Cache.CreateAuth(userID, td)
}

func (srv *Service) FetchAuth(authD *model.AccessDetails) (uint64, error) {
	return srv.Cache.FetchAuth(authD)
}

func (srv *Service) DeleteAuth(givenUuid string) (int64, error) {
	return srv.Cache.DeleteAuth(givenUuid)
}
