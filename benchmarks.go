
// The only way I see this working is by having 3 set of ad hoc programs

// One for the wg
// One for the channels
// One for the mutex

// ======== Ad hoc for the channels =========

// chan aliasing
func main() {

	ch1 := make(chan int)
	ch := ch1

	for {
		_ // put snippet here
	}
}

//make chan in for loop
func main() {

	for {
	ch := make(chan int)
		_ // put snippet here
	}
}

// Infinite for loop
func main() {

	ch := make(chan int)
	for {
		_ // put snippet here
	}
}

// Bounded for loop (statically unknown)
func main() {

	ch := make(chan int)

	for i := 0; i<os.Args[1] {
		_ // put snippet here
	}
}

// Bounded for loop (statically known)
func main() {

	ch := make(chan int)
	
	for i := 0; i<10 {
		_ // put snippet here
	}
}


// Infinite go in for loop
func main() {

	ch := make(chan int)
	for {
		go func() {
			_ // put snippet here
		}()
	}
}

// Go in bounded for loop (statically unknown)
func main() {

	ch := make(chan int)

	for i := 0; i<os.Args[1] {
		go func() {
			_ // put snippet here
		}()
}

// Go in bounded for loop (statically known)
func main() {

	ch := make(chan int)
	
	for i := 0; i<10 {
		go func() {
			_ // put snippet here
		}()
	}
}

// Bounded channel (statically unknown)
func main() {

	ch := make(chan int,.Args[1])

	// put snippet here
	
}

// Bounded channel (statically known)
func main() {

	ch := make(chan int,1)
	
	// put snippet here
	
}

// Defer statement 
func main() {
	ch := make(chan int)

	defer func() {
		_ // put snippet here
	}()
}


// Recursion statement 
func main() {
	ch := make(chan int)
	recursion(ch,10)
}

func recursion(chan int,i int){

	if i > 0 {
		_ // put snippet here
		recursion(ch,i-1)
	}
}


// Timeout (used a lot in real world project)

func main() {
	ch := make(chan int)

	switch {
	case <-time.After(3 * time.Second):
		_ // put snippet here
	}
}

// Don't all the fanciness of the constructs from the channel cause it is already tested with channels

// ====== AD HOC for the Waitgroup ====== 
import "sync"

func main(){
	var wg *sync.WaitGroup 

	_ // snippet here
}

// ====== AD HOC for the mutex ======
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
