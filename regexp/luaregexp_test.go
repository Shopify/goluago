package regexp

import (
	"github.com/Shopify/go-lua"
	"io/ioutil"
	"testing"
)

func TestLuaRegexp(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/regexptest.lua")
	if err != nil {
		t.Fatalf("reading lua test script, %v", err)
	}
	testScript := string(data)

	l := lua.NewState()

	lua.OpenLibraries(l)
	Register(l)

	failHook := func(l *lua.State) int {
		str, ok := lua.ToString(l, -1)
		if !ok {
			t.Fatalf("need a string on the lua stack for calls to fail()")
		}
		lua.Pop(l, 1)
		t.Error(str)
		return 0
	}
	lua.Register(l, "fail", failHook)
	wantTop := lua.Top(l)

	if err := lua.LoadString(l, testScript); err != nil {
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
