package mocks

import (
	"fmt"
	"github.com/onsi/gomega/matchers"
	"net/http/httptest"
	"reflect"
)

func Resource() *ResourceMock {
	return &ResourceMock{form: &resourceFormMock{}}
}

type ResourceMock struct {
	created bool
	form    *resourceFormMock
}

type resourceFormMock struct {
	Name string `binding:"required"`
	Desc string `binding:"required"`
}

func (f *resourceFormMock) ValOf(name string) string {
	r := reflect.ValueOf(f)
	return reflect.Indirect(r).FieldByName(name).String()
}

func (r *ResourceMock) Form() *resourceFormMock {
	return r.form
}

func (r *ResourceMock) Save() (ok bool) {
	ok, r.created = true, true
	return
}

func (r *ResourceMock) Url() string {
	return "/tests/1"
}

func (r *ResourceMock) IsCreated() bool {
	return r.created
}

func (r *ResourceMock) IsBindedWith(m map[string]string) (ok bool) {
	form := r.Form()
	for key, val := range m {
		if ok = form.ValOf(key) == val; !ok {
			return
		}
	}
	return
}

func (r *ResourceMock) String() string {
	return fmt.Sprintf("&ResourceMock{ created: %t }", r.created)
}

// BeCreated matcher
func BeCreated() *beCreatedMatcher {
	return &beCreatedMatcher{}
}

type beCreatedMatcher struct{}

func (m *beCreatedMatcher) Message() string {
	return "to be created"
}

func (m *beCreatedMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.BeTrueMatcher{}).Match(actual.(*ResourceMock).IsCreated())
}

func (m *beCreatedMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %s\n\t%s", actual.(*ResourceMock).String(), m.Message())
}

func (m *beCreatedMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %s\n\tnot %s", actual.(*ResourceMock).String(), m.Message())
}

// HaveHeader matcher
func HaveHeader(name string) *haveHeaderMatcher {
	return &haveHeaderMatcher{Name: name}
}

type haveHeaderMatcher struct {
	Name     string
	Expected string
}

func (m *haveHeaderMatcher) With(url string) *haveHeaderMatcher {
	m.Expected = url
	return m
}

func (m *haveHeaderMatcher) Header(r interface{}) string {
	return r.(*httptest.ResponseRecorder).Header().Get(m.Name)
}

func (m *haveHeaderMatcher) Message() string {
	return fmt.Sprintf("to have header \"%s\" with \"%s\"", m.Name, m.Expected)
}

func (m *haveHeaderMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.EqualMatcher{Expected: m.Expected}).Match(m.Header(actual))
}

func (m *haveHeaderMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %#v\n\t%s", actual.(*httptest.ResponseRecorder).HeaderMap, m.Message())
}

func (m *haveHeaderMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %#v\n\tnot %s", actual.(*httptest.ResponseRecorder).HeaderMap, m.Message())
}

// BindedWith matcher
func BindedWith(expected map[string]string) *bindedWithMatcher {
	return &bindedWithMatcher{Expected: expected}
}

type bindedWithMatcher struct {
	Expected map[string]string
}

func (m *bindedWithMatcher) Message() string {
	return fmt.Sprintf("to have been binded with %#v", m.Expected)
}

func (m *bindedWithMatcher) Match(actual interface{}) (success bool, err error) {
	return (&matchers.BeTrueMatcher{}).Match(actual.(*ResourceMock).IsBindedWith(m.Expected))
}

func (m *bindedWithMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %#v\n\t%s", actual.(*ResourceMock).Form(), m.Message())
}

func (m *bindedWithMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %#v\n\tnot %s", actual.(*ResourceMock).Form(), m.Message())
}
