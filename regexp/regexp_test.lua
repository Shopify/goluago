local re = require("goluago/regexp")

-- quotemeta
equals("quotemeta: example", "\\[foo\\]", re.quotemeta("[foo]"))

-- match/matchstring
istrue("match: foo in seafood", re.match("foo.*", "seafood"))
isfalse("match: bar is not in seafood", re.match("bar.*", "seafood"))

local matched, err = pcall(function() re.match("a(b", "seafood") end)
isfalse("match: bad regexp syntax - matched is false", matched)
equals("match: bad regexp syntax - got an error message", "error parsing regexp: missing closing ): `a(b`", err)

local matched, err = pcall(function() re.match({}, "seafood") end)
isfalse("match: first arg not string - matched is false", matched)
equals("match: first arg not string - got an error message", "match: 1st arg (pattern) must be a string", err)

local matched, err = pcall(function() re.match("foo", {}) end)
isfalse("match: second arg not string - matched is false", matched)
equals("match: second arg not string - got an error message", "match: 2nd arg (s) must be a string", err)
