---
title: "Fix screen tearing on Razor Stealth (Late 2019) with Ubuntu 19.10"
date: 2020-02-10T21:09:37Z
slug: "fix-screen-tearing-on-razor-stealth-late-2019-with-ubuntu-19-10"
categories:
  - Uncategorized
---

<p>So I had my usual problem with screen tearing on Ubuntu 19.10 but instead of my usual switching to Metacity + Compton I wanted to try using Gnome Shell. However Gnome shell doesn't support replacing the compositor.</p>

<p>First step is to allow the NVidia driver to support kernel mode setting:</p>

<pre class="wp-block-code"><code>sudo nano /etc/modprobe.d/zz-nvidia-modeset.conf</code></pre>

<p>Adding</p>

<pre class="wp-block-code"><code>options nvidia-drm modeset=1</code></pre>

<p>And then</p>

<pre class="wp-block-code"><code>sudo update-initramfs -u</code></pre>

<p>Check with</p>

<pre class="wp-block-code"><code>sudo cat /sys/module/nvidia_drm/parameters/modeset</code></pre>

<p>Then add a startup application</p>

<pre class="wp-block-code"><code>xrandr --output eDP-1-1 --set 'PRIME Synchronization' '1'</code></pre>

<p>Finally switch from gdm3 to lightdm.</p>

<pre class="wp-block-code"><code>sudo apt install lightdm</code></pre>

<p>This fixed the following issues:</p>

<ul><li>Vulkan need me to run <code>sudo vulkaninfo</code> before launching a game </li><li>Suspend / Resume were not working</li><li>External displays didn't always work</li><li>Nicer login screen!</li></ul>
