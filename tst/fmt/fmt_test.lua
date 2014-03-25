local fmt = require("goluago/fmt")

equals("can sprintf strings", "hello there", fmt.sprintf("hello %s", "there"))

equals("can sprintf floats", "hello 1", fmt.sprintf("hello %.0f", 1))
equals("can sprintf floats with decimals", "hello 2.2", fmt.sprintf("hello %.1f", 2.22))

-- there's no way to know if it's an int or a float
notequals("cant sprintf integers", "hello 1", fmt.sprintf("hello %d", 1))
notequals("cant sprintf integers", "hello 1", fmt.sprintf("hello %d", 1.0))
