package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

// TODO https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context
// the context package helps you manage long-running processes (e.g. web-servers)
// The context package provides functions to derive new Context values from existing ones.
// These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		// race the two asynchronous processes of fetch and possible cancel
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		// when context is "done" or "cancelled"
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

// TODO continue here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context#refactor
