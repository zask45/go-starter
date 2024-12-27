package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting()
	want := "Konnichiwa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
