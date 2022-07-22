package dev07

import (
	"sync"
)

func MergeChannels(channels ...<-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(locChan <-chan interface{}) {
			defer wg.Done()
			for value := range locChan {
				ch <- value
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
