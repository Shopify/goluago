package json

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/go-lua"
)

func Open(l *lua.State) {
	jsonOpen := func(l *lua.State) int {
		lua.NewLibrary(l, regexpLibrary)
		return 1
	}
	lua.Require(l, "goluago/encoding/json", jsonOpen, false)
	lua.Pop(l, 1)
}

var regexpLibrary = []lua.RegistryFunction{
	{"unmarshall", unmarshall},
}

func unmarshall(l *lua.State) int {
	payload, ok := lua.ToString(l, -1)
	if !ok {
		lua.Errorf(l, "unmarshall: argument must be a string")
		panic("unreachable")
	}
	lua.Pop(l, 1)

	var output interface{} //make(map[string]interface{})

	if err := json.Unmarshal([]byte(payload), &output); err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	var recurseOnArray func([]interface{}, int)

	var recurseOnMap func(map[string]interface{}, int)

	forwardOnType := func(val interface{}, depth int) {

		switch val.(type) {
		case nil:
			lua.PushNil(l)

		case bool:
			lua.PushBoolean(l, val.(bool))

		case string:
			lua.PushString(l, val.(string))

		case uint8, uint16, uint32, uint64, uint:
			lua.PushUnsigned(l, val.(uint))

		case int8, int16, int32, int64, int:
			lua.PushInteger(l, val.(int))

		case float32, float64:
			lua.PushNumber(l, val.(float64))

		case []interface{}:
			a := val.([]interface{})
			lua.CreateTable(l, len(a), 0)
			recurseOnArray(a, depth+1)

		case map[string]interface{}:
			m := val.(map[string]interface{})
			lua.CreateTable(l, 0, len(m))
			recurseOnMap(m, depth+1)

		default:
			lua.Errorf(l, fmt.Sprintf("unmarshall: payload contains unsupported type: %T", val))
			panic("unreachable")
		}
	}

	recurseOnMap = func(input map[string]interface{}, depth int) {
		// -1 is a table
		for key, val := range input {
			lua.PushString(l, key)
			// -1: key, -2: table
			forwardOnType(val, depth)
			// -1: something, -2: key, -3: table
			lua.RawSet(l, -3)
		}
	}

	recurseOnArray = func(input []interface{}, depth int) {
		// -1 is a table
		for i, val := range input {
			forwardOnType(val, depth)
			// -1: something, -2: table
			lua.RawSetInt(l, -2, i+1)
		}
	}

	forwardOnType(output, 0)

	return 1
}
