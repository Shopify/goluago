local json = require "goluago/encoding/json"

equals("unmarshal: can decode empty object", {}, json.unmarshal("{}"))
equals("unmarshal: can decode empty array", array({}), json.unmarshal("[]"))
equals("unmarshal: can decode null", nil, json.unmarshal("null"))
equals("unmarshal: can decode array with null", array({ nil }), json.unmarshal("[null]"))
equals("marshal: can encode empty object", "{}", json.marshal({}))
equals("marshal: can encode null", "null", json.marshal(nil))
equals("marshal: can encode floats", "{\"foo\":1.01}", json.marshal({ foo = 1.01 }))

-- Valid case
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
  "key6": [1, null, 2, null],
  "key7": true
}
]=]

local want = {}
want["key1"]=1
want["key2"]="val2"
want["key3"]=array({"arrKey", 2, true, false, {}, array({})})
want["key4"]={}
want["key4"]["subkey1"]=1
want["key4"]["subkey2"]="val2"
want["key4"]["subkey3"]=array({"arrKey", 2, true, false, {}, array({})})
want["key4"]["subkey4"]={}
want["key4"]["subkey4"]["subsubkey1"]=1
want["key4"]["subkey4"]["subsubkey2"]="val2"
want["key4"]["subkey4"]["subsubkey3"]=array({"arrKey", 2, true, false, {}, array({})})
want["key5"]=nil
want["key6"]=array({1, nil, 2, nil})
want["key7"]=true

equals("unmarshal: can decode sample JSON", want, json.unmarshal(payload))

-- Real use case
local payload = [=[{
  "stuff":"name='updates[1234]'",
  "targetURL":"http://127.0.0.1:7070/targeturl",
  "complete":true
}]=]

local want = {stuff="name='updates[1234]'", targetURL="http://127.0.0.1:7070/targeturl", complete=true}
local got = json.unmarshal(payload)
equals("unmarshal: can decode real use case JSON", want, got)
istrue("unmarshal: true values should be true", got.complete)
equals("marshal: can roundtrip real use case JSON", want, json.unmarshal(json.marshal(got)))

-- Test indented output
local input = {foo = 1 }
local expected_output = [[{
  "foo": 1
}]]

local actual_output = json.marshal(input, 2)

equals("marshal: can encode indented JSON", expected_output, actual_output)


-- Error case
local trailingComma = [=[
{
  "i have": "i trailing comma",
}
]=]

local got, gotErr = pcall(function ()
  return json.unmarshal(trailingComma)
end)

istrue("unmarshal: invalid JSON throws an error", gotErr)
isfalse("unmarshal: invalid JSON decodes a nil value", got)
