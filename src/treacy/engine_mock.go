package treacy

import (
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
)

type EngineMock struct {
	running bool
	port    string
}

func (e *EngineMock) Run(port string) error {
	e.running, e.port = true, port
	return nil
}

func (e *EngineMock) IsRunning() bool {
	return e.running
}

func (e *EngineMock) Port() string {
	return e.port
}

// BeRunning matcher
func BeRunning() *isRunningMatcher {
	return &isRunningMatcher{}
}

type isRunningMatcher struct{}

func (matcher *isRunningMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.BeTrueMatcher{}).Match(actual.(*EngineMock).IsRunning())
}

func (matcher *isRunningMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be running")
}

func (matcher *isRunningMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be running")
}
