package gjson

import (
	"fmt"
	"io"
	"reflect"

	"github.com/antonholmquist/jason"
	"github.com/onsi/gomega/format"
)

type Int64Matcher struct {
	Keys     []string
	Expected int64
}

func (m *Int64Matcher) Match(a interface{}) (bool, error) {
	var obj *jason.Object
	var err error
	switch a := a.(type) {
	case []byte:
		obj, err = jason.NewObjectFromBytes(a)
	case string:
		obj, err = jason.NewObjectFromBytes([]byte(a))
	case io.Reader:
		obj, err = jason.NewObjectFromReader(a)
	case *jason.Object:
		obj = a
	default:
		return false, fmt.Errorf("Cant create jason.Object")
	}
	if err != nil {
		return false, err
	}

	actual, err := obj.GetInt64(m.Keys...)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(actual, m.Expected), nil
}

func (m *Int64Matcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to match json", m.Expected)
}

func (m *Int64Matcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to match FieldError", m.Expected)
}
