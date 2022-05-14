package limiter

import (
	"sync"
	"time"
)

type V1 struct {
	currentTokens, maxCapacity int64
	regenerationInterval       time.Duration

	mu sync.Mutex
}

func (l *V1) regenerateTokens() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.currentTokens += 1
	time.AfterFunc(l.regenerationInterval, l.regenerateTokens)
}

func (l *V1) Take() bool {
	if l.TryTake() {
		return true
	}

	t := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-t.C:
			l.mu.Lock()
			if l.currentTokens > 0 {
				l.currentTokens -= 1
				l.mu.Unlock()
				return true
			}
			l.mu.Unlock()
		}
	}
}

func (l *V1) TryTake() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.currentTokens > 0 {
		l.currentTokens -= 1
		return true
	}

	return false
}

var _ Limiter = (*V1)(nil)

func NewV1(initialTokens, maxCapacity int64, regenerationInterval time.Duration) (Limiter, error) {
	if initialTokens > maxCapacity {
		initialTokens = maxCapacity
	}
	l := &V1{
		currentTokens:        initialTokens,
		maxCapacity:          maxCapacity,
		regenerationInterval: regenerationInterval,
	}
	time.AfterFunc(regenerationInterval, l.regenerateTokens)
	return l, nil
}
