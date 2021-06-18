local sha256 = require("goluago/crypto/sha256")
local hex = require("goluago/encoding/hex")

equals(
  "empty string",
  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
  hex.encode(sha256.digest(""))
)

equals(
  "The quick...",
  "ef537f25c895bfa782526529a9b63d97aa631564d5d789c2b765448c8635fb6c",
  hex.encode(sha256.digest("The quick brown fox jumps over the lazy dog."))
)
