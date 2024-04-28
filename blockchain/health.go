package blockchain

import (
	"time"
)

type ErrorTolerantReportConfig struct {
	Threshold time.Duration `default:"1m"` // report if error persists for a long time
	Reminder  time.Duration `default:"5m"` // remind if error persists for a long time
}

// Health represents an error tolerant health status, which allow errors in short time and periodically
// remind unhealthy if not recovered in a long time.
type Health struct {
	failedAt time.Time
	counter  int
}

func (h *Health) HasError() bool {
	return !h.failedAt.IsZero()
}

// OnSuccess erases error status and return recover information.
//
// `recovered` indicates if recovered from unhealthy status.
//
// `elapsed` is the duration since failed time.
func (h *Health) OnSuccess(config *ErrorTolerantReportConfig) (recovered bool, elapsed time.Duration) {
	return h.onSuccess(config, time.Now())
}

func (h *Health) onSuccess(config *ErrorTolerantReportConfig, now time.Time) (recovered bool, elapsed time.Duration) {
	// last time was healthy
	if h.failedAt.IsZero() {
		return
	}

	// report health now after a long time
	if elapsed = now.Sub(h.failedAt); elapsed > config.Threshold {
		recovered = true
	}

	// reset
	h.failedAt = time.Time{}
	h.counter = 0

	return
}

// OnSuccess marks error status and return unhealthy information.
//
// `unhealthy` indicates continous failures in a long time.
//
// `unrecovered` indicates continous failures and not recovered in a long time.
//
// `elapsed` is the duration since failed time.
func (h *Health) OnFailure(config *ErrorTolerantReportConfig) (unhealthy bool, unrecovered bool, elapsed time.Duration) {
	return h.onFailure(config, time.Now())
}

func (h *Health) onFailure(config *ErrorTolerantReportConfig, now time.Time) (unhealthy bool, unrecovered bool, elapsed time.Duration) {
	// mark unhealthy
	if h.failedAt.IsZero() {
		h.failedAt = now
		return
	}

	// error tolerant
	if elapsed = now.Sub(h.failedAt); elapsed < config.Threshold {
		return
	}

	// become unhealthy
	if h.counter == 0 {
		unhealthy = true
		h.counter++
		return
	}

	// error tolerant
	if remind := config.Threshold + time.Duration(h.counter)*config.Reminder; elapsed < remind {
		return
	}

	// remind unhealthy
	unrecovered = true
	h.counter++

	return
}
