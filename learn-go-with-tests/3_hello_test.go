package main

import "testing"

func TestHello3(t *testing.T) {

	t.Run("Not Empty", func(t *testing.T) {
		got := Hello1("ali")
		want := "Hello ali"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty", func(t *testing.T) {
		got := Hello2("")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {

	//  t.Helper() the line number reported will be in our function call rather than inside our test helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
