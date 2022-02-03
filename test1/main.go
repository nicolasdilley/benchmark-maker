package main

import (
	"sync"
)

// type = ALL

func main() {

	for i := 0; i < 100; i++ {
		var mu sync.Mutex
		mu.Lock()
		mu.Lock()
	}
}