package goluago

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/pkg/encoding/json"
	"github.com/Shopify/goluago/pkg/env"
	"github.com/Shopify/goluago/pkg/fmt"
	"github.com/Shopify/goluago/pkg/net/url"
	"github.com/Shopify/goluago/pkg/regexp"
	"github.com/Shopify/goluago/pkg/strings"
	"github.com/Shopify/goluago/pkg/time"
	"github.com/Shopify/goluago/util"
)

func Open(l *lua.State) {
	regexp.Open(l)
	strings.Open(l)
	json.Open(l)
	time.Open(l)
	fmt.Open(l)
	url.Open(l)
	util.Open(l)
	env.Open(l)
}
