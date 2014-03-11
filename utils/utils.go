package utils

import (
	"fmt"
	"github.com/Shopify/go-lua"
)

// DeepPush will put any basic Go type on the lua stack. If the value
// contains a map or a slice, it will recursively push those values as
// tables on the Lua stack.
//
// Supported types are:
//    Go                     | Lua
//    -----------------------------------------
//    nil                    | nil
//    bool                   | bool
//    string                 | string
//    any int                | number (float64)
//    any float              | number (float64)
//    any complex            | number (real value as float64)
//    []interface{}          | table as array, child `interface{}` recursively resolved
//    map[string]interface{} | table, child `interface{}` recursively resolved
func DeepPush(l *lua.State, v *interface{}) int {

	var recurseOnArray func([]interface{})
	var recurseOnMap func(map[string]interface{})
	var forwardOnType func(val interface{})

	recurseOnArray = func(input []interface{}) {
		// -1 is a table
		for i, val := range input {
			forwardOnType(val)
			// -1: something, -2: table
			lua.RawSetInt(l, -2, i+1)
		}
	}

	recurseOnMap = func(input map[string]interface{}) {
		// -1 is a table
		for key, val := range input {
			forwardOnType(key)
			// -1: key, -2: table
			forwardOnType(val)
			// -1: something, -2: key, -3: table
			lua.RawSet(l, -3)
		}
	}

	forwardOnType = func(val interface{}) {

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

		case complex64:
			lua.PushNumber(l, float64(real(val)))
		case complex128:
			lua.PushNumber(l, real(val))

		case []interface{}:
			lua.CreateTable(l, len(val), 0)
			recurseOnArray(val)

		case map[string]interface{}:
			lua.CreateTable(l, 0, len(val))
			recurseOnMap(val)

		default:
			lua.Errorf(l, fmt.Sprintf("contains unsupported type: %T", val))
			panic("unreachable")
		}
	}

	forwardOnType(*v)

	return 1
}
