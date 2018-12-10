package main

import (
	"fmt"

	ping "github.com/notsu/bazel-example/service/user/ping"
)

func main() {
	fmt.Println(ping.Ping())
}
