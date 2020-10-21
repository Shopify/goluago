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
	l.Pop(1)
}

var timeLibrary = []lua.RegistryFunction{
	{"now", now},
	{"nowUTC", nowUTC},
	{"sleep", sleep},
	{"since", since},
	{"addTime", addTime}
}

func nowUTC(l *State) int {
	l.PushString(time.Now().UTC().Format(time.RFC3339))
	return 1
}},

{"addHourUTC", func(l *State) int{
	start := lua.CheckInteger(l, 1)
	hourToAdd := lua.CheckInteger(l, 2)
	diff := start.Add(time.Hour * time.Duration(hourToAdd))
	l.PushString(diff.Format(time.RFC3339))
	return 1
}},

func sleep(l *lua.State) int {
	ns := lua.CheckInteger(l, 1)
	time.Sleep(time.Nanosecond * time.Duration(ns))
	return 1
}

func now(l *lua.State) int {
	l.PushNumber(float64(time.Now().UnixNano()))
	return 1
}

func since(l *lua.State) int {
	start := lua.CheckNumber(l, 1)
	diff := float64(time.Now().UnixNano()) - start
	l.PushNumber(diff)
	return 1
}
