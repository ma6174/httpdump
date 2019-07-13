# httpdump

```
$ go build httpdump.go && ./httpdump -d 'baidu.com:80'
2019/07/13 15:45:23.875308 http dump server running at :9999 and proxy to baidu.com:80
2019/07/13 15:45:25.167046 bvhynenaz534 ==================== begin [[::1]:54079 <-> 220.181.38.148:80]: ====================
2019/07/13 15:45:25.167231 bvhynenaz534 --> GET http://baidu.com/ HTTP/1.1
2019/07/13 15:45:25.167253 bvhynenaz534 --> Host: baidu.com
2019/07/13 15:45:25.167258 bvhynenaz534 --> User-Agent: curl/7.54.0
2019/07/13 15:45:25.167260 bvhynenaz534 --> Accept: */*
2019/07/13 15:45:25.167262 bvhynenaz534 --> Proxy-Connection: Keep-Alive
2019/07/13 15:45:25.167263 bvhynenaz534 -->
2019/07/13 15:45:25.194879 bvhynenaz534 <-- HTTP/1.1 200 OK
2019/07/13 15:45:25.194909 bvhynenaz534 <-- Date: Sat, 13 Jul 2019 07:45:41 GMT
2019/07/13 15:45:25.194913 bvhynenaz534 <-- Server: Apache
2019/07/13 15:45:25.194916 bvhynenaz534 <-- Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
2019/07/13 15:45:25.194919 bvhynenaz534 <-- ETag: "51-47cf7e6ee8400"
2019/07/13 15:45:25.194922 bvhynenaz534 <-- Accept-Ranges: bytes
2019/07/13 15:45:25.194924 bvhynenaz534 <-- Content-Length: 81
2019/07/13 15:45:25.194932 bvhynenaz534 <-- Cache-Control: max-age=86400
2019/07/13 15:45:25.194935 bvhynenaz534 <-- Expires: Sun, 14 Jul 2019 07:45:41 GMT
2019/07/13 15:45:25.194938 bvhynenaz534 <-- Connection: Keep-Alive
2019/07/13 15:45:25.194940 bvhynenaz534 <-- Content-Type: text/html
2019/07/13 15:45:25.194943 bvhynenaz534 <--
2019/07/13 15:45:25.194946 bvhynenaz534 <-- <html>
2019/07/13 15:45:25.194953 bvhynenaz534 <-- <meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
2019/07/13 15:45:25.194956 bvhynenaz534 <-- </html>
2019/07/13 15:45:25.224281 bvhynenaz534 ==================== end   [[::1]:54079 <-> 220.181.38.148:80]: ====================
```

```
$ curl -v baidu.com -x localhost:9999
* Rebuilt URL to: baidu.com/
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9999 (#0)
> GET http://baidu.com/ HTTP/1.1
> Host: baidu.com
> User-Agent: curl/7.54.0
> Accept: */*
> Proxy-Connection: Keep-Alive
>
< HTTP/1.1 200 OK
< Date: Sat, 13 Jul 2019 07:45:41 GMT
< Server: Apache
< Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
< ETag: "51-47cf7e6ee8400"
< Accept-Ranges: bytes
< Content-Length: 81
< Cache-Control: max-age=86400
< Expires: Sun, 14 Jul 2019 07:45:41 GMT
< Connection: Keep-Alive
< Content-Type: text/html
<
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
* Connection #0 to host localhost left intact
```
