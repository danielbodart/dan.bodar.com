---
title: "AMD ATI Graphics tearing in Chrome on Ubuntu 15.04"
date: 2015-10-14T13:09:39Z
slug: "amd-ati-graphics-tearing-in-chrome-on-ubuntu-15-04"
categories:
  - ubuntu
  - memory
  - linux
---

I had to switch to using proprietor AMD drivers as the open source ones don't support audio via display port, but after switching, Chrome was tearing and just going completely black.

The <a href="chrome://gpu/">chrome GPU</a> page reported no hardware acceleration enabled.

So I went to <a href="chrome://flags/">Chrome Flags</a> page and forced GPU acceleration and rasterisation

<b>Override software rendering list</b> -> Disabled
<b>Enable GPU rasterization</b> -> Force enabled on all layers
<b>Smooth Scrolling Linux</b> -> Enable



