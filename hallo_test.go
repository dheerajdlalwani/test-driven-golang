package main

import "testing"

func TestHallo(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hallo to people", func(t *testing.T) {
		got := Hallo("Dheeraj", "")
		want := "Hallo Dheeraj"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hallo Ingenieur' when an empty string is supplied", func(t *testing.T) {
		got := Hallo("", "")
		want := "Hallo Ingenieur"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T) {
		got := Hallo("Dheeraj", "English")
		want := "Hello Dheeraj"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hallo("Dheeraj", "French")
		want := "Bonjour Dheeraj"
		assertCorrectMessage(t, got, want)
	})
	
}
