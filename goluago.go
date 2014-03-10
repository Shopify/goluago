package goluago

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/encoding/json"
	"github.com/Shopify/goluago/regexp"
	"github.com/Shopify/goluago/strings"
)

func Open(l *lua.State) {
	regexp.Open(l)
	strings.Open(l)
	json.Open(l)
}
