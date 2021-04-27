package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMesage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		// tengo
		got := Hello("Eduardo", "")
		// quiero
		want := "Hello, Eduardo"
		assertCorrectMesage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		// tengo
		got := Hello("", "")
		// quiero
		want := "Hello, world"
		assertCorrectMesage(t, got, want)
	})

	t.Run("say 'Hola, ' when language is Spanish", func(t *testing.T) {
		got := Hello("Eduardo", "Spanish")
		want := "Hola, Eduardo"
		assertCorrectMesage(t, got, want)
	})

	t.Run("say 'Bojour, ' when language is french", func(t *testing.T) {
		got := Hello("Eduardo", "French")
		want := "Bonjour, Eduardo"
		assertCorrectMesage(t, got, want)
	})

	t.Run("say 'Konishiwa' when language is Japan", func(t *testing.T) {
		got := Hello("Eduardo", "Japan")
		want := "Konishiwa, Eduardo"
		assertCorrectMesage(t, got, want)
	})
}
