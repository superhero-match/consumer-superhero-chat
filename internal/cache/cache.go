/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package cache

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/superhero-match/consumer-superhero-chat/internal/config"
)

// Cache is the Redis client.
type Cache struct {
	Redis             *redis.Client
	MessagesKeyFormat string
	TokenKeyFormat    string
}

// NewCache creates a client connection to Redis.
func NewCache(cfg *config.Config) (cache *Cache, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s%s", cfg.Cache.Address, cfg.Cache.Port),
		Password:     cfg.Cache.Password,
		DB:           cfg.Cache.DB,
		PoolSize:     cfg.Cache.PoolSize,
		MinIdleConns: cfg.Cache.MinimumIdleConnections,
		MaxRetries:   cfg.Cache.MaximumRetries,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Cache{
		Redis:             client,
		MessagesKeyFormat: cfg.Cache.MessagesKeyFormat,
		TokenKeyFormat:    cfg.Cache.TokenKeyFormat,
	}, nil
}
