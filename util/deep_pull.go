package util

import (
	"fmt"
	"github.com/Shopify/go-lua"
)

func PullStringTable(l *lua.State, idx int) (map[string]string, error) {
	t := lua.TypeOf(l, idx)
	if t != lua.TypeTable {
		return nil, fmt.Errorf("need a table at index %d, got %v", idx, t)
	}

	// Table at idx
	lua.PushNil(l) // Add free slot for the value, +1

	table := make(map[string]string)
	// +1:nil, idx:table
	for lua.Next(l, idx) {
		// +2:val, +1:key, idx:table
		key, ok := lua.ToString(l, idx+1)
		if !ok {
			return nil, fmt.Errorf("key should be a string (%v)", lua.ToValue(l, idx+1))
		}
		val, ok := lua.ToString(l, idx+2)
		if !ok {
			return nil, fmt.Errorf("value for key '%s' should be a string (%v)", key, lua.ToValue(l, idx+2))
		}
		table[key] = val
		lua.Pop(l, 1) // remove val from top, -1
		// +1:key, idx: table
	}
	lua.Pop(l, 1) // remove the left over table, -1
	return table, nil
}
