package engine

import (
	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

//TestNewPooledConfigOk
func TestNewPooledConfigDefault(t *testing.T) {
	pooledConfig := NewPooledConfig()

	// assert Success
	assert.Equal(t, config.RUNNER_WORKERS_DEFAULT, pooledConfig.NumWorkers)
	assert.Equal(t, config.RUNNER_QUEUE_SIZE_DEFAULT, pooledConfig.WorkQueueSize)
}

//TestNewPooledConfigOk
func TestNewPooledConfigOverride(t *testing.T) {
	previousWorkers := os.Getenv(config.RUNNER_WORKERS_KEY)
	defer os.Setenv(config.RUNNER_WORKERS_KEY, previousWorkers)
	previousQueue := os.Getenv(config.RUNNER_QUEUE_SIZE_KEY)
	defer os.Setenv(config.RUNNER_QUEUE_SIZE_KEY, previousQueue)

	newWorkersValue := 6
	newQueueValue := 60

	// Change values
	os.Setenv(config.RUNNER_WORKERS_KEY, strconv.Itoa(newWorkersValue))
	os.Setenv(config.RUNNER_QUEUE_SIZE_KEY, strconv.Itoa(newQueueValue))

	pooledConfig := NewPooledConfig()

	// assert Success
	assert.Equal(t, newWorkersValue, pooledConfig.NumWorkers)
	assert.Equal(t, newQueueValue, pooledConfig.WorkQueueSize)
}
