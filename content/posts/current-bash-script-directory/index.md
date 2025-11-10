---
title: "Current Bash script directory"
date: 2016-03-07T15:52:46Z
slug: "current-bash-script-directory"
categories:
  - ubuntu
  - memory
  - linux
---

```
DIR=`dirname $(readlink -f $0)`
```