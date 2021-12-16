func bug() {
	ch := make(chan int)
	<-ch
}

func fixed() {
	ch := make(chan int)

	go func(){
		ch <- 0
	}
	<-ch
}