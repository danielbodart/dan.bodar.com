---
title: "Fix screen tearing on Razor Stealth (Late 2019) with Ubuntu 19.10"
date: 2020-02-10T21:09:37Z
slug: "fix-screen-tearing-on-razor-stealth-late-2019-with-ubuntu-19-10"
categories:
  - Uncategorized
---

<!-- wp:paragraph -->
<p>So I had my usual problem with screen tearing on Ubuntu 19.10 but instead of my usual switching to Metacity + Compton I wanted to try using Gnome Shell. However Gnome shell doesn't support replacing the compositor.</p>
<!-- /wp:paragraph -->

<!-- wp:paragraph -->
<p>First step is to allow the NVidia driver to support kernel mode setting:</p>
<!-- /wp:paragraph -->

<!-- wp:code -->
<pre class="wp-block-code"><code>sudo nano /etc/modprobe.d/zz-nvidia-modeset.conf</code></pre>
<!-- /wp:code -->

<!-- wp:paragraph -->
<p>Adding</p>
<!-- /wp:paragraph -->

<!-- wp:code -->
<pre class="wp-block-code"><code>options nvidia-drm modeset=1</code></pre>
<!-- /wp:code -->

<!-- wp:paragraph -->
<p>And then</p>
<!-- /wp:paragraph -->

<!-- wp:code -->
<pre class="wp-block-code"><code>sudo update-initramfs -u</code></pre>
<!-- /wp:code -->

<!-- wp:paragraph -->
<p>Check with</p>
<!-- /wp:paragraph -->

<!-- wp:code -->
<pre class="wp-block-code"><code>sudo cat /sys/module/nvidia_drm/parameters/modeset</code></pre>
<!-- /wp:code -->

<!-- wp:paragraph -->
<p>Then add a startup application</p>
<!-- /wp:paragraph -->

<!-- wp:code -->
<pre class="wp-block-code"><code>xrandr --output eDP-1-1 --set 'PRIME Synchronization' '1'</code></pre>
<!-- /wp:code -->

<!-- wp:paragraph -->
<p>Finally switch from gdm3 to lightdm.</p>
<!-- /wp:paragraph -->

<!-- wp:code -->
<pre class="wp-block-code"><code>sudo apt install lightdm</code></pre>
<!-- /wp:code -->

<!-- wp:paragraph -->
<p>This fixed the following issues:</p>
<!-- /wp:paragraph -->

<!-- wp:list -->
<ul><li>Vulkan need me to run <code>sudo vulkaninfo</code> before launching a game </li><li>Suspend / Resume were not working</li><li>External displays didn't always work</li><li>Nicer login screen!</li></ul>
<!-- /wp:list -->