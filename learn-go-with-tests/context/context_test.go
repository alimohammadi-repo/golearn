package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//func TestServer(t *testing.T) {
//
//	data := "hello, world"
//	svr := Server(&StubStore{data})
//
//	request := httptest.NewRequest(http.MethodGet, "/", nil)
//	response := httptest.NewRecorder()
//
//	svr.ServeHTTP(response, request)
//
//	if response.Body.String() != data {
//		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
//	}
//
//}

// func TestServer(t *testing.T) {
//
//	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
//		data := "hello, world"
//		store := &SpyStore{response: data}
//		svr := Server(store)
//		request := httptest.NewRequest(http.MethodGet, "/", nil)
//
//		cancellingCtx, cancel := context.WithCancel(request.Context())
//		time.AfterFunc(5*time.Millisecond, cancel)
//		request = request.WithContext(cancellingCtx)
//
//		response := httptest.NewRecorder()
//
//		svr.ServeHTTP(response, request)
//
//		if !store.cancelled {
//			t.Error("store was not told to cancel")
//		}
//
//	})
//
//	t.Run("returns data from store", func(t *testing.T) {
//		data := "hello, world"
//		store := &SpyStore{response: data}
//		svr := Server(store)
//
//		request := httptest.NewRequest(http.MethodGet, "/", nil)
//		response := httptest.NewRecorder()
//
//		svr.ServeHTTP(response, request)
//
//		if response.Body.String() != data {
//			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
//		}
//
//		if store.cancelled {
//			t.Error("it should not have cancelled the store")
//		}
//
//	})
//
// }
func TestServer(t *testing.T) {

	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		//store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})

}
