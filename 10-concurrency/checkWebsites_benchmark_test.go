package concurrency

import (
	"testing"
	"time"
)

func stubSlowWebsiteChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

// run benchmarks with `go test -bench=.`
func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = "a url"
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(stubSlowWebsiteChecker, urls)
	}
}
