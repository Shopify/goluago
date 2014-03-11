package time_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/luatest"
	"github.com/Shopify/goluago/time"
	"testing"
)

func TestLuaTime(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	time.Open(l)

	luatest.RunLuaTests(t, l, "time_test.lua")
}
