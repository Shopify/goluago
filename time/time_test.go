package time_test

import (
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaTime(t *testing.T) { luatest.RunLuaTests(t, "time_test.lua") }
