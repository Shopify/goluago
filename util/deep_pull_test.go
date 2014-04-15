package util_test

import (
	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/util"
	"reflect"
	"testing"
)

func TestPullStringTable(t *testing.T) {
	l := lua.NewState()

	want := map[string]string{
		"hello":     "lelele",
		"bye":       "byee",
		"oh":        "hai",
		"this_nil?": "nope",
	}

	util.DeepPush(l, want)
	got, err := util.PullStringTable(l, 1)

	if err != nil {
		t.Fatalf("pulling table, %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestPullStringTableFromLua(t *testing.T) {

	want := map[string]string{
		"hello":     "lelele",
		"bye":       "byee",
		"oh":        "hai",
		"this_nil?": "nope",
	}
	luaWant := `
want = {
    hello = "lelele",
    bye =   "byee",
    oh =    "hai",
}
want["this_nil?"]= "nope"
pull_table(want)`

	l := lua.NewState()

	var got map[string]string
	var err error

	lua.Register(l, "pull_table", func(l *lua.State) int {
		got, err = util.PullStringTable(l, 1)
		return 0
	})
	lua.LoadString(l, luaWant)
	lua.Call(l, 0, 0)

	if err != nil {
		t.Fatalf("pulling table, %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestPullStringTableFromLuaMediumTable(t *testing.T) {

	want := map[string]string{
		"billing_is_shipping":         "on",
		"commit":                      "Continue to next step",
		"order[email]":                "test31805818@example.com",
		"billing_address[first_name]": "Testy-31805818",
		"billing_address[address1]":   "123-31805818 Test Street",
		"billing_address[address2]":   "apt-31805818",
		"billing_address[last_name]":  "McTesterson",
		"billing_address[company]":    "TestCorp",
		"billing_address[city]":       "Winnipeg",
		"billing_address[zip]":        "R2H 0C6",
		"billing_address[country]":    "Canada",
		"billing_address[province]":   "Manitoba",
		"billing_address[phone]":      "2042222222",
	}

	luaWant := `
unique_id=31805818

local order_params = {
  billing_is_shipping = "on",
  commit = "Continue to next step",
}

order_params["order[email]"] = "test"..unique_id.."@example.com"
order_params["billing_address[first_name]"] = "Testy-"..unique_id
order_params["billing_address[address1]"] = "123-"..unique_id.." Test Street"
order_params["billing_address[address2]"] = "apt-"..unique_id
order_params["billing_address[last_name]"] = "McTesterson"
order_params["billing_address[company]"] = "TestCorp"
order_params["billing_address[city]"] = "Winnipeg"
order_params["billing_address[zip]"] = "R2H 0C6"
order_params["billing_address[country]"] = "Canada"
order_params["billing_address[province]"] = "Manitoba"
order_params["billing_address[phone]"] = "2042222222"
order_params["shipping_address[first_name]"] = nil
order_params["shipping_address[last_name]"] = nil
order_params["shipping_address[company]"] = nil
order_params["shipping_address[address1]"] = nil
order_params["shipping_address[address2]"] = nil
order_params["shipping_address[city]"] = nil
order_params["shipping_address[zip]"] = nil
order_params["shipping_address[province]"] = nil
order_params["shipping_address[phone]"] = nil

pull_table(order_params)`

	l := lua.NewState()
	var got map[string]string
	var err error

	lua.Register(l, "pull_table", func(l *lua.State) int {
		got, err = util.PullStringTable(l, 1)
		return 0
	})
	lua.LoadString(l, luaWant)
	lua.Call(l, 0, 0)

	if err != nil {
		t.Fatalf("pulling table, %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %#v, got %#v", want, got)
	}
}
