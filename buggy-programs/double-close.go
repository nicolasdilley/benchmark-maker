func bug() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

func fixed() {
	ch := make(chan int)
	close(ch)
}