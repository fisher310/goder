package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d received %d \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	w := make(chan int)
	go worker(id, w)
	return w
}

func main() {
	var c1, c2 chan int = generator(), generator()

	w := createWorker(1)

	var values []int

	var activeWorker chan<- int
	var activeValue int

	tm := time.After(10 * time.Second)
	for {
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
