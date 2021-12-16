var bound int = *

func haddock(){
 
 for i := 0; i < bound; i++ {
 	ch := make(chan int)

 	go func(){
 		<-ch
 	}()

 	ch <- 0
 }
}