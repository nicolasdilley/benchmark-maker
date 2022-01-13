import "sync"

func bug() {
	var mu *sync.Mutex

	mu.Lock()
	mu.Lock()
}

func fixed() {
	var mu *sync.Mutex

	mu.Lock()
}