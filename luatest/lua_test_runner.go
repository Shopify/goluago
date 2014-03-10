package luatest

import (
	"github.com/Shopify/go-lua"
	"testing"
)

func RunLuaTests(t *testing.T, l *lua.State, program string) {

	// Register the test hook
	failHook := func(l *lua.State) int {
		str := lua.CheckString(l, -1)
		lua.Pop(l, 1)
		t.Error(str)
		return 0
	}
	lua.Register(l, "fail", failHook)

	// Load the harness
	if err := lua.LoadString(l, testHarness); err != nil {
		t.Fatalf("loading lua test harness, %v", err)
	}

	if err := lua.ProtectedCall(l, 0, 0, 0); err != nil {
		t.Fatalf("initializing lua test harness, %v", err)
	}

	// Load and exec the test program
	wantTop := lua.Top(l)

	if err := lua.LoadString(l, program); err != nil {
		t.Fatalf("loading lua test script in VM, %v", err)
	}

	if err := lua.ProtectedCall(l, 0, 0, 0); err != nil {
		t.Errorf("executing lua test script, %v", err)
	}
	gotTop := lua.Top(l)

	if wantTop != gotTop {
		t.Errorf("Unbalanced stack!, want %d, got %d", wantTop, gotTop)
	}
}

const testHarness = `
-- Helper
function istrue(name, cond)
  if not cond then
    print("[FAIL!]\t"..name)
    fail("test '"..name.."': failed assertion, should be true")
  else
    print("[ok]\t"..name)
  end
end

function isfalse(name, cond)
    if cond then
    print("[FAIL!]\t"..name)
    fail("test '"..name.."': failed assertion, should be false")
  else
    print("[ok]\t"..name)
  end
end

function equals(name, want, got)

  local eq = false
  if type(want) == "table" then
    eq = tblEquals(want, got)
    want = tblString(want)
    got = tblString(got)
  else
    eq = (want == got)
  end

  if eq then
    print("[ok]\t"..name)
  else
    if not got then got = "<nil>" end
    print("[FAIL!]\t"..name)
    fail("test '"..name.."': failed assertion, want '"..want.."' but got '"..got.."'")
  end
end

function notequals(name, dontwant, got)
  local eq = false
  if type(want) == "table" then
    eq = tblEquals(dontwant, got)
    dontwant = tblString(dontwant)
    got = tblString(got)
  else
    eq = (want == got)
  end
  if eq then
    if not dontwant then dontwant = "<nil>" end
    print("[FAIL!]\t"..name)
    fail("test '"..name.."': failed assertion, do not want '"..dontwant.."' but.. got it!")
  else
    print("[ok]\t"..name)
  end
end


function tblString(tbl)
  local str = "{"
  for k, v in pairs(tbl) do
    if type(v) == "table" then
      str = str..k..":"..tblString(v)..","
    elseif type(v) == "boolean" then
      if v then
        str = str..k..":true,"
      else
        str = str..k..":false,"
      end
    else
      str = str..k..":"..v..","
    end
  end
  return str.."}"
end

function tblEquals(tbl1, tbl2)
  -- lazy but less duplication of code
  return tblString(tbl1) == tblString(tbl2)
end

istrue("...sanity check: true is true", true)
equals("...sanity check: boolean equality", true, 1==1)
notequals("...sanity check: boolean inequality", false, 1==1)
equals("...sanity check: integer equality", 1, 1)
notequals("...sanity check: integer inequality", 1, 2)
equals("...sanity check: string equality", "hello", "hello")
notequals("...sanity check: string inequality", "hello", "bye")
equals("...sanity check: real equality", 1.0/3.0, 1.0/3.0)
notequals("...sanity check: real inequality", 1.0/3.0, 1.0/6.0)
local mathy = require("math") -- can include
notequals("...sanity check: mathy.pi not equals to 1", 1, mathy.pi)
equals("...sanity check: mathy.pi equals itself", mathy.pi, 3.1415926535897932384626433832795)
`
