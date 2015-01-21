local crypto = require("goluago/crypto/hmac")

--signsha256
equals("compute hmac with sha256", "\40\122\59\216\164\252\119\49\169\76\114\32\121\5\83\35\100\77\135\152\189\41\27\249\135\138\188\155\143\212\177\208", crypto.signsha256("message", "secret-key"))
equals("compute hmac with sha256 and empty secret", "\235\8\193\245\109\93\222\224\127\123\223\128\70\128\131\218\6\182\76\244\250\198\79\227\169\8\131\223\95\234\202\228", crypto.signsha256("message", ""))
equals("compute hmac with sha256 and empty message", "\43\144\206\61\144\91\186\34\107\61\1\135\87\7\27\42\131\151\216\228\45\157\61\187\150\156\150\173\132\85\221\186", crypto.signsha256("", "foobar"))
