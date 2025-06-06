package main

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) (ch chan bool) {
	// ch = make(chan struct{}) would be more memory efficient since nothing needs to actually be allocated
	ch = make(chan bool)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return
}
