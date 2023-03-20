// The following shows the type alphaReader, a trivial implementation of the io.Reader that filters out non-alpha characters from its string source:

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// type alphaReader string

// func (a alphaReader) Read(p []byte) (int, error) {
// 	count := 0
// 	for i := 0; i < len(a); i++ {
// 		if (a[i] >= 'A' && a[i] <= 'Z') ||
// 			(a[i] >= 'a' && a[i] <= 'z') {
// 			p[i] = a[i]
// 		}
// 		count++
// 	}
// 	return count, io.EOF
// }

// func main() {
// 	str := alphaReader("Hello! Where is the sun?")
// 	io.Copy(os.Stdout, &str)
// 	fmt.Println()
// }

// The following snippet shows an updated version of alphaReader. This time, it takes an io.Reader as its source as shown in the following code:

type alphaReader struct {
	src io.Reader
}

func NewAlphaReader(source io.Reader) *alphaReader {
	return &alphaReader{source}
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	count, err := a.src.Read(p) // p has now source data
	if err != nil {
		return count, err
	}
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') ||
			(p[i] >= 'a' && p[i] <= 'z') {
			continue
		} else {
			p[i] = 0
		}
	}
	return count, io.EOF
}

func main() {
	str := strings.NewReader("Hello! Where is the sun?")
	alpha := NewAlphaReader(str)
	io.Copy(os.Stdout, alpha)
	fmt.Println()
}

// The main change to note in this version of the code is that the alphaReader type is now a struct which embeds an io.Reader value. When alphaReader.Read() is invoked, it calls the wrapped reader as a.src.Read(p), which will inject the source data into byte slice p. Then the method loops through p and applies the filter to the data. Now, to use the alphaReader, it must first be provided with an existing reader which is facilitated by the NewAlphaReader() constructor function.


// For instance, the following code snippet shows how the alphaReader type can now be combined with an os.File to filter out non-alphabetic characters from a file (the Go source code itself).

// func main() { 
//    file, _ := os.Open("./reader0.go") 
//    alpha := NewAlphaReader(file) 
//    io.Copy(os.Stdout, alpha) 
//    fmt.Println() 
// } 

