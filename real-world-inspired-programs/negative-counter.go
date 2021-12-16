package main

import "sync"

// adapted from https://github.com/kubernetes/kubernetes/blob/d70ee902fddc682863a3cc4f0d8eac0223ebf70b/test/e2e/storage/vsphere/nodemapper.go\#L62
func main() {
	var wg sync.WaitGroup /*<\label{line:neg-wg-decl}>*/
	someList := []int{1, 2, 3}
	for range someList { /*<\label{line:neg-for-range}>*/
		go func() { /*<\label{line:neg-wg-go}>*/
			wg.Done() // may trigger a run-time error /*<\label{line:neg-wg-done}>*/
		}()
		wg.Add(1) /*<\label{line:neg-wg-add}>*/
	}
	wg.Wait() /*<\label{line:neg-wg-wait}>*/
}
