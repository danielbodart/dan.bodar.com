---
title: "Samsung Series 9 (NP900X3C-A05UK) Ubuntu/Linux setup for power management and keybindings"
date: 2012-12-07T08:35:32Z
slug: "samsung-series-9-np900x3c-a05uk-ubuntulinux-setup-for-power-management-and-keybindings"
categories:
  - Uncategorized
comments:
  - {"author":"yucel karacalar","email":"yucelkaracalar@yahoo.cm","url":"","date":"2012-12-27T22:42:13Z","content":"Hi Dan is it working FN+F11 Silentmode ? I have np900x3d model using 3.5 low latency kernel but i dont know how to enable silent mode above confiuration is enough to use silent mode ? Thanks for answer","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"http://","date":"2013-01-13T18:04:25Z","content":"I don't actually know what silent mode is meant to do as I have never heard the fan actually come on. Do you know what it should do? And which program it might actually run?","parent":0}
  - {"author":"Keith","email":"keith.hunniford@wump.com","url":"http://www.gnoshme.com","date":"2013-01-18T17:09:48Z","content":"Thanks for putting the time into creating the local.rc and sharing it.  Much appreciated.","parent":0}
  - {"author":"Robert Gabriel","email":"ephemeric@gmail.com","url":"","date":"2013-02-21T14:26:45Z","content":"FN+F11 can be silent, normal or overclocked. Either use samsung-laptop or easy-slow-down-manager to get access.\r\n\r\nSee /proc/ or /sys/ for access and/or use samsung-tools package.\r\n\r\nMine worked and the fan comes on and off as needed.\r\n\r\nBeware of samsung-laptop on UEFI systems, could be bricked as articles report.","parent":0}
  - {"author":"Steve Smith","email":"steve@alwaysagileconsulting.com","url":"http://www.alwaysagileconsulting.com","date":"2014-04-10T20:16:41Z","content":"Am I right in saying only /etc/rc.local is a pre-existing file?\r\n\r\nThanks for this","parent":0}
---

UPDATE: [Ubuntu 14.10 settings](http://dan.bodar.com/2014/12/21/updated-power-savings-settings-for-samsung-series-9-np900x3c-a05uk-and-ubuntu-14-10/)

Ubuntu 12.10 works pretty well out of the box (for me it was just power management and keybindings that needed work), see [Ubuntu wiki](https://help.ubuntu.com/community/SamsungSeries9) if you have any other issues . Using [powertop](https://01.org/powertop/) I was able to understand what needed doing. What follows are my settings:

**/etc/rc.local**

```

#!/bin/sh -e



# Temp disable ethernet port

modprobe -r r8169



# Disable wake up on lan if I do use ethernet port

ethtool -s eth2 wol d;



# Temp disable bluetooth

modprobe -r btusb



# Adjust backlight to start much lower

echo 11 > /sys/class/backlight/acpi_video0/brightness



# - NMI Watchdog (turned off)

echo 0 > '/proc/sys/kernel/nmi_watchdog';



# - SATA Active Link Powermanagement

echo 'min_power' > '/sys/class/scsi_host/host0/link_power_management_policy';



# - USB Autosuspend (after 2 secs of inactivity)

for i in `find /sys/bus/usb/devices/*/power/control`; do echo auto > $i; done;

for i in `find /sys/bus/usb/devices/*/power/autosuspend`; do echo 2 > $i; done;



# - Device Power Management

echo auto | tee /sys/bus/i2c/devices/*/power/control > /dev/null;

echo auto | tee /sys/bus/pci/devices/*/power/control > /dev/null;



# - CPU Scaling (on demand scaling governor for all CPU's

for i in `find /sys/devices/system/cpu/*/cpufreq/scaling_governor`; do echo ondemand > $i; done;



exit 0

```

On the keybingings the only keys that didn't work were some of the Fn Keys

**/lib/udev/keymaps/samsung-other**

```

0xCE prog1              # FN+F1 System Settings (NOT WORKING)

0x89 brightnessdown     # Fn+F2

0x88 brightnessup       # Fn+F3

0x82 switchvideomode    # Fn+F4 CRT/LCD (high keycode: "displaytoggle")

0xF7 f22                # Fn+F5 Touchpad on

0xF9 f23                # Fn+F5 Touchpad off

0x97 kbdillumdown	# FN+F9 Keyboard backlight down

0x96 kbdillumup         # FN+F10 Keyboard backlight up

0xB3 silentmode         # FN+F11 Silentmode (NOT WORKING)

0xD5 wlan               # FN+F12 WiFi  (NOT WORKING)

```

**/lib/udev/keymaps/force-release/samsung-other**

```

# list of scancodes (hex or decimal), optional comment

0xCE # FN+F1 System Settings

0x89 # FN+F2 Brightness down

0x88 # FN+F3 Brightness up

0x82 # FN+F4 Switch video mode

0xCE # FN+F1 System Settings

0x89 # FN+F2 Brightness down

0x88 # FN+F3 Brightness up

0x82 # FN+F4 Switch video mode

0xF7 # Fn+F5 Touchpad on

0xF9 # FN+F5 Turn touchpad off

0x97 # FN+F9 Keyboard backlight down

0x96 # FN+F10 Keyboard backlight up

0xB3 # FN+F11 Silentmode

0xD5 # FN+F12 WiFi

```

The volumn keys and trackpad all worked for me so I didn't change them.