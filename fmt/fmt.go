package fmt

import (
	"fmt"
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/util"
)

func Open(l *lua.State) {
	fmtOpen := func(l *lua.State) int {
		lua.NewLibrary(l, fmtLibrary)
		return 1
	}
	lua.Require(l, "goluago/fmt", fmtOpen, false)
	lua.Pop(l, 1)
}

var fmtLibrary = []lua.RegistryFunction{
	{"print", pFamily(fmt.Print)},
	{"println", pFamily(fmt.Println)},
	{"sprint", sFamily(fmt.Sprint)},
	{"sprintln", sFamily(fmt.Sprintln)},
	{"printf", printf},
	{"sprintf", sprintf},
}

func pFamily(f func(a ...interface{}) (int, error)) lua.Function {
	return func(l *lua.State) int {
		n, err := f(getVarArgs(l, 1)...)
		if err != nil {
			lua.Errorf(l, err.Error())
			panic("unreachable")
		}
		return util.DeepPush(l, n)
	}
}

func sFamily(f func(a ...interface{}) string) lua.Function {
	return func(l *lua.State) int {
		return util.DeepPush(l, f(getVarArgs(l, 1)...))
	}
}

func printf(l *lua.State) int {
	format := lua.CheckString(l, 1)
	vargs := getVarArgs(l, 2)
	n, err := fmt.Printf(format, vargs...)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}
	return util.DeepPush(l, n)
}

func sprintf(l *lua.State) int {
	format := lua.CheckString(l, 1)
	vargs := getVarArgs(l, 2)
	return util.DeepPush(l, fmt.Sprintf(format, vargs...))
}

func getVarArgs(l *lua.State, from int) (vargs []interface{}) {
	for i := from; i <= lua.Top(l); i++ {
		vargs = append(vargs, lua.ToValue(l, i))
	}
	return
}
