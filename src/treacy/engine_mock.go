package treacy

import (
	"github.com/gin-gonic/gin"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"reflect"
	"runtime"
)

type (
	EngineMock struct {
		running bool
		routes  []RouteInfo
		port    string
	}

	RouteInfo struct {
		Method  string
		Path    string
		Handler string
	}
)

func (e *EngineMock) Run(port string) error {
	e.running, e.port = true, port
	return nil
}

func (e *EngineMock) POST(relativePath string, handlers ...gin.HandlerFunc) {
	for _, handler := range handlers {
		e.routes = append(e.routes, RouteInfo{Method: "POST", Path: relativePath, Handler: funcName(handler)})
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

func (e *EngineMock) Routes() []RouteInfo {
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
	return &routeMatcher{Method: method}
}

type routeMatcher struct {
	Method  string
	Path    string
	Handler string
}

func (matcher *routeMatcher) On(path string) *routeMatcher {
	matcher.Path = path
	return matcher
}

func (matcher *routeMatcher) By(handler string) *routeMatcher {
	matcher.Handler = handler
	return matcher
}

func (matcher *routeMatcher) Expected() *RouteInfo {
	return &RouteInfo{
		Method:  matcher.Method,
		Path:    matcher.Path,
		Handler: matcher.Handler}
}

func (matcher *routeMatcher) Match(actual interface{}) (success bool, err error) {
	containElementMatcher := &matchers.ContainElementMatcher{
		Element: RouteInfo{
			Method:  matcher.Method,
			Path:    matcher.Path,
			Handler: matcher.Handler}}
	return (containElementMatcher).Match(actual.(*EngineMock).Routes())
}

func (matcher *routeMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*EngineMock).Routes(), "to include", matcher.Expected())
}

func (matcher *routeMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual.(*EngineMock).Routes(), "not to include", matcher.Expected())
}
