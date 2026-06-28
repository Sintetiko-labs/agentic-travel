package ratelimit

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Pacer enforces a minimum delay between HTTP requests.
type Pacer struct {
	Delay     time.Duration
	mu        sync.Mutex
	lastReqAt time.Time
}

// NewPacerFromEnv reads {prefix}_REQUEST_DELAY (e.g. "2s", "500ms", or plain seconds).
func NewPacerFromEnv(prefix string) *Pacer {
	v := strings.TrimSpace(os.Getenv(prefix + "_REQUEST_DELAY"))
	if v == "" {
		return &Pacer{}
	}
	if d, err := time.ParseDuration(v); err == nil {
		return &Pacer{Delay: d}
	}
	if secs, err := strconv.Atoi(v); err == nil && secs >= 0 {
		return &Pacer{Delay: time.Duration(secs) * time.Second}
	}
	return &Pacer{}
}

// Wait blocks until the next request slot is available.
func (p *Pacer) Wait() {
	if p == nil || p.Delay <= 0 {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.lastReqAt.IsZero() {
		if wait := p.Delay - time.Since(p.lastReqAt); wait > 0 {
			time.Sleep(wait)
		}
	}
	p.lastReqAt = time.Now()
}

// FixedThrottle enforces a fixed minimum interval (e.g. 200ms between calls).
func FixedThrottle(p *Pacer, interval time.Duration) {
	if interval <= 0 {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if d := interval - time.Since(p.lastReqAt); d > 0 {
		time.Sleep(d)
	}
	p.lastReqAt = time.Now()
}
