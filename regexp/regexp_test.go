package regexp_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/luatest"
	"github.com/Shopify/goluago/regexp"
	"testing"
)

func TestLuaRegexp(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	regexp.Open(l)
	luatest.RunLuaTests(t, l, "regexp_test.lua")
}
