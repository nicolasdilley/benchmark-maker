var bound int = *

func haddock() {
	ch := make(chan int)

	for i := 0; i < bound; i++ {
		go func(){
			ch <- 0
		}()

		<-ch
	}
}