package util

import (
	"errors"
	"fmt"
	"github.com/Shopify/go-lua"
)

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

func PullTable(l *lua.State, idx int) (map[string]interface{}, error) {
	if !lua.IsTable(l, idx) {
		return nil, fmt.Errorf("need a table at index %d, got %s", idx, lua.TypeNameOf(l, idx))
	}

	return pullTableRec(l, idx)
}

func pullTableRec(l *lua.State, idx int) (map[string]interface{}, error) {
	if !lua.CheckStack(l, 2) {
		return nil, errors.New("pull table, stack exhausted")
	}

	idx = lua.AbsIndex(l, idx)
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

		t := lua.TypeOf(l, -1)
		switch t {
		case lua.TypeString:
			table[key] = lua.CheckString(l, -1)
		case lua.TypeNumber:
			table[key] = lua.CheckInteger(l, -1)
		case lua.TypeTable:
			val, err := pullTableRec(l, -1)
			if err != nil {
				lua.Pop(l, 2)
				return nil, err
			}
			table[key] = val
		default:
			err := fmt.Errorf("pull table, unsupported type %s", lua.TypeNameOf(l, -1))
			lua.Pop(l, 2)
			return nil, err
		}

		lua.Pop(l, 1)
	}

	return table, nil
}
