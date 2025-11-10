---
title: "How to list Active Directory controllers from Linux command line"
date: 2014-03-25T12:26:39Z
slug: "how-to-list-active-directory-controllers-from-linux-command-line"
categories:
  - memory
  - windows
---

```
nslookup -type=srv _ldap._tcp.dc._msdcs.YOUR.DOMAIN.COM
```