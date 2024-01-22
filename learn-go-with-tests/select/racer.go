package _select

import (
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

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}
