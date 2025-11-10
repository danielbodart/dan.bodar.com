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

First install [powertop](https://01.org/powertop) and [TLP](https://linrunner.de/en/tlp/docs/tlp-linux-advanced-power-management.html) . We use powertop to estimate power usage when unplugged and TLP to run as a service

```
sudo apt install powertop tlp
```

I then did a little tweaking in my [/etc/rc.local](https://www.linuxbabe.com/linux-server/how-to-enable-etcrc-local-with-systemd)

```
#!/bin/sh -e

# Temp disable bluetooth
modprobe -r btusb

# Autosuspend USB Razer Keyboard after 5 minutes 
echo '300000' > '/sys/bus/usb/devices/3-8/power/autosuspend_delay_ms'
echo 'auto' > '/sys/bus/usb/devices/3-8/power/control'

# VM writeback timeout
echo '1500' > '/proc/sys/vm/dirty_writeback_centisecs';

exit 0
```

Lastly switch from NVidia GTX 1650 to the Intel Iris Pro

```
sudo prime-select intel
```

After this my power usage dropped to as low as 1.6 W when Razor Turns off or 3.6W while typing this