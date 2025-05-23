package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// Channels are a Go data structure that can both receive and send values.
	// These operations, along with their details, allow communication between different processes
	// TODO resultChannel := make(chan result)

	for _, url := range urls {
		// operation that does not block in Go will run in a separate process called a goroutine
		// To tell Go to start a new goroutine we turn a function call into a go statement
		// by putting the keyword go in front of it: go doSomething(). (often surrounded by an
		// anonymous function)
		// TODO
		// go func() {
		// TODO r := <-resultChannel
		results[url] = wc(url)
		// }()
	}

	return results
}
