local hmac = require("goluago/crypto/hmac")

--signsha256
equals(
  "compute hmac with sha256",
  "\40\122\59\216\164\252\119\49\169\76\114\32\121\5\83\35\100\77\135\152\189\41\27\249\135\138\188\155\143\212\177\208"
  , hmac.signsha256("message", "secret-key"))
equals(
  "compute hmac with sha256 and empty secret",
  "\235\8\193\245\109\93\222\224\127\123\223\128\70\128\131\218\6\182\76\244\250\198\79\227\169\8\131\223\95\234\202\228"
  , hmac.signsha256("message", ""))
equals(
  "compute hmac with sha256 and empty message",
  "\43\144\206\61\144\91\186\34\107\61\1\135\87\7\27\42\131\151\216\228\45\157\61\187\150\156\150\173\132\85\221\186",
  hmac.signsha256("", "foobar"))
equals(
  "compute hmac with sha256 and multiple messages",
  "O\xCC\x06\x91[C\xD8\xA4\x9A\xFF\x194A\xE9\xE1\x86T\xE6\xA2|,B\x8B\x02\xE8\xFC\xC4\x1C\xCC\"\x99\xF9",
  hmac.signsha256_multi(array({ "foo", "bar" }), "secret"))
equals(
  "compute hmac with sha256 and multiple messages and empty secret",
  "\xD7\xAF\x9A\xC40\x19\xEBt\xB1x{\xC2,\xC8\xE8\x17\x91\x04_H\xA9K3M\xAB\x1AT!<O\xC6\t",
  hmac.signsha256_multi(array({ "foo", "bar" }), ""))

--sha1
equals(
  "compute hmac with sha1",
  "\f\xafd\x9f\xee\xe4\x95=\x87\xbf\x90:\xc1\x17lE\xe0(\xdf\x16",
  hmac.signsha1("message", "secret"))
equals(
  "compute hmac with sha1 and empty secret",
  "\xd5\xd1\xed\x05\x12\x14\x17$v\x16\xcf\xc87\x8f6\n9\xda|\xfa",
  hmac.signsha1("message", ""))
equals(
  "compute hmac with sha1 and empty message",
  "%\xafat\xa0\xfc\xec\xc4\xd3Fh\nr\xb7\xcedK\x9a\x88\xe8",
  hmac.signsha1("", "secret"))
