package main_test

import (
	"fmt"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWcwhen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wcwhen Suite")
}

var pathToSprocketCLI string

var _ = Describe("Scraper", func() {
	BeforeSuite(func() {
		var err error
		pathToSprocketCLI, err = gexec.Build("github.com/McBrideR/wcwhen")
		fmt.Println(pathToSprocketCLI)
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("foo", func() {
		command := exec.Command(pathToSprocketCLI)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("mru"))
		Eventually(session).Should(gexec.Exit(0))
	})
})

