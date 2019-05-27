package engine

import (
	"testing"
	"time"
)

func myChan(in chan int, t *testing.T)  {
	for i := 0; i< 10; i++ {
		t.Errorf("T' m not blocked")
		in <- 1
	}
}

var ch = time.Tick(time.Microsecond)

func readChanFromOne(t *testing.T)  {
	for i := 0; i< 10; i++ {
		go func(n int, t *testing.T) {
			<-ch
			t.Logf("No #%d", n)
		}(i, t)
	}
}

func TestChannel (t *testing.T) {
	timer := time.After(time.Second * 5)
	readChanFromOne(t)
	<-timer
}