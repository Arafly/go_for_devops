// Embedding files in Go binaries
// https://golang.org/pkg/embed/
// https://golang.org/pkg/embed/#hdr-Examples
// https://golang.org/pkg/embed/#hdr-Directory_Embedding

// The lines beginning with //go: are not comments, but compiler directives

package main

import (
	_ "embed"
	"fmt"
)

//go:embed goWrite.txt
var s string

func main() {
	fmt.Println(s)
}

// go:embed world.txt
// var b []byte