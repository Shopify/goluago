local aes = require("goluago/crypto/aes")
local hex = require("goluago/encoding/hex")

local original = "my secret message"
local secret_key = hex.decode("6368616e676520746869732070617373776f726420746f206120736563726574")
local ciphertext = aes.encryptCBC(secret_key, original)
local decrypted = aes.decryptCBC(secret_key, ciphertext)

equals("aes can decrypt what was encrypted", original, decrypted)

local value, err = pcall(function ()
  return aes.encryptCBC("a", original)
end)

istrue("#encryptCBC returns error if the secret key is not valid.", err)
isfalse("#encryptCBC returns nil value if the secret key is not valid.", value)

local value, err = pcall(function ()
  return aes.decryptCBC("a", ciphertext)
end)

istrue("#decryptCBC returns error if the secret key is not valid.", err)
isfalse("#decryptCBC returns nil value if the secret key is not valid.", value)

local value, err = pcall(function ()
  return aes.decryptCBC(hex.decode("7368616e676520746869732070617373776f726420746f206120736563726574"), ciphertext)
end)

istrue("#decryptCBC returns error if the secret key is not the correct one.", err)
isfalse("#decryptCBC returns nil value if the secret key not the correct one.", value)
