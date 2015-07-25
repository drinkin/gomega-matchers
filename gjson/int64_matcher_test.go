package gjson_test

import (
	"github.com/drinkin/gomega-matchers/gjson"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Int64Matcher", func() {
	It("", func() {
		body := `{"a":1, "b": "foo"}`

		Expect(body).To(gjson.Get("a").Int64(1))
		Expect(body).ToNot(gjson.Get("a").Int64(2))

		ok, err := gjson.Get("b").Int64(1).Match(body)
		Expect(ok).To(BeFalse())
		Expect(err.Error()).To(Equal("not a number"))
	})
})
