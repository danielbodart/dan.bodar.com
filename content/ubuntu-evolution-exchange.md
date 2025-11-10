---
title: "Ubuntu + Evolution + Exchange"
date: 2015-11-10T10:53:21Z
slug: "ubuntu-evolution-exchange"
categories:
  - Uncategorized
---

By default Ubuntu does not install the EWS (Exchange Web Service) plugin for Evolution
<pre>sudo apt-get install evolution-ews</pre>
Then go into Evolution and choose the following settings:
<ul>
 	<li><b>Server Type:</b> Exchange Web Service</li>
 	<li><b>Username:</b> WINDOWS-DOMAIN\username</li>
 	<li><b>Host URL:</b> https://webmail.company.com/owa/</li>
</ul>
Now click "Fetch URL" and it should ask you for a password and populate the Host and OAB Url correctly.

It will then try and launch Evolution but for me I had to restart for it to correctly work and then it asked me a couple of times for my windows password.