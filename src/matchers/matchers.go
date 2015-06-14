package matchers

import (
	"fmt"
	"github.com/onsi/gomega/types"
)

type MatcherInterface interface {
	Matcher() types.GomegaMatcher
	Prepare(actual interface{}) interface{}
	Format(actual interface{}) string
	Message() string
}

func Matcher(m MatcherInterface) *BaseMatcher {
	return &BaseMatcher{m}
}

type BaseMatcher struct{ MatcherInterface }

func (m *BaseMatcher) Match(actual interface{}) (success bool, err error) {
	return m.Matcher().Match(m.Prepare(actual))
}

func (m *BaseMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %s\n\t%s", m.Format(actual), m.Message())
}

func (m *BaseMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %s\n\tnot %s", m.Format(actual), m.Message())
}
