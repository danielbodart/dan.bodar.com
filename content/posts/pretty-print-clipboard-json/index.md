---
title: "Pretty Print clipboard JSON"
date: 2016-01-06T13:47:04Z
slug: "pretty-print-clipboard-json"
categories:
  - Uncategorized
---

```
alias json='xclip -o | jq -C "." | less -r'
```