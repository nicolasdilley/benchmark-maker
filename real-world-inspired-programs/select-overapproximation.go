package main

import (
	"errors"
	"sync"
)

func buildCommitters() error {
	namespaces := []int{0, 1, 2, 3, 4} // code simplified

	// for each namespace, we build multiple committers (based on maxBatchSize per namespace)
	var wg sync.WaitGroup
	nsCommittersChan := make(chan []int, len(namespaces))
	defer close(nsCommittersChan)
	errsChan := make(chan error, len(namespaces))
	defer close(errsChan)

	// for each namespace, we build committers in parallel. This is because,
	// the committer building process requires fetching of missing revisions
	// that in turn, we want to do in parallel
	for _, ns := range namespaces {
		wg.Add(1)
		go func(ns int) { /*<\label{line:select-go}>*/
			defer wg.Done()
			err := errors.Errorf("database name is illegal, cannot be empty")
			if err != nil {
				errsChan <- err /*<\label{line:send-errsChan}>*/
				return
			}
			nsCommittersChan <- committers
		}(ns)
	}
	wg.Wait()

	// collect all committers
	var allCommitters []int
	select {
	case err := <-errsChan: /*<\label{line:select-rcv}>*/
		return nil, errors.WithStack(err)
	default:
		for i := 0; i < len(namespaces); i++ {
			allCommitters = append(allCommitters, <-nsCommittersChan...) /*<\label{line:select-rcv-1}>*/
		}
	}

	return allCommitters, nil
}
