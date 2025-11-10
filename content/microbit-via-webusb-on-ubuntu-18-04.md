---
title: "Microbit via WebUSB on Ubuntu 18.04"
date: 2019-12-30T21:09:41Z
slug: "microbit-via-webusb-on-ubuntu-18-04"
categories:
  - Uncategorized
tags:
  - Ubuntu
---

Make sure you are running the [latest firmware](https://makecode.microbit.org/device/usb/webusb/troubleshoot) on the microbit

Create a new udev rule

`sudo nano /etc/udev/rules.d/80-microbit.rules`

With the following content

`SUBSYSTEM=="usb", ATTR{idVendor}=="0d28", ATTR{idProduct}=="0204", MODE="0660", GROUP="plugdev"`

Then if you want non-admin users to be able to use it (like your kids) run

`sudo usermod -a -G plugdev non-admin-user`

If you are admin you just need to unplug and replug the microbit (the udev rules run on hot-plug), if you are non-admin you will need to logout and back in to get the updated user groups

Then you can click on the Settings (top right) and click Pair device. After that you won't need to copy the files to the microbit, download will do everything need.