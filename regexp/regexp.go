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
	pattern := lua.CheckString(l, 1)
	bStr := lua.CheckString(l, 2)

	matched, err := regexp.MatchString(pattern, bStr)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	lua.PushBoolean(l, matched)
	return 1
}

func quoteMeta(l *lua.State) int {
	s := lua.CheckString(l, 1)

	quoted := regexp.QuoteMeta(s)

	lua.PushString(l, quoted)
	return 1
}
