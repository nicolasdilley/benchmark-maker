import "sync"

func bug() {
	var wg *sync.WaitGroup

	wg.Done()
}

func fixed() {
	var wg *sync.WaitGroup

	wg.Add(1)
	wg.Done()
}