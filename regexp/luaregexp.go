package regexp

import (
	"github.com/Shopify/go-lua"
	"regexp"
)

// Open exposes the regexp functions to Lua code in the `goluago/regexp`
// namespace.
func Open(l *lua.State) {
	reOpen := func(l *lua.State) int {
		lua.NewLibrary(l, regexpLibrary)
		return 1
	}
	lua.Require(l, "goluago/regexp", reOpen, false)
	lua.Pop(l, 1)
}

var regexpLibrary = []lua.RegistryFunction{
	{"match", match},
	{"quotemeta", quoteMeta},
}

func match(l *lua.State) int {
	pattern, ok := lua.ToString(l, -2)
	if !ok {
		lua.Errorf(l, "match: 1st arg (pattern) must be a string")
		panic("unreachable")
	}
	bStr, ok := lua.ToString(l, -1)
	if !ok {
		lua.Errorf(l, "match: 2nd arg (s) must be a string")
		panic("unreachable")
	}

	matched, err := regexp.MatchString(pattern, bStr)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}
	lua.PushBoolean(l, matched)
	return 1
}

func quoteMeta(l *lua.State) int {
	s, ok := lua.ToString(l, -1)
	if !ok {
		lua.Errorf(l, "quotemeta: argument must be a string")
		panic("unreachable")
	}
	quoted := regexp.QuoteMeta(s)
	lua.PushString(l, quoted)
	return 1
}
