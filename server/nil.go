package main

import (
	"github.com/sirupsen/logrus"
)

type nilProcessor struct{}

func (p *nilProcessor) Process(input *ProcessInput) {
	logrus.Printf("Processing input: %#v", input)
}
