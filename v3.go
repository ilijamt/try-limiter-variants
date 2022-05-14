package limiter

import (
	"sync"
	"time"
)

type V3 struct {
	currentTokens, maxCapacity int64
	regenerationInterval       time.Duration
	cond                       *sync.Cond
}

func (l *V3) regenerateTokens() {
	l.cond.L.Lock()
	defer l.cond.L.Unlock()
	l.currentTokens += 1
	l.cond.Signal()
	time.AfterFunc(l.regenerationInterval, l.regenerateTokens)
}

func (l *V3) Take() bool {
	if l.TryTake() {
		return true
	}

	l.cond.L.Lock()
	l.cond.Wait()
	l.currentTokens -= 1
	l.cond.L.Unlock()
	return true
}

func (l *V3) TryTake() bool {
	l.cond.L.Lock()
	defer l.cond.L.Unlock()

	if l.currentTokens > 0 {
		l.currentTokens -= 1
		return true
	}

	return false
}

var _ Limiter = (*V3)(nil)

func NewV3(initialTokens, maxCapacity int64, regenerationInterval time.Duration) (Limiter, error) {
	if initialTokens > maxCapacity {
		initialTokens = maxCapacity
	}
	l := &V3{
		currentTokens:        initialTokens,
		maxCapacity:          maxCapacity,
		regenerationInterval: regenerationInterval,
		cond:                 sync.NewCond(&sync.Mutex{}),
	}

	time.AfterFunc(regenerationInterval, l.regenerateTokens)
	return l, nil
}
