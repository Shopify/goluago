package json_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/encoding/json"
	"github.com/Shopify/goluago/luatest"
	"io/ioutil"
	"testing"
)

func TestLuaJSON(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	json.Open(l)

	data, err := ioutil.ReadFile("json_test.lua")
	if err != nil {
		t.Fatalf("loading test file, %v", err)
	}

	luatest.RunLuaTests(t, l, string(data))
}
