---
title: "Microbit via WebUSB on Ubuntu 18.04"
date: 2019-12-30T21:09:41Z
slug: "microbit-via-webusb-on-ubuntu-18-04"
categories:
  - Uncategorized
tags:
  - Ubuntu
---

<p>Make sure you are running the <a href="https://makecode.microbit.org/device/usb/webusb/troubleshoot">latest firmware</a> on the microbit</p>

<p>Create a new udev rule</p>

<p><code>sudo nano /etc/udev/rules.d/80-microbit.rules</code></p>

<p>With the following content</p>

<p><code>SUBSYSTEM=="usb", ATTR{idVendor}=="0d28", ATTR{idProduct}=="0204", MODE="0660", GROUP="plugdev"</code></p>

<p>Then if you want non-admin users to be able to use it (like your kids) run</p>

<p><code>sudo usermod -a -G plugdev non-admin-user</code></p>

<p>If you are admin you just need to unplug and replug the microbit (the udev rules run on hot-plug), if you are non-admin you will need to logout and back in to get the updated user groups</p>

<p>Then you can click on the Settings (top right) and click Pair device. After that you won't need to copy the files to the microbit, download will do everything need.</p>
