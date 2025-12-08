package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// Channels are a Go data structure that can both receive and send values.
	// These operations, along with their details, allow communication between different processes
	resultChannel := make(chan result)

	for _, url := range urls {
		// operation that does not block in Go will run in a separate process called a goroutine
		// To tell Go to start a new goroutine we turn a function call into a go statement
		// by putting the keyword go in front of it: go doSomething(). (often surrounded by an
		// anonymous function)
		// In this case we are passing url as a parameter rather than using a reference to it to
		// avoid a "closure variable capture problem.": in the loop the url might still point to
		// the last value in the loop and all goroutines would use this last url. Goroutine creation
		// happens during the loop (very fast) but goroutine execution happens later when the Go
		// scheduler runs them
		go func(url string) {
			// sending a `result` struct for each call to `wc` to the `resultChannel` with a "send statement".
			resultChannel <- result{url, wc(url)}
		}(url) // IIFE

	}

	for i := 0; i < len(urls); i++ {
		// this "receive expression" assigns values received from a channel to a variable
		// it will block execution until it receives the next value.
		// If we would wait for a value that never comes (e.g. `len(urls)+1`) it would
		// block the loop on the last iteration.
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
