package regexp_test

import (
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaRegexp(t *testing.T) { luatest.RunLuaTests(t, "regexp_test.lua") }
