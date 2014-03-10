package strings_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/luatest"
	"github.com/Shopify/goluago/strings"
	"io/ioutil"
	"testing"
)

func TestLuaStrings(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	strings.Open(l)

	data, err := ioutil.ReadFile("strings_test.lua")
	if err != nil {
		t.Fatalf("loading test file, %v", err)
	}

	luatest.RunLuaTests(t, l, string(data))
}
