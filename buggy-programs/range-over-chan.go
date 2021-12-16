func bug() {
	ch := make(chan int)
	
	go func() {
		ch <- 0
	}

	for range ch {

	}
}

func fixed() {
	ch := make(chan int)
	
	go func() {
		ch <- 0
		close(ch)
	}
	
	for range ch {

	}
}