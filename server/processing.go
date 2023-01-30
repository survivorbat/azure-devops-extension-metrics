package main

import (
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"time"
)

type ProcessInput struct {
	ID        uuid.UUID
	Timestamp time.Time
}

func getProcessor(id string) BackendProcessor {
	switch id {
	case "redis":
		log.Printf("Using Redis backend %s", os.Getenv("REDIS_HOST"))
		db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			log.Fatalf("Error while parsing REDIS_DB: %v", err)
		}
		return NewRedisProcessor(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), db)
	default:
		return &nilProcessor{}
	}
}

type BackendProcessor interface {
	Process(input *ProcessInput)
}
