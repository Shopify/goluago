package url

import (
	"github.com/Shopify/go-lua"
	"net/url"
)

// Open makes part of the Go net/url package available to Lua code executing
// in the given lua.State, provided that it requires it:
//    local url = require("goluago/net/url")
func Open(l *lua.State) {
	urlOpen := func(l *lua.State) int {
		lua.NewLibrary(l, urlLibrary)
		return 1
	}
	lua.Require(l, "goluago/net/url", urlOpen, false)
	l.Pop(1)
}

var urlLibrary = []lua.RegistryFunction{
	{"parse", parse},
}

func parse(l *lua.State) int {
	rawurl := lua.CheckString(l, 1)
	u, err := url.Parse(rawurl)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}
	pushURL(l, u)
	return 1
}

func pushURL(l *lua.State, u *url.URL) {

	l.NewTable()

	var urlFunc = map[string]func(*url.URL) lua.Function{
		"isAbs":      urlIsAbs,
		"parse":      urlParse,
		"requestURI": urlRequestURI,
		"string":     urlString,
	}

	for name, goFn := range urlFunc {
		l.PushGoFunction(goFn(u))
		l.SetField(-2, name)
	}

	l.NewTable()

	getHook := func(l *lua.State) int {
		key := lua.CheckString(l, 2)
		switch key {
		case "scheme":
			l.PushString(u.Scheme)
		case "opaque":
			l.PushString(u.Opaque)
		case "host":
			l.PushString(u.Host)
		case "path":
			l.PushString(u.Path)
		case "rawQuery":
			l.PushString(u.RawQuery)
		case "fragment":
			l.PushString(u.Fragment)
		default:
			return 0
		}
		return 1
	}

	l.PushGoFunction(getHook)
	l.SetField(-2, "__index")

	setHook := func(l *lua.State) int {
		key := lua.CheckString(l, 2)
		val := lua.CheckString(l, 3)
		switch key {
		case "scheme":
			u.Scheme = val
		case "opaque":
			u.Opaque = val
		case "host":
			u.Host = val
		case "path":
			u.Path = val
		case "rawQuery":
			u.RawQuery = val
		case "fragment":
			u.Fragment = val
		default:
			l.RawSet(1)
		}
		return 0
	}

	l.PushGoFunction(setHook)
	l.SetField(-2, "__newindex")

	l.SetMetaTable(-2)
}

func urlIsAbs(u *url.URL) lua.Function {
	return func(l *lua.State) int {
		l.PushBoolean(u.IsAbs())
		return 1
	}
}

func urlParse(u *url.URL) lua.Function {
	return func(l *lua.State) int {
		newU, err := u.Parse(lua.CheckString(l, 1))
		if err != nil {
			lua.Errorf(l, err.Error())
			panic("unreachable")
		}
		pushURL(l, newU)
		return 1
	}
}

func urlRequestURI(u *url.URL) lua.Function {
	return func(l *lua.State) int {
		l.PushString(u.RequestURI())
		return 1
	}
}

func urlString(u *url.URL) lua.Function {
	return func(l *lua.State) int {
		l.PushString(u.String())
		return 1
	}
}
