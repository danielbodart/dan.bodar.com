---
title: "Shortcuts to position screens on ultra wide screens"
date: 2014-02-08T11:49:10Z
slug: "shortcuts-to-position-screens-on-ultra-wide-screens"
categories:
  - Uncategorized
  - ubuntu
  - memory
---

<strong>left</strong> (Ctrl+Super+Left): 
<pre>wmctrl -r :ACTIVE: -e 0,0,0,1280,1032</pre>
<strong>right</strong> (Ctrl+Super+Right): 
<pre>wmctrl -r :ACTIVE: -e 0,1280,0,1280,1032</pre>