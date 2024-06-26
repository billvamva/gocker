package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	assertCounter := func(t testing.TB, got *Counter, want int) {
		t.Helper()
		if got.Value() != want {
			t.Errorf("got %d, want %d", got.Value(), want)
		}
	}
	t.Run("increment counter 3 times", func(t *testing.T) {
		counter := Counter{}
		for i:=0;i<3;i++ {
			counter.Inc()
		}
		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})	
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
	
		var wg sync.WaitGroup
		wg.Add(wantedCount)
	
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
	
		assertCounter(t, counter, wantedCount)
	})
}