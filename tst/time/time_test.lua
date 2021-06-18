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

local now = 1.257894e+18
local formattedTime = time.format(now, "2006-01-02T15:04:05Z", "")
equals("time format should be equal", "2009-11-10T23:00:00Z", formattedTime)

local oneHrAdded = time.add(now, {hour= 1})
local formattedTimeAhead = time.format(oneHrAdded, "2006-01-02T15:04:05Z", "")
equals("time: added should be equal", 1.2578976e+18, oneHrAdded)
equals("time format should be equal", "2009-11-11T00:00:00Z", formattedTimeAhead)


equals("time: parse parses a time in a sepcified format", 1.257894e+18, time.parse("2006-01-02T15:04:05Z", "2009-11-10T23:00:00Z"))
equals("time: parseISO parses a time string in ISO8601 encoding", 1.624030596e+18, time.parseISO("2021-06-18T11:36:36-04:00"))
