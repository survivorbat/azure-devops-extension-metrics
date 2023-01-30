package main

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNilProcessor_Process_DoesNothing(t *testing.T) {
	t.Parallel()
	// Arrange
	process := &nilProcessor{}
	input := &ProcessInput{
		ID:        uuid.MustParse("825bce68-9ad0-41f1-83cf-ac4e7f3f85db"),
		Timestamp: time.Time{},
	}

	// Act
	result := func() { process.Process(input) }

	// Assert
	assert.NotPanics(t, result)
}
