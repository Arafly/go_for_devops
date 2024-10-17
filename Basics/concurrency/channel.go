// Channels are a safe way of transferring data between goroutines without using a mutex. You can send data to a channel in one goroutine, then consume data from the same channel in another goroutine. By default, a channel does not have space to store data, so you must simultaneously send and receive data from the channel to avoid a deadlock. A buffered channel allows you to allocate space to temporarily store data in the channel.

package concurrency

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 50)
	go count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

// func main() {
// 	ch := make(chan string, 1)

// 	go func() {
// 		for _, word := range []string{"hello", "world"} {
// 			ch <- word
// 		}
// 		close(ch)
// 	}()

// 	for word := range ch {
// 		fmt.Println(word)
// 	}
// }