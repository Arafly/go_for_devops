package main

import (
	"bytes"
	"testing"
)

func TestGreet (t *testing.T) {
	// Create a buffer to store the output
	// The Buffer is a type that implements the Writer interface
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}