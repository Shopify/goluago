package time_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/luatest"
	"github.com/Shopify/goluago/time"
	"io/ioutil"
	"testing"
)

func TestLuaTime(t *testing.T) {
	l := lua.NewState()

	lua.OpenLibraries(l)
	time.Open(l)

	data, err := ioutil.ReadFile("time_test.lua")
	if err != nil {
		t.Fatalf("loading test file, %v", err)
	}

	luatest.RunLuaTests(t, l, string(data))
}
