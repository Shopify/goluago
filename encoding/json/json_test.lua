local json = require("goluago/encoding/json")

equals("unmarshal: can decode empty object", {}, json.unmarshal("{}"))
equals("unmarshal: can decode empty array", {}, json.unmarshal("[]"))
equals("unmarshal: can decode null", nil, json.unmarshal("null"))
equals("unmarshal: can decode array with null", {nil}, json.unmarshal("[null]"))

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

-- Error case
local trailingComma = [=[
{
  "i have": "i trailing comma",
}
]=]

local wantErr = "invalid character '}' looking for beginning of object key string"

local got, gotErr = pcall(function ()
  return json.unmarshal(trailingComma)
end)

equals("unmarshal: invalid JSON throws an error", wantErr, gotErr)
isfalse("unmarshal: invalid JSON decodes a nil value", got)
