package mocks

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"reflect"
	"runtime"
)

type EngineMock struct {
	running bool
	routes  []string
	port    string
}

func (e *EngineMock) Run(port string) error {
	e.running, e.port = true, port
	return nil
}

func (e *EngineMock) POST(relativePath string, handlers ...gin.HandlerFunc) {
	for _, handler := range handlers {
		e.routes = append(e.routes, fmt.Sprintf("%s %s -> %s", "POST", relativePath, funcName(handler)))
	}
}

func (e *EngineMock) PUT(relativePath string, handlers ...gin.HandlerFunc) {
	for _, handler := range handlers {
		e.routes = append(e.routes, fmt.Sprintf("%s %s -> %s", "PUT", relativePath, funcName(handler)))
	}
}

func (e *EngineMock) GET(relativePath string, handlers ...gin.HandlerFunc) {
	for _, handler := range handlers {
		e.routes = append(e.routes, fmt.Sprintf("%s %s -> %s", "GET", relativePath, funcName(handler)))
	}
}

func (e *EngineMock) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	for _, handler := range handlers {
		e.routes = append(e.routes, fmt.Sprintf("%s %s -> %s", "DELETE", relativePath, funcName(handler)))
	}
}

func funcName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func (e *EngineMock) IsRunning() bool {
	return e.running
}

func (e *EngineMock) Port() string {
	return e.port
}

func (e *EngineMock) Routes() []string {
	return e.routes
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

// Handle route matcher
func Handle(method string) *routeMatcher {
	return &routeMatcher{method: method}
}

type routeMatcher struct {
	method  string
	path    string
	handler string
}

func (matcher *routeMatcher) On(path string) *routeMatcher {
	matcher.path = path
	return matcher
}

func (matcher *routeMatcher) By(handler string) *routeMatcher {
	matcher.handler = handler
	return matcher
}

func (matcher *routeMatcher) ToString() string {
	return fmt.Sprintf("%s %s -> %s", matcher.method, matcher.path, matcher.handler)
}

func (matcher *routeMatcher) Match(actual interface{}) (success bool, err error) {
	containElementMatcher := &matchers.ContainElementMatcher{Element: matcher.ToString()}
	return (containElementMatcher).Match(actual.(*EngineMock).Routes())
}

func (matcher *routeMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*EngineMock).Routes(), "to include", matcher.ToString())
}

func (matcher *routeMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*EngineMock).Routes(), "not to include", matcher.ToString())
}
