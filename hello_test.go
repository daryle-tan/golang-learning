package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello, world"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	got := Hello("Daryle")
	want := "Hello, Daryle"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}