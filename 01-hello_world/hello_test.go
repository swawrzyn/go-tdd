package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Stefan", "english")
		want := "Hello, Stefan!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Polish", func(t *testing.T) {
		got := Hello("Marek", "polish")
		want := "Witam, Marek!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in 中文", func(t *testing.T) {
		got := Hello("李锐", "中文")
		want := "你好, 李锐!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
