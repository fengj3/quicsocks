package socks

import (
	"fmt"
)

func GenDefaultHttpResponse(remoteAddrStr string) string {
	httpResponse := ""
	htmlStr := GenDefaultHtml(remoteAddrStr)
	contentLength := len(htmlStr)
	httpHeader := GenDefaultHttpHeader(contentLength)
	httpResponse += httpHeader
	httpResponse += htmlStr
	return httpResponse
}

func GenDefaultHtml(remoteAddrStr string) string {
	htmlStr := ""
	htmlStr += fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head><title>What is my IP Address?</title></head>
<body bgcolor="#FFFFFF">
<center><b>This page shows your IPv4 or IPv6 address</b>
<table width=600 border=0 cellspacing=5 cellpadding=0>
<tr><td align=center colspan=3>You are connecting with an <font color="#FF0000">IPv4</font> Address of:</td></tr>
<tr><td align=center colspan=3 bgcolor="D0D0F0"><font face="Arial, Monospace" size=+3>%s</font></td></tr>
<tr><td align=left><a href="http://ip4.me/">IPv4 only Test</a></td>
<td align=center><a href="http://ip6.me/">Normal Test</a></td>
<td align=right><a href="http://ip6only.me/">IPv6 only Test</a></td></tr>
<tr><td colspan=3><br>&nbsp;<br>If the IPv6 only test shows "Server not found" or similar error or search page then you do not have working IPv6 connectivity.
"Normal Test" shows which protocol your browser preferrs when you have both IPv4 and IPv6 connectivity.
<br>&nbsp;<br>You can access this page with any of these easy to remember url's:
<br>&nbsp;<br><a href="http://ip4.me">ip4.me</a> - IPv4 only test
<br><a href="http://ip6.me">ip6.me</a> - IPv6 test with IPv4 fallback
<br><a href="http://ip6only.me">ip6only.me</a> - IPv6 only test
<br><a href="http://whatismyv6.com">whatismyv6.com</a> - IPv6 test with IPv4 fallback
<br>&nbsp;<br><b>For automated queries</b> use /api/ on any of the urls for a simple plain text csv result that will not be affected by future html changes on the main page.
Recommended API urls<br>(Don't forget the trailing slash to avoid unnessary 301 redirects):
<br>&nbsp;<br><a href="http://ip4only.me/api/">ip4only.me/api/</a> - IPv4 only test
<br><a href="http://ip6.me/api/">ip6.me/api/</a> - IPv6 test with IPv4 fallback
<br><a href="http://ip6only.me/api/">ip6only.me/api/</a> - IPv6 only test
<br>&nbsp;<br>Some day far in the future ip4.me may have a AAAA record so it is not recommended for "IPv4 only" automated queries.  Use ip4only.me instead.
</td></tr></table>
<p>
<font size="-2">&copy;2020 Dulles Internet Exchange, LLC.  All rights reserved.</font>
</center>
</body>
</html>`, remoteAddrStr)
	return htmlStr
}

func GenDefaultHttpHeader(contentLength int) string {
	httpHeaderStr := ""
	httpHeaderStr += fmt.Sprintf("HTTP/3 200 OK\nServer: Apache/2.4.46 (FreeBSD) OpenSSL/1.1.1d-freebsd\nExpires: 0\nCache-control: no-cache\nPragma: no-cache\nContent-Length: %d\nContent-Type: text/html\n", contentLength)
	return httpHeaderStr
}