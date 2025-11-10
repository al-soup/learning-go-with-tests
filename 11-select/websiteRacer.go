package main

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	// With `select` the first channel wins and execution continues with the first case that is ready.
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) (ch chan bool) {
	// ch = make(chan struct{}) would be more memory efficient since nothing needs to actually be allocated
	// In this case the type which is sent to the channel does not matter. We just need a signal, so struct would be sufficient.
	// Always `make` channels. If you use `var ch chan struct{}` it will be nil because the variable would be implemented as "zero" value
	// <- would block forever because you cannot send nil to a channel
	ch = make(chan bool)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return
}
