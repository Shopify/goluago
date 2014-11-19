package env

import (
	"github.com/Shopify/go-lua"
	"os"
)

// Open makes part of the Go os env functions available to Lua code executing
// in the given lua.State, provided that it requires it:
//    local url = require("goluago/env")
func Open(l *lua.State) {
	envOpen := func(l *lua.State) int {
		lua.NewLibrary(l, fmtLibrary)
		return 1
	}
	lua.Require(l, "goluago/env", envOpen, false)
	lua.Pop(l, 1)
}

var fmtLibrary = []lua.RegistryFunction{
	{"getenv", getenv},
	{"setenv", setenv},
}

func getenv(l *lua.State) int {
	key := lua.CheckString(l, 1)
	val := os.Getenv(key)
	lua.PushString(l, val)
	return 1
}

func setenv(l *lua.State) int {
	key := lua.CheckString(l, 1)
	val := lua.CheckString(l, 2)
	err := os.Setenv(key, val)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}
	return 0
}
