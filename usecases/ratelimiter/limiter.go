package ratelimiter

import "sync"

// Limiter is a thread-safe counter that can be used for rate-limiting requests
type Limiter struct {
	lock    sync.Mutex
	max     int
	current int
}

// New creates a [Limiter] with the specified maximum concurrent requests
func New(maxRequests int) *Limiter {
	return &Limiter{
		max: maxRequests,
	}
}

// If there is still room, TryInc, increases the counter and returns true. If
// there are too many concurrent requests it does not increase the counter and
// returns false
func (l *Limiter) TryInc() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.current < l.max {
		l.current++
		return true

	}

	return false
}

func (l *Limiter) Dec() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.current--
}
