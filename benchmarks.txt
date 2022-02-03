
// The only way I see this working is by having 3 set of context

// One for the channels
// One for the wg
// One for the mutex

// ======== context for the channels =========


// This benchmark test whether the tools supports non-spawning bounded for loops 
// Replace * with actual bounds 
// where the bound is not known statically
func main() {

	bound := *

	ch := make(chan int)

	for i := 0; i<bound; i++ {
		_ // snippet here
	}
}


// Go in bounded for loop (statically unknown)
// This benchmark test whether the tools supports spawning bounded for loops 
// Replace * with actual bounds 
// 50% of known bounds are 10, 75% are 100, max is 11000
func main() {

	bound := *

	ch := make(chan int)

	for i := 0; i<bound; i++{
		go func() {
			_ // snippet here
		}()
	}
}


// Infinite go in for loop
// This benchmark test whether the tools supports spawning infinite for loops
func main() {

	ch := make(chan int)
	for {
		go func() {
			_ // snippet here
		}()
	}
}

//make chan in for loop
// This benchmark test whether the tools supports channel that are spawned in a bounded for loops
func main() {
	for i := 0; i < 4; i++ { {
		ch := make(chan int)
		_ // snippet here
	}
}




// Bounded channel (statically known) size of chan is 50% from empirical analysis
// this benchmark test whether the tool supports asynchronous channels with a the most used bounds
// from the empirical analysis (Q1 and mean from empirical analysis)
func main() {

	ch := make(chan int,1)

	ch <- 0
	<-ch

	_ // snippet here
	
}


// Bounded channel (statically known) size of chan is 75% from empirical analysis
// this benchmark test whether the tool supports asynchronous channels with the Q3 bounds from analysis

// Should we put them in goroutines?? 
// How can we do it with the max value from empirical analysis which is 100 000 ? Spawn them in for loop? (then we use for loops....)
func main() {

	ch := make(chan int,4)

	ch <- 0
	ch <- 0
	ch <- 0
	ch <- 0
	<-ch
	<-ch
	<-ch
	<-ch

	_ // snippet here
}


// Defer statement 
// This Benchmark test whether the tool support bugs icluded in defer statement
func main() {
	ch := make(chan int)
	
	defer func() {
		_ // snippet here
	}()
}


// Recursion statement 
// Test whether recursion is supported by the tool 
func main() {
	ch := make(chan int)
	rec(ch,10)
}

func rec(chan int,i int){

	if i > 0 {
		_ // snippet here
		rec(ch,i-1)
	}
}


// Test whether the tool supports select statements
func main() {
	ch := make(chan int)

	go func() {
		select {
		case <-ch:
			_ // snippet here
			<-ch
		}
	}()

	ch <- 0
	ch <- 0
}

// Timeout (used a lot in real world project)
// test whether the tool support timeout
func main() {
	ch := make(chan int)

	select {
	case <-time.After(3 * time.Second):
		_ // snippet here
	}
}

// Don't all the fanciness of the constructs from the channel cause it is already tested with channels

// ====== Context for the Waitgroup ====== 
import "sync"

func main(){
	var wg *sync.WaitGroup 

	_ // snippet here
}

// ====== Context for the mutex ======
import "sync"

func main(){
	var mu *sync.Mutex

	_ // snippet here
}




// CHANNEL SNIPPETS 

// One receive 
<-ch

// double close 
close(ch)
close(ch)

// send on close 
close(ch)
ch <- 0

// range over chan (without closing it)
for range ch {

}

// WG SNIPPET

// infine wait
wg.Add(1)
wg.Wait()

// negative counter
wg.Done()
wg.Done()

// MUTEX SNIPPET 

// Double Lock
mu.Lock()
mu.Lock()

// Double Unlock
mu.lock()
mu.Unlock()
mu.Unlock()

// Blocking Lock 
go func(){
	mu.Lock()
}
mu.Lock()



// ==== STANDALONE Benchmark ==== 


// chan aliasing
// This benchmark test whether the tool support channel aliasing 
// If su
func main() {

	ch1 := make(chan int)
	ch := ch1

	close(ch) // snippet here
	close(ch1) // snippet here
}