package pong_test

import (
	"testing"

	pong "github.com/notsu/bazel-example/service/post/pong"
)

func TestPong(t *testing.T) {
	got := pong.Pong()
	want := "Pong"

	if got != want {
		t.Errorf("Pong was incorrect, got: %s, want: %s", got, want)
	}
}
