package ping_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	ping "github.com/notsu/bazel-example/service/user/ping"
)

var _ = Describe("Ping", func() {
	Context("when ping", func() {
		It("should ping", func() {
			got := ping.Ping()
			want := "Ping"

			Expect(got).To(Equal(want))
		})
	})

	Context("when ping", func() {
		It("should ping", func() {
			got := ping.PingPong()
			want := "PingPong"

			Expect(got).To(Equal(want))
		})
	})
})
