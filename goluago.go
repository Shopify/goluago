package goluago

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/encoding/json"
	"github.com/Shopify/goluago/fmt"
	"github.com/Shopify/goluago/net/url"
	"github.com/Shopify/goluago/regexp"
	"github.com/Shopify/goluago/strings"
	"github.com/Shopify/goluago/time"
)

func Open(l *lua.State) {
	regexp.Open(l)
	strings.Open(l)
	json.Open(l)
	time.Open(l)
	fmt.Open(l)
	url.Open(l)
}
