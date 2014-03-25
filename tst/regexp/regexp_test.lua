local re = require("goluago/regexp")

-- quotemeta
equals("quotemeta: example", "\\[foo\\]", re.quotemeta("[foo]"))

-- match/matchstring
istrue("match: foo in seafood", re.match("foo.*", "seafood"))
isfalse("match: bar is not in seafood", re.match("bar.*", "seafood"))

local matched, err = pcall(function() re.match("a(b", "seafood") end)
isfalse("match: bad regexp syntax - matched is false", matched)
istrue("match: bad regexp syntax - got an error message", err)

local matched, err = pcall(function() re.match({}, "seafood") end)
isfalse("match: first arg not string - matched is false", matched)
istrue("match: first arg not string - got an error message", err)

local matched, err = pcall(function() re.match("foo", {}) end)
isfalse("match: second arg not string - matched is false", matched)
istrue("match: second arg not string - got an error message", err)

-- re.find
local r = re.compile("fo.?")
equals("find: can find match - present","foo", r.find("seafood"))
equals("find: can find match - absent","", r.find("meat"))

-- re.findSubmatch
local r = re.compile("a(x*)b(y|z)c")
equals("findSubmatch: can find submatch",{"axxxbyc","xxx","y"},r.findSubmatch("-axxxbyc-"))
equals("findSubmatch: can find submatch",{"abzc","","z"},r.findSubmatch("-abzc-"))

local pattern = re.quotemeta('name="updates[').."([0-9]+)"..re.quotemeta(']"')
local r = re.compile(pattern)
equals("findSubmatch: can find submatch",{'name="updates[1234]"',"1234"}, r.findSubmatch('name="updates[1234]"'))

-- re.findAll
local r = re.compile("a.")
equals("findAll: can find all matches", {"ar", "an", "al"}, r.findAll("paranormal", -1))
equals("findAll: can find limited subset matches", {"ar", "an"}, r.findAll("paranormal", 2))
equals("findAll: can find single match", {"aa"}, r.findAll("graal", -1))
equals("findAll: doesn't find absent matches", {}, r.findAll("none", -1))

-- re.findAllSubmatch
local r = re.compile("a(x*)b")

local want = {{"axxb","xx"}}
local got = r.findAllSubmatch("-axxb-", -1)
equals("findAllSubmatch: unique subgroup", want, got)

local want = {{"ab",""},{"axb","x"}}
local got = r.findAllSubmatch("-ab-axb-", -1)
equals("findAllSubmatch: many subgroups", want, got)

local want = {{"ab",""}}
local got = r.findAllSubmatch("-ab-", -1)
equals("findAllSubmatch: unique subgroup w/ empty string", want, got)

local want = {{"axxb","xx"},{"ab",""}}
local got = r.findAllSubmatch("-axxb-ab-", -1)
equals("findAllSubmatch: many subgroup w/ empty string", want, got)
