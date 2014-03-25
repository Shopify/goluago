package strings_test

import (
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaStrings(t *testing.T) { luatest.RunLuaTests(t, "strings_test.lua") }
