package main

import "testing"

func TestGreeting(t *testing.T) {
	t.Run("greet user in japanese", func(t *testing.T) {
		got := Greeting("Yuta-kun", "")
		want := "Konnichiwa Yuta-kun"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet user in spanish", func(t *testing.T) {
		got := Greeting("Yuta-kun", "spanish")
		want := "Hola Yuta-kun"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet user in french", func(t *testing.T) {
		got := Greeting("Yuta-kun", "french")
		want := "Bonjour Yuta-kun"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Konnichiwa' when an empty string is supplied", func(t *testing.T) {
		got := Greeting("", "")
		want := "Konnichiwa"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
