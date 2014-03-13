package util

import (
	"fmt"
	"github.com/Shopify/go-lua"
	"reflect"
)

// DeepPush will put any basic Go type on the lua stack. If the value
// contains a map or a slice, it will recursively push those values as
// tables on the Lua stack.
//
// Supported types are:
//
//    | Go                       | Lua
//    |-------------------------------------------------------------------------
//    | nil                      | nil
//    | bool                     | bool
//    | string                   | string
//    | any int                  | number (float64)
//    | any float                | number (float64)
//    | any complex              | number (real value as float64)
//    |                          |
//    | map[t]t                  | table, key and val `t` recursively
//    |                          | resolved
//    |                          |
//    | []t                      | table with array properties, with `t`
//    |                          | values recursively resolved
func DeepPush(l *lua.State, v interface{}) int {
	forwardOnType(l, v)
	return 1
}

func forwardOnType(l *lua.State, val interface{}) {

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
		forwardOnType(l, []float32{real(val), imag(val)})
	case complex128:
		forwardOnType(l, []float64{real(val), imag(val)})

	default:
		forwardOnReflect(l, val)
	}
}

func forwardOnReflect(l *lua.State, val interface{}) {

	switch v := reflect.ValueOf(val); v.Kind() {

	case reflect.Array, reflect.Slice:
		recurseOnFuncSlice(l, func(i int) interface{} { return v.Index(i).Interface() }, v.Len())

	case reflect.Map:
		lua.CreateTable(l, 0, v.Len())
		for _, key := range v.MapKeys() {
			mapKey := key.Interface()
			mapVal := v.MapIndex(key).Interface()
			forwardOnType(l, mapKey)
			forwardOnType(l, mapVal)
			lua.RawSet(l, -3)
		}

	default:
		lua.Errorf(l, fmt.Sprintf("contains unsupported type: %T", val))
		panic("unreachable")
	}

}

// the hack of using a func(int)interface{} makes it that it is valid for any
// type of slice
func recurseOnFuncSlice(l *lua.State, input func(int) interface{}, n int) {
	lua.CreateTable(l, n, 0)
	for i := 0; i < n; i++ {
		forwardOnType(l, input(i))
		lua.RawSetInt(l, -2, i+1)
	}
}
