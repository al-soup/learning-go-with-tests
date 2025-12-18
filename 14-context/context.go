package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

// The context package helps you manage long-running processes (e.g. web-servers)
// The package provides functions to derive new Context values from existing ones.
// These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.
// Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
// The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue.
// When a Context is canceled, all Contexts derived from it are also canceled.
// But don't pass your values through the context! But you can put information in it like a traceID but it should never required: `context.Value` is for maintainers not users
/*
	Why context is useful in servers: When a request is canceled
	or times out, all the goroutines working on that request
	should exit quickly so the system can reclaim any resources
	they are using
*/
// More on the topic: https://go.dev/blog/context
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// context is passed through and server relies on downstream cancellations
		data, err := store.Fetch(r.Context())

		if err != nil {
			fmt.Println("Request cancelled")
			return
		}

		fmt.Fprint(w, data)
	}
}
