package main

import "testing"

func TestGreeting(t *testing.T) {
	t.Run("saying hello to other people", func(t *testing.T) {
		got := Greeting("Yuta-kun")
		want := "Konnichiwa Yuta-kun"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Konnichiwa' when an empty string is supplied", func(t *testing.T) {
		got := Greeting("")
		want := "Konnichiwa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
