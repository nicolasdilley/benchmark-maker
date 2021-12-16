func bug() {
	ch := make(chan int)
	ch <- 0
}

func fixed() {
	ch := make(chan int, 1)
	ch <- 0
}