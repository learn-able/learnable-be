package models

import (
	"testing"
)

func TestUserModel(t *testing.T) {
	t.Run("Test User Struct", func(t *testing.T) {
		user := User{}
		got := user
		want := user

		if got != want {
			t.Errorf("\nGot: %v\nWant: %v\n", got, want)
		}
	})
}
