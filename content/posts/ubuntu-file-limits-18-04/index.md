---
title: "Ubuntu file limits 18.04"
date: 2019-07-17T08:04:37Z
slug: "ubuntu-file-limits-18-04"
categories:
  - Uncategorized
---

sudo nano /etc/systemd/user.conf

```
DefaultLimitNOFILE=65535
```

sudo nano /etc/systemd/system.conf

```
DefaultLimitNOFILE=65535
```

sudo nano /etc/security/limits.conf

```
* hard nofile 65535
* soft nofile 65535
```

[Source](https://superuser.com/questions/1200539/cannot-increase-open-file-limit-past-4096-ubuntu/1200818#1200818)