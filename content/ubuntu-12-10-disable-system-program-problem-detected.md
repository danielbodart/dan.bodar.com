---
title: "Ubuntu - Disable System Program Problem Detected"
date: 2013-01-26T20:09:09Z
slug: "ubuntu-12-10-disable-system-program-problem-detected"
categories:
  - ubuntu
---

```

echo 'enabled=0' | sudo tee /etc/default/apport

```