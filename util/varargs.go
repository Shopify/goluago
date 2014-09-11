package util

import "github.com/Shopify/go-lua"

func PullVarargs(l *lua.State, startIndex int) ([]interface{}, error) {
	top := lua.Top(l)
	if top < startIndex {
		return []interface{}{}, nil
	}

	varargs := make([]interface{}, top-startIndex+1)
	for i := startIndex; i <= top; i++ {
		var value interface{}
		var err error
		switch lua.TypeOf(l, i) {
		case lua.TypeNil:
			value = nil
		case lua.TypeBoolean:
			value = lua.ToBoolean(l, i)
		case lua.TypeLightUserData:
			value = nil // not supported by go-lua
		case lua.TypeNumber:
			value = lua.CheckNumber(l, i)
		case lua.TypeString:
			value = lua.CheckString(l, i)
		case lua.TypeTable:
			value, err = PullTable(l, i)
			if err != nil {
				return nil, err
			}
		case lua.TypeFunction:
			value = lua.ToGoFunction(l, i)
		case lua.TypeUserData:
			value = lua.ToUserData(l, i)
		case lua.TypeThread:
			value = lua.ToThread(l, i)
		}
		varargs[i-startIndex] = value
	}
	return varargs, nil
}
