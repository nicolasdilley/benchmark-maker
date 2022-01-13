func bug() {
	ch := make(chan int)

	close(ch)
	ch <- 0
}

func fixed() {
	ch := make(chan int)

	ch <- 0
	close(ch)
}