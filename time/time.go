package time

import (
	"github.com/Shopify/go-lua"
	"time"
)

func Open(l *lua.State) {
	timeOpen := func(l *lua.State) int {
		lua.NewLibrary(l, timeLibrary)
		return 1
	}
	lua.Require(l, "goluago/time", timeOpen, false)
	lua.Pop(l, 1)
}

var timeLibrary = []lua.RegistryFunction{
	{"now", now},
	{"sleep", sleep},
	{"since", since},
}

func sleep(l *lua.State) int {
	ns, ok := lua.ToInteger(l, -1)
	if !ok {
		lua.Errorf(l, "sleep: argument (nanosec) must be an integer")
		panic("unreachable")
	}

	time.Sleep(time.Nanosecond * time.Duration(ns))

	return 1
}

func now(l *lua.State) int {
	lua.PushInteger(l, int(time.Now().UnixNano()))
	return 1
}

func since(l *lua.State) int {
	start, ok := lua.ToInteger(l, -1)
	if !ok {
		lua.Errorf(l, "since: argument (nanosec) must be an integer")
		panic("unreachable")
	}

	diff := int(time.Now().UnixNano()) - start

	lua.PushInteger(l, diff)
	return 1
}
