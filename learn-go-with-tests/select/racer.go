package _select

import (
	"fmt"
	"net/http"
	"time"
)

//func Racer(a, b string) (winner string) {
//
//	startA := time.Now()
//	http.Get(a)
//	aDuration := time.Since(startA)
//
//	startB := time.Now()
//	http.Get(a)
//	bDuration := time.Since(startB)
//
//	if aDuration > bDuration {
//		return b
//	}
//
//	return a
//}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

//func Racer(a, b string) (winner string) {
//	aDuration := measureResponseTime(a)
//	bDuration := measureResponseTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//
//	return b
//}

// func Racer(a, b string) (winner string, err error) {
//
//		select {
//		case <-ping(a):
//			return a, nil
//		case <-ping(b):
//			return b, nil
//		case <-time.After(10 * time.Second):
//			return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
//		}
//	}
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {

	ch := make(chan struct{})

	go func() {

		http.Get(url)
		close(ch)

	}()

	return ch
}
