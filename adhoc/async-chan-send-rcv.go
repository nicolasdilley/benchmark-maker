var bound int = *


func adhoc() {

	ch := make(chan int, bound)


	for i := 0; i < bound; i++ {
		ch <- 0
	}
	

	for i := 0; i < bound; i++ {
		<-ch
	}
}