local sha256 = require("goluago/crypto/md5")

equals(
  "go-lua-go md5 string",
  "b0dec8aeb0ab1bef807bb7fa5bde5792",
  md5.sum(b0dec8aeb0ab1bef807bb7fa5bde5792)
)
