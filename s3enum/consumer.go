package s3enum

import (
    "sync"
)

// Consumer struct
type Consumer struct {
    inputChannel  chan string
    resultChannel chan string
    quit          chan bool
    resolver      Resolver
    wg            *sync.WaitGroup
}

// NewConsumer initializer
func NewConsumer(resolver Resolver, input chan string, result chan string, quit chan bool, wg *sync.WaitGroup) *Consumer {
    consumer := &Consumer{
	resolver:      resolver,
	inputChannel:  input,
	resultChannel: result,
	quit:          quit,
	wg:            wg,
    }

    return consumer
}

// Consume reads messages from 'input', and outputs results to 'result'.
func (c *Consumer) Consume() {
    defer c.wg.Done() // This ensures that the wait group is done when this goroutine finishes

    for {
	j, more := <-c.inputChannel
	if more {
	    if c.resolver.IsBucket(j) {
		c.resultChannel <- j
	    }
	} else {
	    return
	}
    }
}
