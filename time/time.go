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
	ns := lua.CheckInteger(l, 1)
	time.Sleep(time.Nanosecond * time.Duration(ns))
	return 1
}

func now(l *lua.State) int {
	lua.PushNumber(l, float64(time.Now().UnixNano()))
	return 1
}

func since(l *lua.State) int {
	start := lua.CheckNumber(l, 1)
	diff := float64(time.Now().UnixNano()) - start
	lua.PushNumber(l, diff)
	return 1
}
