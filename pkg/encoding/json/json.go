package json

import (
	"encoding/json"
	"strings"

	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/util"
)

func Open(l *lua.State) {
	jsonOpen := func(l *lua.State) int {
		lua.NewLibrary(l, jsonLibrary)
		return 1
	}
	lua.Require(l, "goluago/encoding/json", jsonOpen, false)
	l.Pop(1)
}

var jsonLibrary = []lua.RegistryFunction{
	{"marshal", marshal},
	{"unmarshal", unmarshal},
}

func check(l *lua.State, err error) {
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}
}

func marshal(l *lua.State) int {
	var t interface{}
	var err error
	var b []byte

	if !l.IsNil(1) {
		t, err = util.PullTable(l, 1)
		check(l, err)
	}

	indent := lua.OptInteger(l, 2, 0)

	if indent == 0 {
		b, err = json.Marshal(t)
	} else {
		b, err = json.MarshalIndent(t, "", strings.Repeat(" ", indent))
	}

	check(l, err)
	l.PushString(string(b))
	return 1
}

func unmarshal(l *lua.State) int {
	payload := lua.CheckString(l, 1)
	var output interface{}
	check(l, json.Unmarshal([]byte(payload), &output))
	return util.DeepPush(l, output)
}
