package main

import "testing"

const helloWorld = "Hello, world"

func TestHello(t *testing.T) {
	got := Hello()
	want := helloWorld

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Hello() should return "Hello, world".
func Hello() string {
	return helloWorld
}
