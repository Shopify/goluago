package strings

import (
	"github.com/Shopify/go-lua"
	"strings"
)

func Open(l *lua.State) {
	strOpen := func(l *lua.State) int {
		lua.NewLibrary(l, stringLibrary)
		return 1
	}
	lua.Require(l, "goluago/strings", strOpen, false)
	lua.Pop(l, 1)
}

var stringLibrary = []lua.RegistryFunction{
	{"split", split},
	{"trim", trim},
	{"replace", replace},
}

func split(l *lua.State) int {
	str := lua.CheckString(l, 1)
	sep := lua.CheckString(l, 2)

	strArr := strings.Split(str, sep)

	lua.CreateTable(l, len(strArr), 0)
	for i, strVal := range strArr {
		lua.PushString(l, strVal)
		lua.RawSetInt(l, -2, i+1)
	}

	return 1
}

func trim(l *lua.State) int {
	str := lua.CheckString(l, 1)
	lua.PushString(l, strings.TrimSpace(str))
	return 1
}

func replace(l *lua.State) int {
	s := lua.CheckString(l, 1)
	old := lua.CheckString(l, 2)
	new := lua.CheckString(l, 3)
	n := lua.CheckInteger(l, 4)

	lua.PushString(l, strings.Replace(s, old, new, n))
	return 1
}
