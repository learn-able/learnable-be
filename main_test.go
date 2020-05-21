package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	got := main()
	want := "Hello Learnable"

	if got != want {
		fmt.Errorf("\nGot: %s\nWant: %s", got, want)
	}
}
