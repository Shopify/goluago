package luatest

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago"
	"reflect"
	"testing"
)

func RunLuaTests(t *testing.T, filename string) {

	l := lua.NewState()
	lua.OpenLibraries(l)
	goluago.Open(l)

	// Register the test hook
	lua.Register(l, "istrue", isTrue(t))
	lua.Register(l, "isfalse", isFalse(t))
	lua.Register(l, "equals", isEqual(t))
	lua.Register(l, "notequals", isNotEqual(t))

	// Load and exec the test program
	wantTop := lua.Top(l)

	if err := lua.LoadFile(l, filename, ""); err != nil {
		t.Fatalf("loading lua test script in VM, %v", err)
	}

	if err := lua.ProtectedCall(l, 0, 0, 0); err != nil {
		t.Errorf("executing lua test script, %v", err)
	}
	gotTop := lua.Top(l)

	if wantTop != gotTop {
		t.Errorf("Unbalanced stack!, want %d, got %d", wantTop, gotTop)
	}
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
