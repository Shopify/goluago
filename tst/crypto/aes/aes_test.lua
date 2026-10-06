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

-- Fixed IV + ciphertext: with a random IV, wrong-key output occasionally ends
-- in valid padding (unauthenticated CBC can't always detect a wrong key).
local fixed_ciphertext = hex.decode("e68ad6852aa5a7633b2959587f8b902d2e3cee9ae76ccc5d59122edd75d37f6ffb5a7a8e8826b7583853054cccdd8610")

equals("aes can decrypt a known ciphertext", original, aes.decryptCBC(secret_key, fixed_ciphertext))

local value, err = pcall(function ()
  return aes.decryptCBC(hex.decode("7368616e676520746869732070617373776f726420746f206120736563726574"), fixed_ciphertext)
end)

istrue("#decryptCBC returns error if the secret key is not the correct one.", err)
isfalse("#decryptCBC returns nil value if the secret key not the correct one.", value)

local value, err = pcall(function ()
  return aes.decryptCBC(secret_key, "too short")
end)

istrue("#decryptCBC returns error if the ciphertext is truncated.", err)
isfalse("#decryptCBC returns nil value if the ciphertext is truncated.", value)
