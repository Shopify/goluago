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

-- All the following should work
local json = require("goluago/encoding/json")

equals("unmarshall: can decode empty object", {}, json.unmarshall("{}"))
equals("unmarshall: can decode empty array", {}, json.unmarshall("[]"))
equals("unmarshall: can decode null", nil, json.unmarshall("null"))
equals("unmarshall: can decode array with null", {}, json.unmarshall("[null]"))

local payload = [=[
{
  "key1":1,
  "key2":"val2",
  "key3":["arrKey", 2, true, false, {}, []],
  "key4": {
    "subkey1":1,
    "subkey2":"val2",
    "subkey3":["arrKey", 2, true, false, {}, []],
    "subkey4": {
      "subsubkey1":1,
      "subsubkey2":"val2",
      "subsubkey3":["arrKey", 2, true, false, {}, []]
    }
  },
  "key5": null,
  "key6": [1, null, 2, null]
}
]=]


local want = {}
want["key1"]=1
want["key2"]="val2"
want["key3"]={"arrKey", 2, true, false, {}, {}}
want["key4"]={}
want["key4"]["subkey1"]=1
want["key4"]["subkey2"]="val2"
want["key4"]["subkey3"]={"arrKey", 2, true, false, {}, {}}
want["key4"]["subkey4"]={}
want["key4"]["subkey4"]["subsubkey1"]=1
want["key4"]["subkey4"]["subsubkey2"]="val2"
want["key4"]["subkey4"]["subsubkey3"]={"arrKey", 2, true, false, {}, {}}
want["key5"]=nil
want["key6"]={1, nil, 2, nil}

local got = json.unmarshall(payload)

equals("unmarshall: can decode sample JSON", want, got)

local trailingComma = [=[
{
  "i have": "i trailing comma",
}
]=]

local wantErr = "invalid character '}' looking for beginning of object key string"

local got, gotErr = pcall(function ()
  return json.unmarshall(trailingComma)
end)

equals("unmarshall: invalid JSON throws an error", wantErr, gotErr)
isfalse("unmarshall: invalid JSON decodes a nil value", got)
