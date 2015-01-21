package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"github.com/Shopify/go-lua"
)

func Open(l *lua.State) {
	hmacOpen := func(l *lua.State) int {
		lua.NewLibrary(l, hmacLibrary)
		return 1
	}
	lua.Require(l, "goluago/crypto/hmac", hmacOpen, false)
	lua.Pop(l, 1)
}

var hmacLibrary = []lua.RegistryFunction{
	{"signsha256", g},
}

func g(l *lua.State) int {
	message := lua.CheckString(l, 1)
	key := lua.CheckString(l, 2)

	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	lua.PushString(l, string(mac.Sum(nil)))
	return 1
}
