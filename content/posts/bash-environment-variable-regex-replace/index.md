---
title: "Bash Environment Variable Regex Replace"
date: 2015-09-15T14:19:14Z
slug: "bash-environment-variable-regex-replace"
categories:
  - memory
  - linux
---

```

HOST=tcp://server:443/

echo ${HOST//tcp/https} # prints https://server:443/

```