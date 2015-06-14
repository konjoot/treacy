package mocks

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	. "matchers"
	"reflect"
	"runtime"
	"strings"
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

func (e *EngineMock) FmtRoutes() (s string) {
	s += "[\n\t"
	s += strings.Join(e.routes, "\n\t")
	s += "\n]"
	return
}

func (e *EngineMock) String() string {
	return fmt.Sprintf("&EngineMock{running:\"%t\"}", e.running)
}

// BeRunning matcher
func BeRunning() *BaseMatcher {
	return Matcher(&isRunningMatcher{})
}

type isRunningMatcher struct{}

func (_ *isRunningMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (_ *isRunningMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*EngineMock).IsRunning()
}

func (_ *isRunningMatcher) Format(actual interface{}) string {
	return actual.(*EngineMock).String()
}

func (_ *isRunningMatcher) Message() string {
	return "to be running"
}

// Handle route matcher
func Handle(method string) *handleMatcher {
	return &handleMatcher{method: method}
}

type handleMatcher struct {
	method  string
	path    string
	handler string
}

func (m *handleMatcher) On(path string) *handleMatcher {
	m.path = path
	return m
}

func (m *handleMatcher) By(handler string) *handleMatcher {
	m.handler = handler
	return m
}

func (m *handleMatcher) String() string {
	return fmt.Sprintf("%s %s -> %s", m.method, m.path, m.handler)
}

func (m *handleMatcher) Message() string {
	return fmt.Sprintf("to include \"%s\"", m)
}

func (m *handleMatcher) Match(actual interface{}) (success bool, err error) {
	containElementMatcher := &matchers.ContainElementMatcher{Element: m.String()}
	return (containElementMatcher).Match(actual.(*EngineMock).Routes())
}

func (m *handleMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %#s\n\t%s", actual.(*EngineMock).FmtRoutes(), m.Message())
}

func (m *handleMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %s\n\tnot %s", actual.(*EngineMock).FmtRoutes(), m.Message())
}
