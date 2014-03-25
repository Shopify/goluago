package url

import (
	"github.com/Shopify/go-lua"
	"net/url"
)

func Open(l *lua.State) {
	urlOpen := func(l *lua.State) int {
		lua.NewLibrary(l, urlLibrary)
		return 1
	}
	lua.Require(l, "goluago/net/url", urlOpen, false)
	lua.Pop(l, 1)
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

	lua.NewTable(l)

	var urlFunc = map[string]func(*url.URL) lua.Function{
		"isAbs":      urlIsAbs,
		"parse":      urlParse,
		"requestURI": urlRequestURI,
		"string":     urlString,
	}

	for name, goFn := range urlFunc {
		lua.PushGoFunction(l, goFn(u))
		lua.SetField(l, -2, name)
	}

	lua.NewTable(l)

	getHook := func(l *lua.State) int {
		key := lua.CheckString(l, 2)
		switch key {
		case "scheme":
			lua.PushString(l, u.Scheme)
		case "opaque":
			lua.PushString(l, u.Opaque)
		case "host":
			lua.PushString(l, u.Host)
		case "path":
			lua.PushString(l, u.Path)
		case "rawQuery":
			lua.PushString(l, u.RawQuery)
		case "fragment":
			lua.PushString(l, u.Fragment)
		default:
			return 0
		}
		return 1
	}

	lua.PushGoFunction(l, getHook)
	lua.SetField(l, -2, "__index")

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
			lua.RawSet(l, 1)
		}
		return 0
	}

	lua.PushGoFunction(l, setHook)
	lua.SetField(l, -2, "__newindex")

	lua.SetMetaTable(l, -2)
}

func urlIsAbs(u *url.URL) lua.Function {
	return func(l *lua.State) int {
		lua.PushBoolean(l, u.IsAbs())
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
		lua.PushString(l, u.RequestURI())
		return 1
	}
}

func urlString(u *url.URL) lua.Function {
	return func(l *lua.State) int {
		lua.PushString(l, u.String())
		return 1
	}
}
