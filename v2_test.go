package limiter_test

import (
	"github.com/stretchr/testify/require"
	"limiter"
	"testing"
	"time"
)

func TestV2(t *testing.T) {
	t.Parallel()
	t.Run("15 calls in 5 seconds, serially, all succeed", func(t *testing.T) {
		t.Parallel()
		timeStart := time.Now()
		l, err := limiter.NewV2(15, 10, time.Second)
		require.NoError(t, err)
		testLimiterSyncTake(t, l, timeStart, 15, 15, 5*time.Second)
	})
	t.Run("15 calls in 5 seconds, async, all succeed", func(t *testing.T) {
		t.Parallel()
		timeStart := time.Now()
		l, err := limiter.NewV2(10, 100, time.Second)
		require.NoError(t, err)
		testLimiterAsyncTake(t, l, timeStart, 15, 15, 5*time.Second)
	})
	t.Run("15 calls in 5 seconds, serially, take, 10 succeed, 5 fail", func(t *testing.T) {
		t.Parallel()
		l, err := limiter.NewV2(10, 100, time.Second)
		require.NoError(t, err)
		testLimiterSyncTryTake(t, l, 15, 15, 10, 5)
	})

	t.Run("15 calls in 5 seconds, async, take, 10 succeed, 5 fail", func(t *testing.T) {
		t.Parallel()
		l, err := limiter.NewV2(10, 100, time.Second)
		require.NoError(t, err)
		testLimiterAsyncTryTake(t, l, 15, 15, 10, 5)
	})

}

func BenchmarkV2_Sync_Take(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l, _ := limiter.NewV2(5, 10, time.Millisecond)
		benchLimiterSyncTake(l, 5)
	}
}

func BenchmarkV2_Async_Take(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l, _ := limiter.NewV2(5, 10, time.Millisecond)
		benchLimiterAsyncTake(l, 5)
	}
}

func BenchmarkV2_Sync_TryTake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l, _ := limiter.NewV2(5, 10, time.Second)
		benchLimiterSyncTryTake(l, 5)
	}
}

func BenchmarkV2_Async_TryTake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l, _ := limiter.NewV2(5, 10, time.Second)
		benchLimiterAsyncTryTake(l, 5)
	}
}
