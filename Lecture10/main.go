package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	counter int
	lastElm string
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	cp.Add(2)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()
			if cp.counter == times {
				cp.Unlock()
				break
			}
			if cp.lastElm != "foo" {
				fmt.Print("foo")
				cp.lastElm = "foo"
				cp.counter++
			}
			cp.Unlock()
			time.Sleep(1000 * time.Millisecond)
		}

	}()
	//	cp.Done()
}

func (cp *ConcurrentPrinter) printBar(times int) {
	//cp.Add(1)
	go func() {
		defer cp.Done()
		for {
			cp.Lock()
			if cp.counter == times {
				cp.Unlock()
				break
			}
			if cp.lastElm != "bar" {
				fmt.Print("bar")
				cp.lastElm = "bar"
				cp.counter++
			}
			cp.Unlock()
		}

	}()
	//	cp.Done()
}

func main() {

	times := 10
	cp := &ConcurrentPrinter{}
	cp.printFoo(times)
	cp.printBar(times)
	cp.Wait()
}
