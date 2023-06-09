// Go provides the Writer Interface to raed and write data to a stream or local disk.
package main

import (
	"fmt"
	"io"
	"os"
)

// go doc io.Writer
// go doc os.File
// go doc os.File.Write

// The os.Create function creates a new file with the specified path. If the file already exists, os.Create will overwrite it
func main() {
	file, err := os.Create("goWrite.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    writer := io.Writer(file)
    data, err := writer.Write([]byte("Hello World from Reader and Writer"))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(data)

}




