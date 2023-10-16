package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Doli")
	got := buffer.String()
	want := "Hello, Doli"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
