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
package reader

import (
	"go.uber.org/zap"

	"github.com/superhero-match/consumer-superhero-chat/internal/cache"
	"github.com/superhero-match/consumer-superhero-chat/internal/config"
	"github.com/superhero-match/consumer-superhero-chat/internal/consumer"
	"github.com/superhero-match/consumer-superhero-chat/internal/db"
	"github.com/superhero-match/consumer-superhero-chat/internal/firebase"
)

const timeFormat = "2006-01-02T15:04:05"

// Reader holds all the data relevant.
type Reader struct {
	DB         *db.DB
	Consumer   *consumer.Consumer
	Cache      *cache.Cache
	Firebase   *firebase.Firebase
	Logger     *zap.Logger
	TimeFormat string
}

// NewReader configures Reader.
func NewReader(cfg *config.Config) (r *Reader, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	ch, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	cs, err := consumer.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	f := firebase.NewFirebase(cfg)

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Reader{
		DB:         dbs,
		Consumer:   cs,
		Cache:      ch,
		Firebase:   f,
		Logger:     logger,
		TimeFormat: timeFormat,
	}, nil
}
