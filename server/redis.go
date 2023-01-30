package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func NewRedisProcessor(address string, password string, db int) BackendProcessor {
	return &redisProcessor{
		client: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		}),
	}
}

type redisProcessor struct {
	client *redis.Client
}

func (p *redisProcessor) Process(input *ProcessInput) {
	var id = input.ID.String()

	existQuery := p.client.Exists(context.Background(), input.ID.String())
	if existQuery.Err() != nil {
		logrus.Errorf("Failed to query for existing item: %v", existQuery.Err())
	}

	if count, _ := existQuery.Uint64(); count > 0 {
		logrus.Infof("Existing item found, marking this one as the ending")
		id = fmt.Sprintf("%s_end", id)
	}

	logrus.Infof("Created %s with timestamp %v", id, input.Timestamp)
	if err := p.client.Set(context.Background(), id, input.Timestamp, 0).Err(); err != nil {
		logrus.Errorf("Error while processing input: %v", err)
	}
}
