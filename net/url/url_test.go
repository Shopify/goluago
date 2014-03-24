package url_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago"
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaURL(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	goluago.Open(l)
	luatest.RunLuaTests(t, l, "url_test.lua")
}
