---
title: "Fix screen tearing on Razor Stealth (Late 2019) with Ubuntu 19.10"
date: 2020-02-10T21:09:37Z
slug: "fix-screen-tearing-on-razor-stealth-late-2019-with-ubuntu-19-10"
categories:
  - Uncategorized
---

So I had my usual problem with screen tearing on Ubuntu 19.10 but instead of my usual switching to Metacity + Compton I wanted to try using Gnome Shell. However Gnome shell doesn't support replacing the compositor.

First step is to allow the NVidia driver to support kernel mode setting:

```
sudo nano /etc/modprobe.d/zz-nvidia-modeset.conf
```

Adding

```
options nvidia-drm modeset=1
```

And then

```
sudo update-initramfs -u
```

Check with

```
sudo cat /sys/module/nvidia_drm/parameters/modeset
```

Then add a startup application

```
xrandr --output eDP-1-1 --set 'PRIME Synchronization' '1'
```

Finally switch from gdm3 to lightdm.

```
sudo apt install lightdm
```

This fixed the following issues:

- Vulkan need me to run `sudo vulkaninfo` before launching a game
- Suspend / Resume were not working
- External displays didn't always work
- Nicer login screen!