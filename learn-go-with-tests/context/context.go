package context

import (
	"fmt"
	"net/http"
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

func Server(store Store) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprintln(writer, d)
		case <-ctx.Done():
			store.Cancel()
		}

	}

}

type Store interface {
	Fetch() string
	Cancel()
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}
