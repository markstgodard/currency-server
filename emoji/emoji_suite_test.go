package emoji_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEmoji(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Emoji Suite")
}
