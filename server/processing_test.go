package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProcessor_ReturnsNilAsDefault(t *testing.T) {
	t.Parallel()
	// Act
	result := getProcessor("any")

	// Assert
	assert.IsType(t, &nilProcessor{}, result)
}

func TestGetProcessor_ReturnsRedisOnRedis(t *testing.T) {
	// Arrange
	t.Setenv("REDIS_DB", "0")

	// Act
	result := getProcessor("redis")

	// Assert
	assert.IsType(t, &redisProcessor{}, result)
}
