func haddock() {
	num_messages := 5

	ch := make(chan int)
	closeCh := make(chan int)

	go send(ch, num_messages)
	rcv(ch, num_messages)

	close(closeCh)
}

func send(ch chan int, closeCh chan int) {
	select {
	case ch <- ch:
		send(ch, closeCh)
	case <-closeCh:
	}
}

func rcv(ch chan int, closeCh chan int) {
	select {
	case <-ch:
		rcv(ch, closeCh)
	case <-closeCh:
	}
}