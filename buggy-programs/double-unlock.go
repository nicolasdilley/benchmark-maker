import "sync"

func bug() {
	var mu *sync.Mutex

	mu.Lock()
	mu.Unlock()
	mu.Unlock()
}

func fixed() {
	var mu *sync.Mutex

	mu.Lock()
	mu.Unlock()
}