// The following code snippet shows the implementation of the channelWriter type, a writer that decomposes and serializes its stream that is sent over a Go channel as consecutive bytes:

package main

import (
	"fmt"
	"io"
	"os"
)

type channelWriter struct {
	Channel chan byte
}

func NewChannelWriter() *channelWriter {
	return &channelWriter{
		Channel: make(chan byte, 1024),
	}
}

// The Write method uses a goroutine to copy each byte, from p, and sends it across the c.Channel. Upon completion, the goroutine closes the channel so that consumers are notified when to stop consuming from the channel. As an implementation convention, writers should not modify slice p or hang on to it. When an error occurs, the writer should return the current number of bytes processed and an error.

func (c *channelWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	go func() {
		defer close(c.Channel) // when done
		for _, b := range p {
			c.Channel <- b
		}
	}()

	return len(p), nil
}

// The following snippet uses the fmt.Fprint function to serialize the "Stream me!" string as a sequence of bytes over a channel using channelWriter:
// func main() {
// 	cw := NewChannelWriter()
// 	go func() {
// 		fmt.Fprint(cw, "Stream me!")
// 	}()

// 	for c := range cw.Channel {
// 		fmt.Printf("%c\n", c)
// 	}
// }

// The following snippet shows another example where the content of a file is serialized over a channel using the same channelWriter. In this implementation, an io.File value and io.Copy function are used to source the data instead of the fmt.Fprint function:
func main() {
	cw := NewChannelWriter()
	file, err := os.Open("./write.go")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	_, err = io.Copy(cw, file)
	if err != nil {
		fmt.Println("Error copying:", err)
		os.Exit(1)
	}

	// consume channel
	for c := range cw.Channel {
		fmt.Printf("%c\n", c)
	}
}
