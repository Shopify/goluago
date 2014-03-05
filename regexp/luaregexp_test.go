package regexp_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/regexp"
	"testing"
)

func TestLuaRegexp(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	regexp.Open(l)

	failHook := func(l *lua.State) int {
		str := lua.CheckString(l, -1)
		lua.Pop(l, 1)
		t.Error(str)
		return 0
	}
	lua.Register(l, "fail", failHook)
	wantTop := lua.Top(l)

	if err := lua.LoadFile(l, "testdata/regexptest.lua", "t"); err != nil {
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
