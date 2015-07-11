package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPassrServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PassrServer Suite")
}
