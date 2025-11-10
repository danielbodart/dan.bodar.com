---
title: "How to connect to CloudFoundry's Loggregator directly via WebSockets"
date: 2015-04-22T10:50:19Z
slug: "how-to-connect-to-cloudfoundrys-loggregator-directly-via-websockets"
categories:
  - cloudfoundry
---

Firstly make sure you can run the standard CloudFoundry cli:

<pre>
cf logs [APP_NAME]
</pre>

Then turn on tracing:

<pre>
export CF_TRACE=true
</pre>

You will see a HTTP request and then it will switch to a secure WebSocket:

<pre>
WEBSOCKET REQUEST: [2015-04-22T11:27:22+01:00]
GET /tail/?app=[APP_GUID] HTTP/1.1
Host: wss://loggregator.[YOUR_CF_DOMAIN]:443
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: [HIDDEN]
Origin: http://localhost
Authorization: [PRIVATE DATA HIDDEN]
</pre>

As you can see the Authorisation header is hidden, but fear not if look in <code>~/.cf/config.json</code> under "AccessToken" thats all you need (<code>"jq .AccessToken ~/.cf/config.json"</code>). Combine this with <a href="https://www.npmjs.com/package/wscat">wscat</a> and you are good to go:

<pre>
wscat -c wss://loggregator.[YOUR_CF_DOMAIN]:443/tail/?app=[APP_GUID] -H Authorization:$(jq .AccessToken ~/.cf/config.json)
</pre>