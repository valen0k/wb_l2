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
	<-MergeChannels(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	since := time.Since(start)
	if since > time.Second+2*time.Millisecond {
		t.Errorf("%v > %v", since, time.Second+2*time.Millisecond)
	}
}
