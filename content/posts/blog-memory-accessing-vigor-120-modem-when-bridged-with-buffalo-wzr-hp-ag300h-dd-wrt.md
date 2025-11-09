---
title: "Blog Memory: Accessing Vigor 120 Modem when bridged with Buffalo WZR-HP-AG300H (DD-WRT) "
date: 2012-10-06T19:26:49Z
slug: "blog-memory-accessing-vigor-120-modem-when-bridged-with-buffalo-wzr-hp-ag300h-dd-wrt"
categories:
  - Uncategorized
---

I found this article http://ip6.com/projects/?p=363

So I just changed the IP address to match the default that the Vigor modem uses:

<pre>
ifconfig eth1:1 192.168.2.9 netmask 255.0.0.0
iptables -I POSTROUTING -t nat -o eth1 -d 192.0.0.0/8 -j MASQUERADE
</pre>

It's not perfect but it's a start.
