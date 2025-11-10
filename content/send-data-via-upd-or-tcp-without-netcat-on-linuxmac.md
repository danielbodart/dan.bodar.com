---
title: "Send data via UPD or TCP without Netcat on Linux/Mac"
date: 2016-06-09T14:29:39Z
slug: "send-data-via-upd-or-tcp-without-netcat-on-linuxmac"
categories:
  - ubuntu
  - memory
  - linux
comments:
  - {"author":"Martin","email":"martin.hynar@gmail.com","url":"","date":"2016-06-14T07:13:43Z","content":"This is indeed cool and your post ignited my curiosity on this!\r\nThere are some examples here, also on how to listen - http://www.tldp.org/LDP/abs/html/devref1.html\r\nThe cool thing is that you can also use domain name.","parent":0}
---

I only just found out about this...

```
# Send data via TCP to host 127.0.0.1 port 12201
echo "hello" >/dev/tcp/127.0.0.1/12201

# Send data via UDP to host 127.0.0.1 port 12201
echo "hello" >/dev/udp/127.0.0.1/12201
```

Now is there a way to listen via the same /dev/tcp|udp end point?