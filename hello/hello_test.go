package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// t.Run("saying hello to people", func(t *testing.T) {
	// 	got := Hello("Tanu")
	// 	want := "hello, Tanu"
	// 	assertCorrectMessage(t, got, want)
	// })
	t.Run("say 'hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ben", "Spanish")
		want := "hola, Ben"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Gom", "French")
		want := "bonjour, Gom"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
