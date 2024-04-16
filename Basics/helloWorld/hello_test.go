package main

import "testing"

func TestHello (t *testing.T) {
	t.Run ("saying hello to people", func (t *testing.T) {
		got := Hello("Christie", "") // Add an empty string as the second argument
		want := "Hello, Christie"
		assertCorrectMessage(t, got, want)
		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
	})

	t.Run ("say 'Hello, world!' when an empty string is supplied", func (t *testing.T) {
		got := Hello("", "") // Add an empty string as the second argument
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
		// The code below has been transferred to the assertCorrectMessage function
		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
	})

	t.Run ("in Español", func (t *testing.T) {
		got := Hello("Ines", "Español")
		want := "Hola, Ines"
		assertCorrectMessage(t, got, want)
	})

	t.Run ("in Français", func (t *testing.T) {
		got := Hello("Pierre", "Français")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage (t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}