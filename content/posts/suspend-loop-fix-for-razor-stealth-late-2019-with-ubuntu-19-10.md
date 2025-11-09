---
title: "Suspend Loop Fix for Razor Stealth Late 2019 with Ubuntu 19.10"
date: 2020-01-01T08:11:11Z
slug: "suspend-loop-fix-for-razor-stealth-late-2019-with-ubuntu-19-10"
categories:
  - Uncategorized
tags:
  - Ubuntu
  - Razor
---

<p>I was experiencing a loop where after suspending and resuming the laptop it would go back into suspend over and over again. To fix this just update <code>/etc/default/grub</code> and add <code>button.lid_init_state=open</code> to <code>GRUB_CMDLINE_LINUX_DEFAULT</code></p>

<pre class="wp-block-code"><code>GRUB_CMDLINE_LINUX_DEFAULT="quiet splash button.lid_init_state=open"</code></pre>

<p>I also had to add the following to stop errors during the update</p>

<pre class="wp-block-code"><code>GRUB_DISABLE_OS_PROBER=true</code></pre>

<p>Then update GRUB</p>

<pre class="wp-block-code"><code>sudo update-grub</code></pre>

<p>Update:</p>

<p>I found that I actually needed to replace gdm3 with lightdm to fix a different issue and this also resolved the suspend problems</p>
