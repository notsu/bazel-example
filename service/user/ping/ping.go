package ping

import (
	pong "github.com/notsu/bazel-example/service/post/pong"
)

// Ping a request
func Ping() string {
	return "Ping 2"
}

// PingPong returns a combination string with Ping and Pong
func PingPong() string {
	return Ping() + pong.Pong()
}
