package main

import (
	"github.com/google/uuid"
	"log"
	"time"
)

type ProcessInput struct {
	ID        uuid.UUID
	Timestamp time.Time
}

func getProcessor(id string) IBackendProcessor {
	return &nilProcessor{}
}

type IBackendProcessor interface {
	Process(input *ProcessInput)
}

type nilProcessor struct{}

func (p *nilProcessor) Process(input *ProcessInput) {
	log.Printf("Processing input: %#v\n", input)
}
