// type = ALL
package main 

func main() {
	ch2 := make(chan int)
	ch1 := make(chan int)

	// CP
	go func() {
		select {
		case <-ch1:
		case <-ch2;
			// CS
		}
	}()

	ch2 <- 0
}