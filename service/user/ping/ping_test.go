package ping_test

import (
	"fmt"
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

func TestPingPong(t *testing.T) {
	got := ping.PingPong()
	want := "PingPong"

	fmt.Println(got)
	fmt.Println(want)

	if got != want {
		t.Errorf("PingPong was incorrect, got: %s, want: %s", got, want)
	}
}
