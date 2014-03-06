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

function tblEquals(tbl1, tbl2)
  local count = 0
  for k, v in pairs(tbl1) do
    count = count+1
    if (tbl2[k] ~= v) then return false end
  end
  for k, v in pairs(tbl2) do
    count = count-1
    if (tbl1[k] ~= v) then return false end
  end
  return count == 0
end

function tblString(tbl)
  local str = "{"
  for k, v in pairs(tbl) do
    str = str..k..":"..v..","
  end
  return str..'}'
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
local strings = require("goluago/strings")

-- trim
equals("trim: extra space before", "loll", strings.trim(" loll"))
equals("trim: extra space after", "loll", strings.trim("loll "))
equals("trim: extra space surrounds", "loll", strings.trim(" loll "))
equals("trim: multiple white space surrounds", "loll", strings.trim(" \t\n \t\tloll \t\t \n \n\t  "))

-- join
local want = {"cat", "dog", "elephant", "walrus"}
local got = strings.split("cat,dog,elephant,walrus", ",")
print("Want:", tblString(want))
print("Got:", tblString(got))
equals("join: comma separated list", want, got)
equals("join: comma separated list with extra comma", {"", "cat", "dog", "elephant", "walrus",""}, strings.split(",cat,dog,elephant,walrus,", ","))
