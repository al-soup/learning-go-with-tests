package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of server, returning the URL of the faster one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// Prefixing a function with `defer` will call that function at the end of the containing function.
		// This improves readability and we don't have to remember doing certain things at the end of the function.
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		fmt.Printf("url %s", fastURL)

		want := fastURL
		got, err := ConfigurableRacer(slowURL, fastURL, 30*time.Millisecond)

		if err != nil {
			// Different opinions: don't use Fatalf for failing tests. Only for errors in setup or to terminate the testing process
			// Other opinion: from this point on it does no longer make sense to test
			// t.Fatalf("did not expect an error but got one %v", err)
			t.Errorf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't return within 10s", func(t *testing.T) {

		server := makeDelayedServer(20 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
