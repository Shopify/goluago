local hex = require("goluago/encoding/hex")

local original = "test value"
local hex_value = hex.encode(original)
local decoded = hex.decode(hex_value)

equals("hex can encode to a string to hex value", hex_value, "746573742076616c7565")
equals("hex can encode and decode to same value", original, decoded)

local value, err = pcall(function ()
  return hex.decode("Z")
end)

istrue("#decode returns error if the value is not a hex value.", err)
isfalse("#decode returns nil value is the value is not a hex value.", value)
