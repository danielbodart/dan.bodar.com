---
title: "How to connect to CloudFoundry's Loggregator directly via WebSockets"
date: 2015-04-22T10:50:19Z
slug: "how-to-connect-to-cloudfoundrys-loggregator-directly-via-websockets"
categories:
  - cloudfoundry
---

Firstly make sure you can run the standard CloudFoundry cli:

```
cf logs [APP_NAME]
```

Then turn on tracing:

```
export CF_TRACE=true
```

You will see a HTTP request and then it will switch to a secure WebSocket:

```
WEBSOCKET REQUEST: [2015-04-22T11:27:22+01:00]
GET /tail/?app=[APP_GUID] HTTP/1.1
Host: wss://loggregator.[YOUR_CF_DOMAIN]:443
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: [HIDDEN]
Origin: http://localhost
Authorization: [PRIVATE DATA HIDDEN]
```

As you can see the Authorisation header is hidden, but fear not if look in `~/.cf/config.json` under "AccessToken" thats all you need (`"jq .AccessToken ~/.cf/config.json"`). Combine this with [wscat](https://www.npmjs.com/package/wscat) and you are good to go:

```
wscat -c wss://loggregator.[YOUR_CF_DOMAIN]:443/tail/?app=[APP_GUID] -H Authorization:$(jq .AccessToken ~/.cf/config.json)
```