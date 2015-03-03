package testing

import (
	"testing"

	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/util"
)

func RunLuaTestString(t *testing.T, requireCallback func(l *lua.State), luaSource string) {
	runLuaTest(t, requireCallback, func(l *lua.State) error {
		return lua.LoadString(l, luaSource)
	})
}

func RunLuaTestFile(t *testing.T, requireCallback func(l *lua.State), luaSourceFileName string) {
	runLuaTest(t, requireCallback, func(l *lua.State) error {
		return lua.LoadFile(l, luaSourceFileName, "")
	})
}

func runLuaTest(t *testing.T, requireCallback func(l *lua.State), loadCallback func(l *lua.State) error) {
	l := lua.NewState()
	lua.OpenLibraries(l)
	openTestingLibrary(l, t)
	requireCallback(l)

	wantTop := l.Top()

	if err := loadCallback(l); err != nil {
		t.Fatalf("loading lua test script in VM, %v", err)
	}

	if err := l.ProtectedCall(0, 0, 0); err != nil {
		t.Errorf("executing lua test script, %v", err)
	}

	gotTop := l.Top()
	if wantTop != gotTop {
		t.Errorf("unbalanced stack, height before %d, height after %d", wantTop, gotTop)
	}
}

func openTestingLibrary(l *lua.State, t *testing.T) {
	open := func(l *lua.State) int {
		lua.NewLibrary(l, []lua.RegistryFunction{
			{"error", error_(t)},
			{"errorf", errorf(t)},
			{"fatal", fatal(t)},
			{"fatalf", fatalf(t)},
			{"log", log(t)},
			{"logf", logf(t)},
		})
		return 1
	}
	lua.Require(l, "goluago/testing", open, false)
	l.Pop(1)
}

func error_(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		args := util.MustPullVarargs(l, 1)
		t.Error(args)
		return 0
	}
}

func errorf(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		format := lua.CheckString(l, 1)
		args := util.MustPullVarargs(l, 2)
		t.Errorf(format, args)
		return 0
	}
}

func fatal(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		args := util.MustPullVarargs(l, 1)
		t.Fatal(args)
		return 0
	}
}

func fatalf(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		format := lua.CheckString(l, 1)
		args := util.MustPullVarargs(l, 2)
		t.Fatalf(format, args)
		return 0
	}
}

func log(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		args := util.MustPullVarargs(l, 1)
		t.Log(args)
		return 0
	}
}

func logf(t *testing.T) lua.Function {
	return func(l *lua.State) int {
		format := lua.CheckString(l, 1)
		args := util.MustPullVarargs(l, 2)
		t.Logf(format, args)
		return 0
	}
}
