local uuid = require("goluago/uuid")

local u1 = uuid.new()
local u2 = uuid.new()
equals("uuid must be a string", type(u1), "string")
notequals("uuids must differ", u1, u2)
