package fmt_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/fmt"
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaFmt(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	fmt.Open(l)
	luatest.RunLuaTests(t, l, "fmt_test.lua")
}
