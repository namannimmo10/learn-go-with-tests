package main

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	// `Get` method is used for requesting data
	// from a specified source or a server.
	return time.Since(start)
}

func Racer(a, b string) (winnner string) {
	aResponseTime := measureResponseTime(a)
	bResponseTime := measureResponseTime(b)

	if aResponseTime < bResponseTime {
		return a
	}

	return b
}
