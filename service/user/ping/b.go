package ping

import (
	pong "github.com/notsu/bazel-example/service/post/pong"
)

func B() string {
	return "B" + pong.Hello()
}
