---
title: "Ubuntu 13.10 - Reset Gnome Panels"
date: 2013-11-16T09:57:39Z
slug: "ubuntu-13-10-reset-gnome-panels"
categories:
  - ubuntu
  - memory
---

```
dconf reset -f /org/gnome/gnome-panel/
killall gnome-panel
```