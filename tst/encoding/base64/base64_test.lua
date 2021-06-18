local base64 = require("goluago/encoding/base64")

--encode
equals("an empty string outputs nothing", "", base64.encode(""))
equals("encodes a string to base64", "Zm9vYmFy", base64.encode("foobar"))

--decode
equals("an empty string outputs nothing", "", base64.decode(""))
equals("decodes a base64 encded string", "ZXhhbXBsZS1tZXNzYWdlLWZvb2Jhcg==", base64.encode("example-message-foobar"))

local value, err = pcall(function ()
  return base64.decode("invalid base64")
end)

istrue("#decode returns error if the value is not a base64.", err)
isfalse("#decode returns nil value is the value is not a base64.", value)

equals(
  "urlEncode",
  "aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8_dGVzdD10cnVl",
  base64.urlEncode("https://www.google.com/?test=true")
)

equals(
  "urlDecode",
  "https://www.google.com/?test=true",
  base64.urlDecode("aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8_dGVzdD10cnVl")
)

local value, err = pcall(function ()
  return base64.urlDecode("invalid base64")
end)

istrue("#urlDecode returns error if the value is not a base64.", err)
isfalse("#urlDecode returns nil value is the value is not a base64.", value)
