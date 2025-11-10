---
title: "Shortcuts to position screens on ultra wide screens"
date: 2014-02-08T11:49:10Z
slug: "shortcuts-to-position-screens-on-ultra-wide-screens"
categories:
  - Uncategorized
  - ubuntu
  - memory
---

**left** (Ctrl+Super+Left):

```
wmctrl -r :ACTIVE: -e 0,0,0,1280,1032
```

**right** (Ctrl+Super+Right):

```
wmctrl -r :ACTIVE: -e 0,1280,0,1280,1032
```