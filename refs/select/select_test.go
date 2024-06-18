package selecttd

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	makeDelayedServer := func(tb testing.TB,delay time.Duration) *httptest.Server{
		tb.Helper()
		currServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
		return currServer

	}
	t.Run("faster url", func(t *testing.T) {

		slowServer := makeDelayedServer(t, 20*time.Millisecond)
		fastServer := makeDelayedServer(t, 0)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got := SelectRacer(slowURL, fastURL)
		want := fastURL

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})
	
}