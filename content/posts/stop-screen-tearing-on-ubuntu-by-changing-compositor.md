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
<pre><code>sudo apt-get install compton 
</code></pre>
Disable existing compositor (I used Metacity with Gnome Flashback)
<pre><code>gsettings set org.gnome.metacity compositing-manager false
</code></pre>
Add Compton to "Startup Applications"
<pre><code>compton --backend glx --paint-on-overlay --vsync opengl-swc</code></pre>