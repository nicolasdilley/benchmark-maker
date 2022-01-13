func adhoc() {
	ch := make(chan int)
	doneCh := make(chan int, bound)

	for i := 0; i < bound; i++ {
		go func() {
			doneCh <- 0
		}()
	}

	select {
	case <-doneCh:
	default:
		for i := 0; i < bound; i++ {
			<-ch
		}
	}

}