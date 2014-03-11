package json

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/go-lua"
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

	var recurseOnArray func([]interface{})

	var recurseOnMap func(map[string]interface{})

	forwardOnType := func(val interface{}) {

		switch val := val.(type) {
		case nil:
			lua.PushNil(l)

		case bool:
			lua.PushBoolean(l, val)

		case string:
			lua.PushString(l, val)

		case uint8:
			lua.PushNumber(l, float64(val))
		case uint16:
			lua.PushNumber(l, float64(val))
		case uint32:
			lua.PushNumber(l, float64(val))
		case uint64:
			lua.PushNumber(l, float64(val))
		case uint:
			lua.PushNumber(l, float64(val))

		case int8:
			lua.PushNumber(l, float64(val))
		case int16:
			lua.PushNumber(l, float64(val))
		case int32:
			lua.PushNumber(l, float64(val))
		case int64:
			lua.PushNumber(l, float64(val))
		case int:
			lua.PushNumber(l, float64(val))

		case float32:
			lua.PushNumber(l, float64(val))
		case float64:
			lua.PushNumber(l, val)

		case []interface{}:
			lua.CreateTable(l, len(val), 0)
			recurseOnArray(val)

		case map[string]interface{}:
			lua.CreateTable(l, 0, len(val))
			recurseOnMap(val)

		default:
			lua.Errorf(l, fmt.Sprintf("payload contains unsupported type: %T", val))
			panic("unreachable")
		}
	}

	recurseOnMap = func(input map[string]interface{}) {
		// -1 is a table
		for key, val := range input {
			lua.PushString(l, key)
			// -1: key, -2: table
			forwardOnType(val)
			// -1: something, -2: key, -3: table
			lua.RawSet(l, -3)
		}
	}

	recurseOnArray = func(input []interface{}) {
		// -1 is a table
		for i, val := range input {
			forwardOnType(val)
			// -1: something, -2: table
			lua.RawSetInt(l, -2, i+1)
		}
	}

	forwardOnType(output)

	return 1
}
