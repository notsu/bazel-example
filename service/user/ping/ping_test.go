package ping_test

import (
	"testing"

	ping "github.com/notsu/bazel-example/service/user/ping"
)

func TestPing(t *testing.T) {
	got := ping.Ping()
	want := "Ping"

	if got != want {
		t.Errorf("Ping was incorrect, got: %s, want: %s", got, want)
	}
}
