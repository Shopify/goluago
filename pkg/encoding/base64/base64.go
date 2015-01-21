package base64

import (
	"encoding/base64"
	"github.com/Shopify/go-lua"
)

func Open(l *lua.State) {
	base64Open := func(l *lua.State) int {
		lua.NewLibrary(l, base64Library)
		return 1
	}
	lua.Require(l, "goluago/encoding/base64", base64Open, false)
	lua.Pop(l, 1)
}

var base64Library = []lua.RegistryFunction{
	{"encode", encode},
	{"decode", decode},
}

func encode(l *lua.State) int {
	data := lua.CheckString(l, 1)
	lua.PushString(l, base64.StdEncoding.EncodeToString([]byte(data)))
	return 1
}

func decode(l *lua.State) int {
	data := lua.CheckString(l, 1)
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}
	lua.PushString(l, string(decoded))
	return 1
}
