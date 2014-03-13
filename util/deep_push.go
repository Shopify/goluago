package util

import (
	"fmt"
	"github.com/Shopify/go-lua"
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
//    |------------------------------------------------------------------------
//    | map[string]interface{}   | table, child `interface{}` recursively
//    |                          | resolved if childs are of a type
//    |                          | in the first part of this table
//    |------------------------------------------------------------------------
//
// For string, int, float64, interface{}.
//
//    | Go                       | Lua
//    |-------------------------------------------------------------------------
//    | []                       | table with array properties
//    | [][]                     | 2d table with array properties
//    | [][][]                   | 3d table with array properties
//
// ...where interface{} is not resolved past a fourth level.  This means
//
//    var arr [][][]interface{} = {{{"string", 1, 0.65, true}, {}}, {{}, {}}}
//
// ...will give a Lua table such as:
//
//    {{{string, number, number, true}, {}}, {{}, {}}}
//
// ...but this will not work:
//
//    var arr [][][][]interface{} = {{{{"string", 1, 0.65, true}, {}}, {{}, {}}}}
//
// ... because the fourth level is not of a basic Go type in the first table.
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

	case map[string]interface{}:
		lua.CreateTable(l, 0, len(val))
		recurseOnMap(l, val)

	case []int:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][]int:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][][]int:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))

	case []float64:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][]float64:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][][]float64:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))

	case []string:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][]string:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][][]string:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))

	case []interface{}:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][]interface{}:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))
	case [][][]interface{}:
		recurseOnFuncArray(l, func(i int) interface{} { return val[i] }, len(val))

	default:
		lua.Errorf(l, fmt.Sprintf("contains unsupported type: %T", val))
		panic("unreachable")
	}
}

func recurseOnMap(l *lua.State, input map[string]interface{}) {
	// -1 is a table
	for key, val := range input {
		forwardOnType(l, key)
		forwardOnType(l, val)
		// -1: something, -2: key, -3: table
		lua.RawSet(l, -3)
	}
}

// the hack of using a func(int)interface{} makes it that it is valid for any
// combination of []something
func recurseOnFuncArray(l *lua.State, input func(int) interface{}, n int) {
	lua.CreateTable(l, n, 0)
	for i := 0; i < n; i++ {
		forwardOnType(l, input(i))
		lua.RawSetInt(l, -2, i+1)
	}
}
