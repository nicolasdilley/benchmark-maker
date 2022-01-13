package main

import "sync"
import "time"

func main() {
	var mu sync.Mutex

	go func() {
		mu.Lock()
	}()
	mu.Lock()
}