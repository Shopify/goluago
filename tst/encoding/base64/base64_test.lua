local base64 = require("goluago/encoding/base64")

--encode
equals("an empty string outputs nothing", "", base64.encode(""))
equals("encodes a string to base64", "Zm9vYmFy", base64.encode("foobar"))

--decode
equals("an empty string outputs nothing", "", base64.decode(""))
equals("decodes a base64 encded string", "ZXhhbXBsZS1tZXNzYWdlLWZvb2Jhcg==", base64.encode("example-message-foobar"))
