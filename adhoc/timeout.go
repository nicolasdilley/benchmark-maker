func adhoc() {
	ch := make(chan int)
	timeout := make(chan int)
	go func() {
		select {
		case <-ch:
			timeout <- 0
		case <-time.After(5 * time.Second):
			timeout <- 0
		}
	}()

	<-timeout
}