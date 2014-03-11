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

func replace(l *lua.State) int {
	s, ok := lua.ToString(l, -4)
	if !ok {
		lua.Errorf(l, "replace: 1st argument (source) must be a string")
		panic("unreachable")
	}
	old, ok := lua.ToString(l, -3)
	if !ok {
		lua.Errorf(l, "replace: 2nd argument (old) must be a string")
		panic("unreachable")
	}
	new, ok := lua.ToString(l, -2)
	if !ok {
		lua.Errorf(l, "replace: 3rd argument (new) must be a string")
		panic("unreachable")
	}
	n, ok := lua.ToInteger(l, -1)
	if !ok {
		lua.Errorf(l, "replace: 4th argument (n) must be an integer")
		panic("unreachable")
	}

	lua.PushString(l, strings.Replace(s, old, new, n))
	return 1
}
