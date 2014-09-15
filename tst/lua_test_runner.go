package tst

import (
	"github.com/Shopify/go-lua"
	"reflect"
	"testing"

	luatesting "github.com/Shopify/goluago/pkg/testing"
)

func RunLuaTests(t *testing.T, libraryCallback func(l *lua.State), filename string) {
	requireCallback := func(l *lua.State) {
		lua.Register(l, "istrue", isTrue(t))
		lua.Register(l, "isfalse", isFalse(t))
		lua.Register(l, "equals", isEqual(t))
		lua.Register(l, "notequals", isNotEqual(t))
		libraryCallback(l)
	}
	luatesting.RunLuaTestFile(t, requireCallback, filename)
}

func isTrue(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		name := lua.CheckString(l, 1)
		cond := lua.ToBoolean(l, 2)

		if !cond {
			t.Errorf("%s: condition should be true", name)
		}

		return 0
	}
}

func isFalse(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		name := lua.CheckString(l, 1)
		cond := lua.ToBoolean(l, 2)

		if cond {
			t.Errorf("%s: condition should be false", name)
		}

		return 0
	}
}

func isEqual(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		name := lua.CheckString(l, 1)
		want := lua.ToValue(l, 2)
		got := lua.ToValue(l, 3)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("%s: want `%#v`, got `%#v`", name, want, got)
		}

		return 0
	}
}

func isNotEqual(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		name := lua.CheckString(l, 1)
		want := lua.ToValue(l, 2)
		got := lua.ToValue(l, 3)

		if reflect.DeepEqual(want, got) {
			t.Errorf("%s: dont want `%#v` but got it", name, want)
		}

		return 0
	}
}
