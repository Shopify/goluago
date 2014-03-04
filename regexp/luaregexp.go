package regexp

import (
	"github.com/Shopify/go-lua"
	"regexp"
)

// Register makes the regexp function available to Lua code in the `regexp`
// namespace, once Lua code invokes `require('regexp')`.
func Register(l *lua.State) {
	reOpen := func(l *lua.State) int {
		lua.NewLibrary(l, regexpLibrary)
		return 1
	}
	lua.Require(l, "regexp", reOpen, true)
	lua.Pop(l, 1)
}

var regexpLibrary = []lua.RegistryFunction{
	{"match", Match},
	{"quotemeta", QuoteMeta},
}

// Match checks whether a textual regular expression matches a string. More
// complicated queries need to use Compile and the full Regexp interface.
func Match(l *lua.State) int {
	pattern, ok := lua.ToString(l, -2)
	if !ok {
		lua.Errorf(l, "match: 1st arg (pattern) must be a string")
		lua.Pop(l, 2)
		return 0
	}
	bStr, ok := lua.ToString(l, -1)
	if !ok {
		lua.Errorf(l, "match: 2nd arg (s) must be a string")
		lua.Pop(l, 2)
		return 0
	}

	matched, err := regexp.MatchString(pattern, bStr)
	if err != nil {
		lua.Errorf(l, err.Error())
		return 0
	}
	lua.PushBoolean(l, matched)
	return 1
}

// QuoteMeta returns a string that quotes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression
// matching the literal text. For example, QuoteMeta(`[foo]`) returns `\[foo\]`.
func QuoteMeta(l *lua.State) int {
	s, ok := lua.ToString(l, -1)
	if !ok {
		lua.Errorf(l, "expected s to be a string")
		lua.Pop(l, 1)
		return 0
	}
	quoted := regexp.QuoteMeta(s)
	lua.PushString(l, quoted)
	return 1
}
