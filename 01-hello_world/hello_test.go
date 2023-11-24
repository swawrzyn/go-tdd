package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Stefan")
	want := "Hello, Stefan!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
