---
title: "Where did I download that file from?"
date: 2014-07-21T15:05:26Z
slug: "where-did-i-download-that-file-from"
categories:
  - Uncategorized
---

Assuming you used Chrome and a modern Linux file system...

```

$ attr -g xdg.origin.url Downloads/google-chrome-stable_current_x86_64.rpm 

Attribute "xdg.origin.url" had a 74 byte value for /home/dan/Downloads/google-chrome-stable_current_x86_64.rpm:

https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm

```