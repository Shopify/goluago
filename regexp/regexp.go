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
	{"compile", compile},
}

func match(l *lua.State) int {
	pattern := lua.CheckString(l, 1)
	s := lua.CheckString(l, 2)

	matched, err := regexp.MatchString(pattern, s)
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

func compile(l *lua.State) int {
	expr := lua.CheckString(l, 1)
	re, err := regexp.Compile(expr)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	lua.NewTable(l)
	for name, goFn := range regexpFunc {
		// -1: tbl
		lua.PushGoFunction(l, goFn(re))
		// -1: fn, -2:tbl
		lua.SetField(l, -2, name)
	}

	return 1
}

var regexpFunc = map[string]func(*regexp.Regexp) lua.Function{
	"findAll":         reFindAll,
	"findAllSubmatch": reFindAllSubmatch,
}

func reFindAll(re *regexp.Regexp) lua.Function {
	return func(l *lua.State) int {
		s := lua.CheckString(l, 1)
		n := lua.CheckInteger(l, 2)

		all := re.FindAllString(s, n)

		lua.CreateTable(l, len(all), 0)
		for i, str := range all {
			lua.PushString(l, str)
			lua.RawSetInt(l, -2, i+1)
		}

		return 1
	}
}

func reFindAllSubmatch(re *regexp.Regexp) lua.Function {
	return func(l *lua.State) int {
		s := lua.CheckString(l, 1)
		n := lua.CheckInteger(l, 2)

		allSubmatch := re.FindAllStringSubmatch(s, n)

		lua.CreateTable(l, len(allSubmatch), 0)
		for i, submatches := range allSubmatch {
			// -1: outTbl
			lua.CreateTable(l, len(submatches), 0)
			// -1: inTbl, -2:outTbl
			for j, str := range submatches {
				lua.PushString(l, str)
				// -1: str, -2: inTbl, -3: outTbl
				lua.RawSetInt(l, -2, j+1)
				// -1: inTbl, -2:outTbl
			}
			lua.RawSetInt(l, -2, i+1)
			// -1:outTbl
		}

		return 1
	}
}
