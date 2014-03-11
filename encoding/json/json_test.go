package json_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/encoding/json"
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaJSON(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	json.Open(l)
	luatest.RunLuaTests(t, l, "json_test.lua")
}
