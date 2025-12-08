package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// `select` let's you wait on multiple channels (`myVar := <- ch` would be blocking call while waiting for a value)
	// in a `select` the first channel wins and execution continues with the first case that is ready.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// returns a chan an will send a signal after the defined time
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out after waiting for %s and %s", a, b)
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
