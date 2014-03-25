package fmt_test

import (
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaFmt(t *testing.T) { luatest.RunLuaTests(t, "fmt_test.lua") }
