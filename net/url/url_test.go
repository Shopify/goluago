package url_test

import (
	"github.com/Shopify/goluago/luatest"
	"testing"
)

func TestLuaURL(t *testing.T) { luatest.RunLuaTests(t, "url_test.lua") }
