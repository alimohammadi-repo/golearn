package main

import "testing"

func TestHello1(t *testing.T) {

	got := Hello1("ali")
	want := "Hello ali"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
