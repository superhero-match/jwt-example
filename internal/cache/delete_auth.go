package cache

func (c *Cache) DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := c.Redis.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}
