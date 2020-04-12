package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Peter")
	want := "Hello, Peter"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}