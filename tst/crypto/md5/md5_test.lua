local md5 = require("goluago/crypto/md5")

equals(
  "go-lua-go md5 string",
  "b0804ec967f48520697662a204f5fe72",
  md5.sum("These pretzels are making me thirsty.")
)
