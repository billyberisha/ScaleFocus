package main

import (
	"fmt"
	"sync"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) printFoo(times int) <-chan string {
	channFoo := make(chan string)
	go func() {
		for i := 0; i < times/2; i++ {
			channFoo <- "bar"
		}
		close(channFoo)
	}()
	return channFoo
}

func (cp *ConcurrentPrinter) printBar(times int) <-chan string {
	channBar := make(chan string)
	go func() {
		for i := 0; i < times/2; i++ {
			channBar <- "foo"
		}
		close(channBar)
	}()
	return channBar
}

func (cp *ConcurrentPrinter) fanIn(input1, input2 <-chan string) <-chan string {
	channFan := make(chan string)
	counter := 0
	go func() {
		for {
			if counter == 1 {
				channFan <- <-input1
				counter = 0
			} else {
				channFan <- <-input2
				counter = 1
			}

		}

	}()
	return channFan
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.printBar(times)
	cp.printFoo(times)
	for s := range cp.fanIn(cp.printBar(times), cp.printFoo(times)) {
		fmt.Print(s)
	}

}
