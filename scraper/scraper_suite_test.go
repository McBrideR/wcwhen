package scraper_test

import (
	. "github.com/McBrideR/wcwhen/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestScraper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scraper Suite")
}

var _ = Describe("Scraper", func() {
	Context("when supplying a country", func() {
		It("should respond with that country that is entered", func() {
			GetMatchStats("Ireland")
			Expect(os.Stdout).To(Equal("Ireland"))
		})
	})
})
