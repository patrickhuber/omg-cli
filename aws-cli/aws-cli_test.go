package awscli_test

import (
	"gopkg.in/urfave/cli.v2"
	awscli "github.com/enaml-ops/omg-cli/aws-cli"
	"github.com/enaml-ops/pluginlib/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("given the photon cli", func() {
	Context("when called with a complete set of flags", func() {
		It("then it should NOT panic", func() {
			action := awscli.GetAction(func(s string) {})
			var ctx *cli.Context
			ctx = pluginutil.NewContext([]string{"someapp",
				"--aws-instance-size", "some",
				"--aws-subnet", "stuff",
				"--aws-pem-path", "to",
				"--aws-access-key", "do",
				"--aws-keyname", "92895-35-2975340-34346346",
				"--aws-secret", "10.0.0.3",
				"--gateway", "10.0.0.254",
				"--cidr", "10.0.0.1/24",
			}, awscli.GetFlags())
			err := action(ctx)
			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Context("when called with an incomplete set of flags", func() {
		It("then it should panic and exit", func() {
			action := awscli.GetAction(func(s string) {})
			var ctx *cli.Context
			ctx = pluginutil.NewContext([]string{"someapp"}, awscli.GetFlags())
			err := action(ctx)
			Ω(err).Should(HaveOccurred())
		})
	})
})
