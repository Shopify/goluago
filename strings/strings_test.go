package strings_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/luatest"
	"github.com/Shopify/goluago/strings"
	"testing"
)

func TestLuaStrings(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	strings.Open(l)

	luatest.RunLuaTests(t, l, "strings_test.lua")
}
