package regexp_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/luatest"
	"github.com/Shopify/goluago/regexp"
	"io/ioutil"
	"testing"
)

func TestLuaRegexp(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	regexp.Open(l)

	data, err := ioutil.ReadFile("regexp_test.lua")
	if err != nil {
		t.Fatalf("loading test file, %v", err)
	}

	luatest.RunLuaTests(t, l, string(data))
}
