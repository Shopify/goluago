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
  if want ~= got then
    if not got then got = "<nil>" end

    print("[FAIL!]\t"..name)
    fail("test '"..name.."': failed assertion, want '"..want.."' but got '"..got.."'")
  else
    print("[ok]\t"..name)
  end
end

function notequals(name, dontwant, got)
  if dontwant == got then
    if not dontwant then dontwant = "<nil>" end
    print("[FAIL!]\t"..name)
    fail("test '"..name.."': failed assertion, do not want '"..dontwant.."' but.. got it!")
  else
    print("[ok]\t"..name)
  end
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
require("math") -- can include
notequals("...sanity check: math.pi not equals to 1", 1, math.pi)
equals("...sanity check: math.pi equals itself", math.pi, 3.1415926535897932384626433832795)

-- All the following should work
require("regexp")

-- quotemeta
equals("quotemeta: example", "\\[foo\\]", regexp.quotemeta("[foo]"))

-- match/matchstring
istrue("match: foo in seafood", regexp.match("foo.*", "seafood"))
isfalse("match: bar is not in seafood", regexp.match("bar.*", "seafood"))

local matched, err = pcall(function() regexp.match("a(b", "seafood") end)
isfalse("match: bad regexp syntax - matched is false", matched)
equals("match: bad regexp syntax - got an error message", "error parsing regexp: missing closing ): `a(b`", err)

local matched, err = pcall(function() regexp.match({}, "seafood") end)
isfalse("match: first arg not string - matched is false", matched)
equals("match: first arg not string - got an error message", "match: 1st arg (pattern) must be a string", err)

local matched, err = pcall(function() regexp.match("foo", {}) end)
isfalse("match: second arg not string - matched is false", matched)
equals("match: second arg not string - got an error message", "match: 2nd arg (s) must be a string", err)
