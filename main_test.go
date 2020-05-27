package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	assertPorts := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("\nGot: %s\nWant: %s", got, want)
		}
	}

	t.Run("Test that it will set a default port number", func(t *testing.T) {
		got := getPort()
		want := ":4747"

		assertPorts(t, got, want)
	})

	t.Run("Test that it can set a port", func(t *testing.T){
		os.Setenv("PORT", "8080")

		got := getPort()
		want := ":8080"

		assertPorts(t, got, want)
	})
}
