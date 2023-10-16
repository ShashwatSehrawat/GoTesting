package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare the server based on response time", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(13 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Error("did not expect an error but got one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("return an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(30 * time.Millisecond)
		fastServer := makeDelayedServer(22 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		_, err := ConfigurableRacer(slowURL, fastURL, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
