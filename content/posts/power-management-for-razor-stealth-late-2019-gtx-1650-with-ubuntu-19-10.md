---
title: "Power Management for Razor Stealth Late 2019 (GTX 1650) with Ubuntu 19.10"
date: 2020-01-01T07:55:45Z
slug: "power-management-for-razor-stealth-late-2019-gtx-1650-with-ubuntu-19-10"
categories:
  - Uncategorized
tags:
  - Ubuntu
  - power
  - Razor
---

<p>First install <a href="https://01.org/powertop">powertop</a> and <a href="https://linrunner.de/en/tlp/docs/tlp-linux-advanced-power-management.html">TLP</a> . We use powertop to estimate power usage when unplugged and TLP to run as a service </p>

<pre class="wp-block-code"><code>sudo apt install powertop tlp</code></pre>

<p>I then did a little tweaking in my <a href="https://www.linuxbabe.com/linux-server/how-to-enable-etcrc-local-with-systemd">/etc/rc.local</a></p>

<pre class="wp-block-code"><code>#!/bin/sh -e

# Temp disable bluetooth
modprobe -r btusb

# Autosuspend USB Razer Keyboard after 5 minutes 
echo '300000' > '/sys/bus/usb/devices/3-8/power/autosuspend_delay_ms'
echo 'auto' > '/sys/bus/usb/devices/3-8/power/control'

# VM writeback timeout
echo '1500' > '/proc/sys/vm/dirty_writeback_centisecs';

exit 0</code></pre>

<p>Lastly switch from NVidia GTX 1650 to the Intel Iris Pro</p>

<pre class="wp-block-code"><code>sudo prime-select intel</code></pre>

<p>After this my power usage dropped to as low as 1.6 W when Razor Turns off  or 3.6W while typing this</p>
