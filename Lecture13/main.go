package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	/* Add other fields you might need */
	buffer chan string
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	/*Implement the rest */
	buff := make(chan string, bufferSize)
	newBufferCTX := &BufferedContext{Context: ctx, buffer: buff, CancelFunc: cancel}
	return newBufferCTX
}
func (bc *BufferedContext) Done() <-chan struct{} {

	if len(bc.buffer) == cap(bc.buffer) {
		fmt.Println("Buffer limit is reached")
		bc.CancelFunc()
	}
	return bc.Context.Done()
}
func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fn(bc, bc.buffer)
}

func main() {
	ctx := NewBufferedContext(time.Second*20, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done, Time Out reached")
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond) // try different values here
				fmt.Println("bar")
			}
		}
	})
}
