// Go provides the Writer Interface to raed and write data to a stream or local disk.
package main

import (
	"fmt"
	"io"
	"os"
)

// go doc io.Reader
// go doc os.File
// go doc os.File.Read

// // The os.Open function, on the other hand, opens an existing file for reading. If the file does not exist, os.Open will return an error.
// func main() {
// 	file, err := os.Open("goWrite.txt")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer file.Close()

//     reader := io.Reader(file)
//     buffer := make([]byte, 30)
//     count, err := reader.Read(buffer)
//     if err != nil {
//         log.Fatal(err)
//         return
//     }
// 	fmt.Printf("Read count=%v\n err=%v\n buffer=%v\n", count, err, string(buffer))
// }

// The following source snippet opens an existing file and creates a copy of its content using the io.Copy function.
// One common, and recommended practice to notice is the deferred call to the method Close on the file. This ensures a graceful release of OS resources when the function exits:

func main() {
	f1, err := os.Open("goWrite.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer f1.Close()

	f2, err := os.Create("./file0.bkp")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer f2.Close()

	n, err := io.Copy(f2, f1)
	if err != nil {
		fmt.Println("Failed to copy:", err)
		os.Exit(1)
	}

	fmt.Printf("Copied %d bytes from %s to %s\n", n, f1.Name(), f2.Name())
}




