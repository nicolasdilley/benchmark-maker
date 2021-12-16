func haddock(){
	ch := make(chan int)

	defer func(){
		<-ch
	}

	go func(){
		ch <- 0
	}()
}