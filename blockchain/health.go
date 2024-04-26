package blockchain

import (
	"time"
)

type ErrorTolerantReportConfig struct {
	Threshold time.Duration `default:"1m"`
	Reminder  time.Duration `default:"5m"`
}

type Health struct {
	unhealthyAt time.Time
	counter     int
}

func (h *Health) OnSuccess(config *ErrorTolerantReportConfig) (recovered bool, elapsed time.Duration) {
	// last time was healthy
	if h.unhealthyAt.IsZero() {
		return
	}

	// report health now after a long time
	if elapsed = time.Since(h.unhealthyAt); elapsed > config.Threshold {
		recovered = true
	}

	// reset
	h.unhealthyAt = time.Time{}
	h.counter = 0

	return
}

func (h *Health) OnFailure(config *ErrorTolerantReportConfig) (unhealthy bool, unrecovered bool, elapsed time.Duration) {
	// mark unhealthy
	if h.unhealthyAt.IsZero() {
		h.unhealthyAt = time.Now()
		return
	}

	// error tolerant
	elapsed = time.Since(h.unhealthyAt)
	if elapsed < config.Threshold {
		return
	}

	// become unhealthy
	if h.counter == 0 {
		unhealthy = true
		h.counter++
		return
	}

	// error tolerant
	remind := config.Threshold + time.Duration(h.counter)*config.Reminder
	if elapsed < remind {
		return
	}

	// remind unhealthy
	unrecovered = true
	h.counter++

	return
}
