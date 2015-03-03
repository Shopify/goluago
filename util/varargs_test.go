package util

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/Shopify/go-lua"
)

func TestVarargsFrom1(t *testing.T) {
	l := lua.NewState()
	l.PushString("foo")
	l.PushString("bar")
	l.PushString("baz")

	testVarargs(t, l, 1, []interface{}{"foo", "bar", "baz"})
}

func TestVarargsFrom2(t *testing.T) {
	l := lua.NewState()
	l.PushString("foo")
	l.PushString("bar")
	l.PushString("baz")

	testVarargs(t, l, 2, []interface{}{"bar", "baz"})
}

func TestVarargsFromEnd(t *testing.T) {
	l := lua.NewState()
	l.PushString("foo")
	l.PushString("bar")
	l.PushString("baz")

	testVarargs(t, l, 4, []interface{}{})
}

func TestVarargsFromBeyondEnd(t *testing.T) {
	l := lua.NewState()
	l.PushString("foo")
	l.PushString("bar")
	l.PushString("baz")

	testVarargs(t, l, 5, []interface{}{})
}

func TestVarargsWithNil(t *testing.T) {
	l := lua.NewState()
	l.PushNil()

	testVarargs(t, l, 1, []interface{}{nil})
}

func TestVarargsWithBoolean(t *testing.T) {
	l := lua.NewState()
	l.PushBoolean(true)
	l.PushBoolean(false)

	testVarargs(t, l, 1, []interface{}{true, false})
}

func TestVarargsWithNumber(t *testing.T) {
	l := lua.NewState()
	l.PushNumber(-20.5)
	l.PushNumber(0)
	l.PushNumber(10000)

	testVarargs(t, l, 1, []interface{}{-20.5, 0., 10000.})
}

func TestVarargsWithString(t *testing.T) {
	l := lua.NewState()
	l.PushString("foo")
	l.PushString("")

	testVarargs(t, l, 1, []interface{}{"foo", ""})
}

func TestVarargsWithTable(t *testing.T) {
	l := lua.NewState()
	l.NewTable()
	l.PushString("bar")
	l.SetField(-2, "foo")

	testVarargs(t, l, 1, []interface{}{map[string]interface{}{"foo": "bar"}})
}

func TestVarargsWithFunction(t *testing.T) {
	l := lua.NewState()
	l.PushGoClosure(func(l *lua.State) int { return 0 }, 0)

	actual, err := PullVarargs(l, 1)
	if err != nil {
		t.Fatalf("Failed, %s", err.Error())
	}
	switch actual[0].(type) {
	case lua.Function:
		return
	default:
		t.Fatalf("Expected function, got %#v", actual[0])
	}
}

func TestVarargsWithUserData(t *testing.T) {
	userdata := regexp.MustCompile(".")
	l := lua.NewState()
	l.PushUserData(userdata)

	testVarargs(t, l, 1, []interface{}{userdata})
}

func TestVarargsWithThread(t *testing.T) {
	l := lua.NewState()
	l.PushThread()

	testVarargs(t, l, 1, []interface{}{l})
}

func testVarargs(t *testing.T, l *lua.State, index int, expected []interface{}) {
	actual, err := PullVarargs(l, index)
	if err != nil {
		t.Fatalf("Failed, %s", err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %#v, got %#v", expected, actual)
	}
}
