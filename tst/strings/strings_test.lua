local strings = require("goluago/strings")

-- trim
equals("trim: extra space before", "loll", strings.trim(" loll"))
equals("trim: extra space after", "loll", strings.trim("loll "))
equals("trim: extra space surrounds", "loll", strings.trim(" loll "))
equals("trim: multiple white space surrounds", "loll", strings.trim(" \t\n \t\tloll \t\t \n \n\t  "))

-- join
local want = array({"cat", "dog", "elephant", "walrus"})
local got = strings.split("cat,dog,elephant,walrus", ",")
equals("join: comma separated list", want, got)
equals("join: comma separated list with extra comma", array({"", "cat", "dog", "elephant", "walrus",""}), strings.split(",cat,dog,elephant,walrus,", ","))

-- replace
equals("replace: 2 times", "oinky oinky oink", strings.replace("oink oink oink", "k", "ky", 2))
equals("replace: all the times", "moo moo moo", strings.replace("oink oink oink", "oink", "moo", -1))
