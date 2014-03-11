package json

import (
	"encoding/json"
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/utils"
)

func Open(l *lua.State) {
	jsonOpen := func(l *lua.State) int {
		lua.NewLibrary(l, jsonLibrary)
		return 1
	}
	lua.Require(l, "goluago/encoding/json", jsonOpen, false)
	lua.Pop(l, 1)
}

var jsonLibrary = []lua.RegistryFunction{
	{"unmarshal", unmarshal},
}

func unmarshal(l *lua.State) int {
	payload := lua.CheckString(l, 1)

	var output interface{}

	if err := json.Unmarshal([]byte(payload), &output); err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}
	return utils.DeepPush(l, &output)
}
