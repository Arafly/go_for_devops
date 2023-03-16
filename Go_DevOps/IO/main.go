// Go provides the Writer Interface to write data to a stream or local disk.
package main

import (
	"fmt"
	"io"
	"os"
)

// type Writer interface {
// 	Write (p []byte) (n int, err error)
// }

func main() {
	file, err := os.Create("goWrite.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    writer := io.Writer(file)
    data, err := writer.Write([]byte("Hello World"))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(data)

}

