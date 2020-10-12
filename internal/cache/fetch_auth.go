package cache

import (
	"github.com/jwt-example/internal/cache/model"
	"strconv"
)

func (c *Cache) FetchAuth(authD *model.AccessDetails) (uint64, error) {
	userid, err := c.Redis.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)

	return userID, nil
}

