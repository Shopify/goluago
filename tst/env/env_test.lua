local env = require("goluago/env")

equals("undefined variable is not set", "", env.getenv("A_NEW_VARIABLE"))
env.setenv("A_NEW_VARIABLE","FOO")
equals("setenv of a new variable changes value", "FOO", env.getenv("A_NEW_VARIABLE"))
env.setenv("A_NEW_VARIABLE","BAR")
equals("setenv of an existing variable changes value", "BAR", env.getenv("A_NEW_VARIABLE"))
