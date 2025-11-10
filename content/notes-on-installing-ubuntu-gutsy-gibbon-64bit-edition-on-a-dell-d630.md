---
title: "Notes on installing Ubuntu Gutsy Gibbon 64bit edition on a Dell D630"
date: 2007-11-08T16:35:48Z
slug: "notes-on-installing-ubuntu-gutsy-gibbon-64bit-edition-on-a-dell-d630"
categories:
  - ubuntu
---

I had to start the install CD in safe VGA mode and have not got the slash to display yet.

To get sound to work you don't need to recompile your kernel or go back to an earlier version just follow method G one this page:

[Gutsy Intel HD Audio Controller](https://wiki.ubuntu.com/Gutsy_Intel_HD_Audio_Controller#head-0e5a1c0b384a3886c7776913e401a039809c84c9)

or

`sudo aptitude install linux-backports-modules-generic`

`sudo gedit /etc/modprobe.d/alsa-base`

In the editor, add the following line at the end of the file:

`options snd-hda-intel model=dell-m42`

Save the file and reboot to get sound working correctly.

If sound is too low, go to Volume Control's Preferences and add "Front" (and any other playback tracks) and make sure they are set to the maximum.