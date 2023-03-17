// Go provides the Writer Interface to raed and write data to a stream or local disk.
package main

import (
	"fmt"
	"io"
	"os"
    "log"
)

// go doc io.Reader
// go doc os.File
// go doc os.File.Read

func main() {
	file, err := os.Open("goWrite.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    reader := io.Reader(file)
    buffer := make([]byte, 30)
    count, err := reader.Read(buffer)
    if err != nil {
        log.Fatal(err)
        return
    }
	fmt.Printf("Read count=%v\n err=%v\n buffer=%v\n", count, err, string(buffer))
}



