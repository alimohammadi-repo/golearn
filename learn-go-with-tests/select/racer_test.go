package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//func TestRacer(t *testing.T) {
//
//	slowURL := "http://www.facebook.com"
//	fastURL := "http://www.quii.dev"
//
//	want := fastURL
//	got := Racer(slowURL, fastURL)
//
//	if got != want {
//		t.Errorf("got %q, want %q", got, want)
//	}
//
//}

//func TestRacer(t *testing.T) {
//
//	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		time.Sleep(20 * time.Millisecond)
//		w.WriteHeader(http.StatusOK)
//	}))
//
//	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//	}))
//
//	slowURL := slowServer.URL
//	fastURL := fastServer.URL
//
//	want := fastURL
//	got := Racer(slowURL, fastURL)
//
//	if got != want {
//		t.Errorf("got %q, want %q", got, want)
//	}
//
//	slowServer.Close()
//	fastServer.Close()
//
//}

//func TestRacer(t *testing.T) {
//
//	slowServer := makeDelayedServer(20 * time.Millisecond)
//	fastServer := makeDelayedServer(0 * time.Millisecond)
//
//	defer slowServer.Close()
//	defer fastServer.Close()
//
//	slowURL := slowServer.URL
//	fastURL := fastServer.URL
//
//	want := fastURL
//	got := Racer(slowURL, fastURL)
//
//	if got != want {
//		t.Errorf("got %q, want %q", got, want)
//	}
//
//}

func makeDelayedServer(delay time.Duration) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}

//func TestRacer(t *testing.T) {
//	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
//		slowServer := makeDelayedServer(20 * time.Millisecond)
//		fastServer := makeDelayedServer(0 * time.Millisecond)
//
//		defer slowServer.Close()
//		defer fastServer.Close()
//
//		slowURL := slowServer.URL
//		fastURL := fastServer.URL
//
//		want := fastURL
//		got, _ := Racer(slowURL, fastURL)
//
//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//	})
//
//	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
//		serverA := makeDelayedServer(11 * time.Second)
//		serverB := makeDelayedServer(12 * time.Second)
//
//		defer serverA.Close()
//		defer serverB.Close()
//
//		_, err := Racer(serverA.URL, serverB.URL)
//
//		if err == nil {
//			t.Error("expected an error but didn't get one")
//		}
//	})
//
//}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Second)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, time.Second*20)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}
