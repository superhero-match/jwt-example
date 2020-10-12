package cache

import (
	"github.com/jwt-example/internal/cache/model"
	"strconv"
	"time"
)

func(c *Cache) CreateAuth(userID uint64, td model.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)

	now := time.Now()

	errAccess := c.Redis.Set(td.AccessUuid, strconv.Itoa(int(userID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := c.Redis.Set(td.RefreshUuid, strconv.Itoa(int(userID)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}
