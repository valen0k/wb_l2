package dev07

import (
	"testing"
	"time"
)

func TestMergeChannels(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	maxTime := 10 * time.Minute
	<-MergeChannels(
		sig(maxTime),
		sig(maxTime-2*time.Minute),
		sig(maxTime-8*time.Minute),
		sig(maxTime-4*time.Minute),
		sig(maxTime-6*time.Minute),
	)

	since := time.Since(start) - maxTime
	res := time.Millisecond * 100
	if since > res {
		t.Errorf("%v > %v", since, res)
	}
}
