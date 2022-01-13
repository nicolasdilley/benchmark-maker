var bound int = *

func adhoc() {
	ch := make(chan int)

	for i := 0; i < bound; i++ {
		go func(){
			ch <- 0
		}()

		<-ch
	}
}