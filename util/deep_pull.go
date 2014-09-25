package util

import (
	"errors"
	"fmt"

	"github.com/Shopify/go-lua"
)

func Open(l *lua.State) {
	lua.Register(l, "array", luaArray)
}

func PullStringTable(l *lua.State, idx int) (map[string]string, error) {
	if !lua.IsTable(l, idx) {
		return nil, fmt.Errorf("need a table at index %d, got %s", idx, lua.TypeNameOf(l, idx))
	}

	// Table at idx
	lua.PushNil(l) // Add free slot for the value, +1

	table := make(map[string]string)
	// -1:nil, idx:table
	for lua.Next(l, idx) {
		// -1:val, -2:key, idx:table
		key, ok := lua.ToString(l, -2)
		if !ok {
			return nil, fmt.Errorf("key should be a string (%v)", lua.ToValue(l, -2))
		}
		val, ok := lua.ToString(l, -1)
		if !ok {
			return nil, fmt.Errorf("value for key '%s' should be a string (%v)", key, lua.ToValue(l, -1))
		}
		table[key] = val
		lua.Pop(l, 1) // remove val from top, -1
		// -1:key, idx: table
	}

	return table, nil
}

func PullTable(l *lua.State, idx int) (interface{}, error) {
	if !lua.IsTable(l, idx) {
		return nil, fmt.Errorf("need a table at index %d, got %s", idx, lua.TypeNameOf(l, idx))
	}

	return pullTableRec(l, idx)
}

func pullTableRec(l *lua.State, idx int) (interface{}, error) {
	if !lua.CheckStack(l, 2) {
		return nil, errors.New("pull table, stack exhausted")
	}

	idx = lua.AbsIndex(l, idx)
	if isArray(l, idx) {
		return pullArrayRec(l, idx)
	}

	table := make(map[string]interface{})

	lua.PushNil(l)
	for lua.Next(l, idx) {
		// -1: value, -2: key, ..., idx: table
		key, ok := lua.ToString(l, -2)
		if !ok {
			err := fmt.Errorf("key should be a string (%s)", lua.TypeNameOf(l, -2))
			lua.Pop(l, 2)
			return nil, err
		}

		value, err := toGoValue(l, -1)
		if err != nil {
			lua.Pop(l, 2)
			return nil, err
		}

		table[key] = value

		lua.Pop(l, 1)
	}

	return table, nil
}

const arrayMarkerField = "_is_array"

func luaArray(l *lua.State) int {
	lua.NewTable(l)
	lua.PushBoolean(l, true)
	lua.SetField(l, -2, arrayMarkerField)
	lua.SetMetaTable(l, -2)
	return 1
}

func isArray(l *lua.State, idx int) bool {
	if !lua.IsTable(l, idx) {
		return false
	}

	if !lua.MetaField(l, idx, arrayMarkerField) {
		return false
	}
	defer lua.Pop(l, 1)

	return lua.ToBoolean(l, -1)
}

func pullArrayRec(l *lua.State, idx int) (interface{}, error) {
	table := make([]interface{}, lua.LengthEx(l, idx))

	lua.PushNil(l)
	for lua.Next(l, idx) {
		k, ok := lua.ToInteger(l, -2)
		if !ok {
			lua.Pop(l, 2)
			return nil, fmt.Errorf("pull array: expected numeric index, got '%s'", lua.TypeOf(l, -2))
		}

		v, err := toGoValue(l, -1)
		if err != nil {
			lua.Pop(l, 2)
			return nil, err
		}

		table[k-1] = v
		lua.Pop(l, 1)
	}

	return table, nil
}

func toGoValue(l *lua.State, idx int) (interface{}, error) {
	t := lua.TypeOf(l, idx)
	switch t {
	case lua.TypeString:
		return lua.CheckString(l, idx), nil
	case lua.TypeNumber:
		return lua.CheckInteger(l, idx), nil
	case lua.TypeTable:
		return pullTableRec(l, idx)
	default:
		err := fmt.Errorf("pull table, unsupported type %s", lua.TypeNameOf(l, idx))
		return nil, err
	}
}
