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
	String() string
}

func Matcher(m MatcherInterface) *BaseMatcher {
	return &BaseMatcher{m}
}

type BaseMatcher struct{ MatcherInterface }

func (m *BaseMatcher) Match(actual interface{}) (success bool, err error) {
	return m.Matcher().Match(m.Prepare(actual))
}

func (m *BaseMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf(m.Template(false), m.Format(actual), m.Message())
}

func (m *BaseMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf(m.Template(true), m.Format(actual), m.Message())
}

func (m *BaseMatcher) Template(negate bool) (s string) {
	s = "Expected\n\t%s\n"

	if negate {
		s += "not "
	}

	s += "%s"

	if str := m.String(); len(str) > 0 {
		s += fmt.Sprintf("\n\t%s", str)
	}

	return
}
