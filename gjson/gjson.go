package gjson

import "github.com/onsi/gomega/types"

type MatchBuilder struct {
	Keys []string
}

func Get(keys ...string) *MatchBuilder {
	return &MatchBuilder{keys}
}

func (b *MatchBuilder) Int64(v int64) types.GomegaMatcher {
	return &Int64Matcher{
		Keys:     b.Keys,
		Expected: v,
	}
}
