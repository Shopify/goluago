local url = require("goluago/net/url")

local u = url.parse("http://bing.com/search?q=dotnet")

equals("has host", "bing.com", u.host)
equals("has scheme", "http", u.scheme)
equals("has string", "http://bing.com/search?q=dotnet", u.string())


u.scheme = "https"
u.host = "google.com"
equals("can change host", "google.com", u.host)
equals("can change scheme", "https", u.scheme)
equals("can change string", "https://google.com/search?q=dotnet", u.string())

equals("can escape", "https%3A%2F%2Fa.com%3Fsearch%3Dbob%2B1%40a.com", url.escape("https://a.com?search=bob+1@a.com"))
equals("can unescape", "https://a.com?search=bob+1@a.com", url.unescape("https%3A%2F%2Fa.com%3Fsearch%3Dbob%2B1%40a.com"))
