---
title: "Updated power savings settings for Samsung Series 9 (NP900X3C-A05UK) and Ubuntu 14.10 "
date: 2014-12-21T10:58:02Z
slug: "updated-power-savings-settings-for-samsung-series-9-np900x3c-a05uk-and-ubuntu-14-10"
categories:
  - Uncategorized
comments:
  - {"author":"Samsung Series 9 (NP900X3C-A05UK) Ubuntu/Linux setup for power management and keybindings | Yesterday I was wrong","email":"","url":"http://dan.bodar.com/2012/12/07/samsung-series-9-np900x3c-a05uk-ubuntulinux-setup-for-power-management-and-keybindings/","date":"2014-12-21T10:59:28Z","content":"[...] UPDATE: Ubuntu 14.10 settings [...]","parent":0}
  - {"author":"passing by","email":"a@b.com","url":"","date":"2015-04-11T08:43:06Z","content":"Super useful. Thanks for posting.","parent":0}
  - {"author":"Steve Smith","email":"steve@alwaysagileconsulting.com","url":"http://www.alwaysagileconsulting.com","date":"2015-12-30T15:40:25Z","content":"Thanks Dan as ever. Do you know if these settings work with 15.10?","parent":0}
  - {"author":"Steve Smith","email":"steve@alwaysagileconsulting.com","url":"http://www.alwaysagileconsulting.com","date":"2015-12-30T15:47:37Z","content":"The brightness setting does not work on 15.10, it seems to be overridden afterwards","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"http://","date":"2016-01-08T15:48:25Z","content":"Steve you might need to adjust the sleep value (5) to be higher or move it into a user profile setting as when you login it can reset the brightness value.\r\n\r\nAlso check the correct path under sys:\r\n\u003cpre\u003e\u003ccode\u003e/sys/class/backlight/intel_backlight/brightness\u003c/code\u003e\u003c/pre\u003e\r\nAs I noticed with different versions of Ubuntu and/or models of the laptop the path is slightly different","parent":0}
---

Here is my updated `/etc/rc.local`for Ubuntu 14.10

```

#!/bin/sh -e



# Sleep so all services have started before we change settings

sleep 5



# Set Intel Audio to power save 

echo '1' > '/sys/module/snd_hda_intel/parameters/power_save';



# Temp disable ethernet port

modprobe -r r8169



# Wireless Power Saving for interface wlan0

iw dev wlan0 set power_save on



# VM writeback timeout

echo '1500' > '/proc/sys/vm/dirty_writeback_centisecs';



# Temp disable bluetooth

modprobe -r btusb



# Adjust backlight to start much lower

echo 800 > '/sys/class/backlight/intel_backlight/brightness'



# - NMI Watchdog (turned off)

echo 0 > '/proc/sys/kernel/nmi_watchdog';



# - SATA Active Link Power management

for i in `find /sys/class/scsi_host/*/link_power_management_policy`; do echo 'min_power' > $i; done;



# - USB Autosuspend (after 2 secs of inactivity)

for i in `find /sys/bus/usb/devices/*/power/control`; do echo auto > $i; done;

for i in `find /sys/bus/usb/devices/*/power/autosuspend`; do echo 2 > $i; done;



# - Device Power Management

echo 'auto' | tee /sys/bus/i2c/devices/*/power/control > /dev/null;

echo 'auto' | tee /sys/bus/pci/devices/*/power/control > /dev/null;



# - CPU Scaling (power saving scaling governor for all CPU's

for i in `find /sys/devices/system/cpu/*/cpufreq/scaling_governor`; do echo 'powersave' > $i; done;



exit 0

```