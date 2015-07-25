package gjson_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGjson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gjson Suite")
}
