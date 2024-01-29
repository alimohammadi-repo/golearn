package context

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

//func Server(store Store) http.HandlerFunc {
//
//	return func(writer http.ResponseWriter, request *http.Request) {
//		store.Cancel()
//		fmt.Fprint(writer, store.Fetch())
//	}
//
//}

//func Server(store Store) http.HandlerFunc {
//
//	return func(writer http.ResponseWriter, request *http.Request) {
//
//		ctx := request.Context()
//		data := make(chan string, 1)
//		go func() {
//			data <- store.Fetch()
//		}()
//
//		select {
//		case d := <-data:
//			fmt.Fprintln(writer, d)
//		case <-ctx.Done():
//			store.Cancel()
//		}
//
//	}
//
//}

//type Store interface {
//	Fetch() string
//	Cancel()
//}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())

		if err != nil {
			return
		}

		fmt.Fprint(writer, data)
	}
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {

			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil

	}

}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

//type StubStore struct {
//	response string
//}
//
//func (s *StubStore) Fetch() string {
//	return s.response
//}

//type SpyStore struct {
//	response  string
//	cancelled bool
//}

//type SpyStore struct {
//	response  string
//	cancelled bool
//	t         *testing.T
//}

//func (s SpyStore) assertWasCancelled() {
//	s.t.Helper()
//	if !s.cancelled {
//		s.t.Error("store was not told to cancel")
//	}
//}
//
//func (s SpyStore) assertWasNotCancelled() {
//	s.t.Helper()
//	if s.cancelled {
//		s.t.Error("store was told to cancel")
//	}
//}
//
//func (s *SpyStore) Fetch() string {
//	time.Sleep(100 * time.Millisecond)
//	return s.response
//}
//
//func (s *SpyStore) Cancel() {
//	s.cancelled = true
//}
