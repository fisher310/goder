package atomic

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	expected := 2
	for i := 0; i < 10000; i++ {
		var a atomicInt
		a.increment()
		ch := make(chan int)
		go func(c chan int) {
			a.increment()
			c <- 0
		}(ch)
		<-ch
		if a.get() != expected {
			t.Errorf("expected value is %d, but the actual is %d\n", expected, a)
		}
	}
}
