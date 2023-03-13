package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every two seconds"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// printWords spins off a gourinte that prints words off input channels
// until we receive a signal on exitCh. wg can be used to know that
// printWords exits.
func printWords(in1, in2 chan string, exit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-exit:
			fmt.Println("exiting")
			return
		case str := <-in1:
			fmt.Println("in1: ", str)
		case str := <-in2:
			fmt.Println("in2: ", str)
		}
	}
}

// func main() {
// 	in1 := make(chan string)
// 	in2 := make(chan string)
// 	wg := &sync.WaitGroup{}
// 	exit := make(chan struct{})

// 	wg.Add(1)
// 	go printWords(in1, in2, exit, wg)

// 	in1 <- "hello"
// 	in2 <- "world"

// 	close(exit)

// 	wg.Wait()
// }