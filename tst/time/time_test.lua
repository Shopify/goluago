local time = require("goluago/time")

local before = time.now()
time.sleep(1 * 1000) -- 1µs = 1000ns
local after = time.now()

istrue("time: after must be after before (!)", after > before)
isfalse("time: after must be after before (!)", after < before)
notequals("time: before must not be same as after", after, before)

local tenMicroSec = 10 * 1000
local start = time.now()
time.sleep(tenMicroSec)
local diff = time.since(start)
istrue("time: time since sleep must be bigger than sleep duration", diff > tenMicroSec)

now := 1.257894e+18
formattedTime := time.format(now, "2008-10-04T23:00:00Z")
isEqual("time: format should be equal",  "2009-11-10T23:00:00Z", formattedTime)

oneHrAdded := time.add(now, {"hour": 1})
formattedOneHrLater := time.format(oneHrAdded, "2008-10-04T23:00:00Z")
isEqual("time: format should be equal", "2009-11-11T00:00:00Z", formattedOneHrLater)