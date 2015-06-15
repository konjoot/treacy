package mocks

import (
	"fmt"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	. "matchers"
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
func BeCreated() *BaseMatcher {
	return Matcher(&beCreatedMatcher{})
}

type beCreatedMatcher struct{}

func (_ *beCreatedMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (_ *beCreatedMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*ResourceMock).IsCreated()
}

func (_ *beCreatedMatcher) Format(actual interface{}) string {
	return actual.(*ResourceMock).String()
}

func (_ *beCreatedMatcher) Message() string {
	return "to be created"
}

func (_ *beCreatedMatcher) String() (s string) {
	return
}

// HaveHeader matcher
func HaveHeader(name string) *haveHeaderMatcher {
	return &haveHeaderMatcher{Name: name}
}

type haveHeaderMatcher struct {
	Name     string
	Expected string
}

func (m *haveHeaderMatcher) With(url string) *BaseMatcher {
	m.Expected = url
	return Matcher(m)
}

func (m *haveHeaderMatcher) Matcher() types.GomegaMatcher {
	return &matchers.EqualMatcher{Expected: m.Expected}
}

func (m *haveHeaderMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*httptest.ResponseRecorder).Header().Get(m.Name)
}

func (_ *haveHeaderMatcher) Format(actual interface{}) string {
	return fmt.Sprintf("%#v", actual.(*httptest.ResponseRecorder).HeaderMap)
}

func (m *haveHeaderMatcher) Message() string {
	return "to have header"
}

func (m *haveHeaderMatcher) String() string {
	return fmt.Sprintf("\"%s\" with \"%s\"", m.Name, m.Expected)
}

// BindedWith matcher
func BindedWith(expected map[string]string) *BaseMatcher {
	return Matcher(&bindedWithMatcher{Expected: expected})
}

type bindedWithMatcher struct {
	Expected map[string]string
}

func (m *bindedWithMatcher) Matcher() types.GomegaMatcher {
	return &matchers.BeTrueMatcher{}
}

func (m *bindedWithMatcher) Prepare(actual interface{}) interface{} {
	return actual.(*ResourceMock).IsBindedWith(m.Expected)
}

func (_ *bindedWithMatcher) Format(actual interface{}) string {
	return fmt.Sprintf("%#v", actual.(*ResourceMock).Form())
}

func (_ *bindedWithMatcher) Message() string {
	return "to have been binded with"
}

func (m *bindedWithMatcher) String() string {
	return fmt.Sprintf("%#v", m.Expected)
}
