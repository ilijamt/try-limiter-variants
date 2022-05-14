package limiter_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"limiter"
	"sync"
	"testing"
	"time"
)

func benchLimiterSyncTryTake(l limiter.Limiter, maxCalls int64) {
	var callee = func() {}
	for i := int64(0); i < maxCalls; i++ {
		l.TryTake()
		callee()
	}
}

func benchLimiterSyncTake(l limiter.Limiter, maxCalls int64) {
	var callee = func() {}
	for i := int64(0); i < maxCalls; i++ {
		l.Take()
		callee()
	}
}

func benchLimiterAsyncTake(l limiter.Limiter, maxCalls int64) {
	var callee = func() {}
	var wg = sync.WaitGroup{}
	wg.Add(int(maxCalls))
	for i := int64(0); i < maxCalls; i++ {
		go func() {
			defer wg.Done()
			l.Take()
			callee()
		}()
	}
	wg.Wait()
}

func benchLimiterAsyncTryTake(l limiter.Limiter, maxCalls int64) {
	var callee = func() {}
	var wg = sync.WaitGroup{}
	wg.Add(int(maxCalls))
	for i := int64(0); i < maxCalls; i++ {
		go func() {
			defer wg.Done()
			l.TryTake()
			callee()
		}()
	}
	wg.Wait()
}

func testLimiterSyncTryTake(t *testing.T, l limiter.Limiter, maxCalls, expectedCalls int64,
	expectedSuccess, expectedFails int64) {
	require.NotNil(t, l)
	var called uint64
	var callee = func() {
		called++
	}

	var success, fails int64

	for i := int64(0); i < maxCalls; i++ {
		if l.TryTake() {
			success += 1
		} else {
			fails += 1
		}
		callee()
	}

	assert.EqualValues(t, expectedCalls, called)
	assert.EqualValues(t, expectedSuccess, success)
	assert.EqualValues(t, expectedFails, fails)
}

func testLimiterAsyncTryTake(t *testing.T, l limiter.Limiter, maxCalls, expectedCalls, expectedSuccess, expectedFails int64) {
	require.NotNil(t, l)

	var called uint64
	var mu1 = &sync.Mutex{}
	var mu2 = &sync.Mutex{}
	var callee = func() {
		mu1.Lock()
		defer mu1.Unlock()
		called++
	}

	var success, fails int64

	var wg = &sync.WaitGroup{}
	wg.Add(int(maxCalls))

	for i := int64(0); i < maxCalls; i++ {
		go func(mu *sync.Mutex, wg *sync.WaitGroup) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			if l.TryTake() {
				success += 1
			} else {
				fails += 1
			}
			callee()
		}(mu2, wg)
	}

	wg.Wait()

	assert.EqualValues(t, expectedCalls, called)
	assert.EqualValues(t, expectedSuccess, success)
	assert.EqualValues(t, expectedFails, fails)
}

func testLimiterSyncTake(t *testing.T, l limiter.Limiter, timeStart time.Time, maxCalls, expectedCalls int64, duration time.Duration) {
	require.NotNil(t, l)
	var called uint64
	var callee = func() {
		called++
	}

	for i := int64(0); i < maxCalls; i++ {
		assert.True(t, l.Take())
		callee()
	}

	assert.EqualValues(t, expectedCalls, called)
	assert.WithinDuration(t, timeStart.Add(duration), time.Now(), 100*time.Millisecond)
}

func testLimiterAsyncTake(t *testing.T, l limiter.Limiter, timeStart time.Time, maxCalls, expectedCalls int64, duration time.Duration) {
	require.NotNil(t, l)

	var called uint64
	var mu = sync.Mutex{}
	var callee = func() {
		mu.Lock()
		defer mu.Unlock()
		called++
	}

	var wg = sync.WaitGroup{}
	wg.Add(int(maxCalls))

	for i := int64(0); i < maxCalls; i++ {
		go func() {
			defer wg.Done()
			assert.True(t, l.Take())
			callee()
		}()
	}

	wg.Wait()
	assert.EqualValues(t, expectedCalls, called)
	assert.WithinDuration(t, timeStart.Add(duration), time.Now(), 100*time.Millisecond)
}
