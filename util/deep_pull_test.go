package util_test

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"

	luatesting "github.com/Shopify/goluago/pkg/testing"

	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago/util"
	"github.com/bradfitz/iter"
)

////////////////////////
// From Go to lua to Go

func TestPullString_Random0(t *testing.T)    { testRandomMap(t, 0) }
func TestPullString_Random10(t *testing.T)   { testRandomMap(t, 10) }
func TestPullString_Random100(t *testing.T)  { testRandomMap(t, 100) }
func TestPullString_Random1000(t *testing.T) { testRandomMap(t, 1000) }

func testRandomMap(t *testing.T, size int) {
	l := lua.NewState()
	want := randomMap(size)

	util.DeepPush(l, want)
	got, err := util.PullStringTable(l, 1)

	if err != nil {
		t.Fatalf("pulling table, %v", err)
	}

	checkMaps(t, want, got)
}

func checkMaps(t *testing.T, want, got map[string]string) {
	for wantKey, wantVal := range want {
		if gotVal, ok := got[wantKey]; !ok {
			t.Errorf("key '%s' missing", wantKey)
		} else if wantVal != gotVal {
			t.Errorf("key '%s', want %v got %v", wantKey, wantVal, gotVal)
		}
	}

	for gotKey := range got {
		if _, ok := want[gotKey]; !ok {
			t.Errorf("key '%s' extra", gotKey)
		}
	}
}

func randomMap(size int) map[string]string {
	out := make(map[string]string, size)
	for _ = range iter.N(size) {
		out[randomString(40)] = randomString(40)
	}
	return out
}

var alphabet = func() []rune {
	buf := bytes.NewBuffer(nil)
	for i := rune(' '); i <= rune('~'); i++ {
		buf.WriteRune(i)
	}
	return []rune(buf.String())
}()

func randomString(size int) string {

	buf := bytes.NewBuffer(make([]byte, 0, size))
	for _ = range iter.N(size) {
		buf.WriteRune(alphabet[rand.Intn(len(alphabet))])
	}
	return buf.String()
}

/////////////////
// From lua to Go

var fromLuaTT = []struct {
	want map[string]string
	code string
}{
	// small test
	{
		want: map[string]string{
			"hello":     "lelele",
			"bye":       "byee",
			"oh":        "hai",
			"this_nil?": "nope"},
		code: `want = {
        hello = "lelele",
        bye =   "byee",
        oh =    "hai",
    }
    want["this_nil?"]= "nope"
    pull_table(want)`,
	},

	// real usage test
	{
		want: map[string]string{
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
			"billing_address[phone]":      "2042222222"},
		code: `unique_id=31805818

    local want = {
      billing_is_shipping = "on",
      commit = "Continue to next step",
    }

    want["order[email]"] = "test"..unique_id.."@example.com"
    want["billing_address[first_name]"] = "Testy-"..unique_id
    want["billing_address[address1]"] = "123-"..unique_id.." Test Street"
    want["billing_address[address2]"] = "apt-"..unique_id
    want["billing_address[last_name]"] = "McTesterson"
    want["billing_address[company]"] = "TestCorp"
    want["billing_address[city]"] = "Winnipeg"
    want["billing_address[zip]"] = "R2H 0C6"
    want["billing_address[country]"] = "Canada"
    want["billing_address[province]"] = "Manitoba"
    want["billing_address[phone]"] = "2042222222"

    pull_table(want)`,
	},
}

func TestPullStringTableFromLua(t *testing.T) {

	for _, tt := range fromLuaTT {
		want, code := tt.want, tt.code
		l := lua.NewState()

		var got map[string]string
		var err error

		l.Register("pull_table", func(l *lua.State) int {
			got, err = util.PullStringTable(l, 1)
			return 0
		})
		lua.LoadString(l, code)
		l.Call(0, 0)

		if err != nil {
			t.Fatalf("pulling table, %v", err)
		}

		checkMaps(t, want, got)
	}

}

var fromLuaRecTT = []struct {
	want map[string]interface{}
	code string
}{
	{
		want: map[string]interface{}{
			"hello": "lelele",
			"oh":    float64(1234),
			"foo": map[string]interface{}{
				"go": "og",
				"ruby": map[string]interface{}{
					"ru": "ur",
					"by": "yb",
				},
				"c": "c",
			},
		},
		code: `
		want = {
		    hello = "lelele",
		    oh = 1234,
		    foo = {
		        go = "og",
		        ruby = {
		            ru = "ur",
		            by = "yb",
		        },
		        c = "c",
		    },
		}
		pullTable(want)`,
	},
}

func TestPullTableFromLua(t *testing.T) {

	for _, tt := range fromLuaRecTT {
		want, code := tt.want, tt.code
		l := lua.NewState()

		var got interface{}
		var err error

		l.Register("pullTable", func(l *lua.State) int {
			got, err = util.PullTable(l, 1)
			return 0
		})
		lua.LoadString(l, code)
		l.Call(0, 0)

		if err != nil {
			t.Fatalf("pulling table, %v", err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("maps are not equal, expected %v, got %v", want, got)
		}
	}

}

func TestPullTableWithArraysFromLua(t *testing.T) {
	require := func(l *lua.State) {
		util.Open(l)
		l.Register("validateTable", func(l *lua.State) int {
			actual, err := util.PullTable(l, -1)
			if err != nil {
				t.Fatalf("failed to pull table: %s", err.Error())
			}

			expected := map[string]interface{}{
				"foo": []interface{}{
					"I",
					"love",
					map[string]interface{}{
						"lua": []interface{}{"LuaJIT", "go-lua"},
						"go":  []interface{}{"6g", "gcc-go"},
					},
				},
			}
			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("expected %v, actual %v", expected, actual)
			}

			return 0
		})
	}
	luatesting.RunLuaTestString(t, require, `
		validateTable({
			foo = array({
			  "I",
				"love",
				{
					lua = array({"LuaJIT", "go-lua"}),
					go = array({"6g", "gcc-go"}),
				},
			})
		})
	`)
}

func TestPullTableFailsWhenNotATable(t *testing.T) {
	l := lua.NewState()

	l.PushString("not a table")
	_, err := util.PullTable(l, l.Top())

	if err == nil {
		t.Fatalf("strings should not be convertible to tables")
	}
}

func TestPullTableFailsGracefullyOnCyclicStructures(t *testing.T) {
	l := lua.NewState()

	l.NewTable()
	l.PushValue(-1)
	l.SetField(-2, "foo")

	_, err := util.PullTable(l, l.Top())

	if err == nil {
		t.Fatalf("cyclic tables should not not be convertible to maps")
	}
}

func TestPullTableFailsGracefullyOnUnconvertableValues(t *testing.T) {
	l := lua.NewState()

	l.NewTable()
	l.PushGoClosure(func(l *lua.State) int { return 0 }, 0)
	l.SetField(-2, "foo")

	_, err := util.PullTable(l, l.Top())

	if err == nil {
		t.Fatalf("should not be able to convert closure")
	}
}

func TestPullStringArrayFromLua(t *testing.T) {
	require := func(l *lua.State) {
		util.Open(l)
		l.Register("validateStringArray", func(l *lua.State) int {
			actual, err := util.PullStringArray(l, -1)
			if err != nil {
				t.Fatalf("failed to pull string array: %s", err.Error())
			}

			expected := []string{ "foo", "bar" }
			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("expected %v, actual %v", expected, actual)
			}

			return 0
		})
	}
	luatesting.RunLuaTestString(t, require, `
		validateStringArray(array({"foo", "bar"}))
	`)
}

func TestPullStringArrayFailsWhenNotATable(t *testing.T) {
	l := lua.NewState()

	l.PushString("not a table")
	_, err := util.PullStringArray(l, l.Top())

	if err == nil {
		t.Fatalf("strings should not be convertible to string arrays")
	}
}

func TestPullStringArrayFromLuaFailsWhenATableButNotArray(t *testing.T) {
	require := func(l *lua.State) {
		util.Open(l)
		l.Register("validateStringArray", func(l *lua.State) int {
			_, err := util.PullStringArray(l, -1)
			if err == nil {
				t.Fatalf("non-array tables should not be convertible to string arrays")
			}

			return 0
		})
	}
	luatesting.RunLuaTestString(t, require, `
		validateStringArray({ foo = "bar" })
	`)
}

func TestPullStringArrayFromLuaFailsWhenANonStringArray(t *testing.T) {
	require := func(l *lua.State) {
		util.Open(l)
		l.Register("validateStringArray", func(l *lua.State) int {
			_, err := util.PullStringArray(l, -1)
			if err == nil {
				t.Fatalf("non-string arrays should not be convertible to string arrays")
			}

			return 0
		})
	}
	luatesting.RunLuaTestString(t, require, `
		validateStringArray(array({ "foo", array({}) }))
	`)
}
