package json_test

import (
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaJSON(t *testing.T) { luatest.RunLuaTests(t, "json_test.lua") }
