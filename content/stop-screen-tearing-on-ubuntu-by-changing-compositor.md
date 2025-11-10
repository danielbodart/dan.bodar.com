---
title: "Stop screen tearing on Ubuntu by changing compositor"
date: 2018-02-26T14:55:37Z
slug: "stop-screen-tearing-on-ubuntu-by-changing-compositor"
categories:
  - Uncategorized
  - ubuntu
  - linux
---

Install Compton (https://github.com/chjj/compton/)

```
sudo apt-get install compton 
```

Disable existing compositor (I used Metacity with Gnome Flashback)

```
gsettings set org.gnome.metacity compositing-manager false
```

Add Compton to "Startup Applications"

```
compton --backend glx --paint-on-overlay --vsync opengl-swc
```