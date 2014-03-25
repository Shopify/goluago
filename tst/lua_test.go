package tst

import (
	"testing"
)

var tests = []struct {
	pkg      string
	filename string
}{
	{"fmt", "fmt/fmt_test.lua"},
	{"json", "encoding/json/json_test.lua"},
	{"regexp", "regexp/regexp_test.lua"},
	{"strings", "strings/strings_test.lua"},
	{"time", "time/time_test.lua"},
	{"url", "net/url/url_test.lua"},
}

func TestAllPackages(t *testing.T) {
	for _, test := range tests {
		t.Logf("Testing package '%s'", test.pkg)
		runLuaTests(t, test.filename)
	}
}
