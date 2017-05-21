package emoji_test

import (
	"github.com/markstgodard/currency-service/emoji"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Emoji Converter", func() {

	Context("when currency is USD", func() {

		It("converts the amount to dollar emoji", func() {
			v, err := emoji.Convert(100.00, "USD")
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("💵"))
		})
	})

	Context("when currency is in Britsh Pounds", func() {

		It("converts the amount to pound emoji", func() {
			v, err := emoji.Convert(100.00, "GBP")
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("💷"))
		})
	})

	Context("when currency is in CAD", func() {

		It("converts the amount to poop emoji", func() {
			v, err := emoji.Convert(100.00, "CAD")
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("💩"))
		})
	})

	Context("when currency is unknown", func() {
		It("converts the amount to poop emoji", func() {
			v, err := emoji.Convert(100.00, "FOO")
			Expect(err).To(MatchError(ContainSubstring("Unknown currency code: FOO")))
			Expect(v).To(Equal("🤷"))
		})
	})

})
