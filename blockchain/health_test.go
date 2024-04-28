package blockchain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTestErrorTolerantReportConfig() ErrorTolerantReportConfig {
	return ErrorTolerantReportConfig{
		Threshold: time.Minute,
		Reminder:  5 * time.Minute,
	}
}

func TestHealthContinousSuccess(t *testing.T) {
	config := newTestErrorTolerantReportConfig()
	var health Health

	recovered, elapsed := health.onSuccess(&config, time.Now().Add(config.Threshold+1))
	assert.False(t, recovered)
	assert.Equal(t, time.Duration(0), elapsed)
}

func TestHealthFailedShortTime(t *testing.T) {
	config := newTestErrorTolerantReportConfig()
	var health Health

	now := time.Now()

	// first failure
	unhealthy, unrecovered, elapsed := health.onFailure(&config, now)
	assert.False(t, unhealthy)
	assert.False(t, unrecovered)
	assert.Equal(t, time.Duration(0), elapsed)

	// continous failure in short time
	unhealthy, unrecovered, elapsed = health.onFailure(&config, now.Add(config.Threshold-2))
	assert.False(t, unhealthy)
	assert.False(t, unrecovered)
	assert.Equal(t, config.Threshold-2, elapsed)

	// recovered
	recovered, elapsed := health.onSuccess(&config, now.Add(config.Threshold-1))
	assert.False(t, recovered)
	assert.Equal(t, config.Threshold-1, elapsed)
}

func TestHealthThreshold(t *testing.T) {
	config := newTestErrorTolerantReportConfig()
	var health Health

	now := time.Now()

	// first failure
	health.onFailure(&config, now)

	// continous failure in short time
	health.onFailure(&config, now.Add(config.Threshold-1))

	// continous failure in long time
	unhealthy, unrecovered, elapsed := health.onFailure(&config, now.Add(config.Threshold+1))
	assert.True(t, unhealthy)
	assert.False(t, unrecovered)
	assert.Equal(t, config.Threshold+1, elapsed)

	// recovered
	recovered, elapsed := health.onSuccess(&config, now.Add(config.Threshold+2))
	assert.True(t, recovered)
	assert.Equal(t, config.Threshold+2, elapsed)
}

func TestHealthRemind(t *testing.T) {
	config := newTestErrorTolerantReportConfig()
	var health Health

	now := time.Now()

	// first failure
	health.onFailure(&config, now)

	// continous failure in short time
	health.onFailure(&config, now.Add(config.Threshold-1))

	// continous failure in long time
	health.onFailure(&config, now.Add(config.Threshold+1))

	// continous failure in long time, but not reached remind time
	unhealthy, unrecovered, elapsed := health.onFailure(&config, now.Add(config.Threshold+2))
	assert.False(t, unhealthy)
	assert.False(t, unrecovered)
	assert.Equal(t, config.Threshold+2, elapsed)

	// continous failure and reached remind time
	unhealthy, unrecovered, elapsed = health.onFailure(&config, now.Add(config.Threshold+2+config.Reminder))
	assert.False(t, unhealthy)
	assert.True(t, unrecovered)
	assert.Equal(t, config.Threshold+2+config.Reminder, elapsed)

	// recovered
	recovered, elapsed := health.onSuccess(&config, now.Add(config.Threshold+3+config.Reminder))
	assert.True(t, recovered)
	assert.Equal(t, config.Threshold+3+config.Reminder, elapsed)
}
