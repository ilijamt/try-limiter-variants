package limiter

import (
	"time"
)

type V2 struct {
	regenerationInterval time.Duration
	tokens               chan interface{}
}

func (l *V2) regenerateTokens() {
	l.tokens <- nil
	time.AfterFunc(l.regenerationInterval, l.regenerateTokens)
}

func (l *V2) Take() bool {
	<-l.tokens
	return true
}

func (l *V2) TryTake() bool {
	select {
	case <-l.tokens:
	default:
		return false
	}

	return true
}

var _ Limiter = (*V2)(nil)

func NewV2(initialTokens, maxCapacity int64, regenerationInterval time.Duration) (Limiter, error) {
	l := &V2{
		tokens:               make(chan interface{}, maxCapacity),
		regenerationInterval: regenerationInterval,
	}
	if initialTokens > maxCapacity {
		initialTokens = maxCapacity
	}
	for i := int64(0); i < initialTokens; i++ {
		l.tokens <- nil
	}
	time.AfterFunc(regenerationInterval, l.regenerateTokens)
	return l, nil
}
