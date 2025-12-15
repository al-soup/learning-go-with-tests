package sync

import "sync"

// NOTE analyze your code with `go vet`. Also good in a pipeline

type Counter struct {
	count int
	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	mu sync.Mutex
}

// Constructor which shows readers of your API that it would be better to not initialize the type yourself
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// while a goroutine acquires the lock, all other will have to wait until the lock is unlocked
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

// NOTE Use channels when passing ownership of data
// NOTE Use mutexes for managing state
