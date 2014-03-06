package strings

import (
	"github.com/Shopify/go-lua"
	"strings"
)

func Open(l *lua.State) {
	strOpen := func(l *lua.State) int {
		lua.NewLibrary(l, regexpLibrary)
		return 1
	}
	lua.Require(l, "goluago/strings", strOpen, false)
	lua.Pop(l, 1)
}

var regexpLibrary = []lua.RegistryFunction{
	{"split", split},
	{"trim", trim},
}

func split(l *lua.State) int {
	str, ok := lua.ToString(l, -2)
	if !ok {
		lua.Errorf(l, "split: 1st argument (to split) must be a string")
		panic("unreachable")
	}
	sep, ok := lua.ToString(l, -1)
	if !ok {
		lua.Errorf(l, "split: 2nd argument (separator) must be a string")
		panic("unreachable")
	}

	strArr := strings.Split(str, sep)
	lua.CreateTable(l, len(strArr), 0)
	for i, strVal := range strArr {
		lua.PushString(l, strVal)
		lua.RawSetInt(l, -2, i+1)
	}

	return 1
}

func trim(l *lua.State) int {
	str, ok := lua.ToString(l, -1)
	if !ok {
		lua.Errorf(l, "trim: must have a string argument")
		panic("unreachable")
	}
	lua.PushString(l, strings.TrimSpace(str))
	return 1
}
