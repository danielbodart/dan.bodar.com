---
title: "Chrome WebGL not working with NVIDIA RTX 5070 Ti on Linux"
date: 2025-02-04T14:30:00Z
slug: "chrome-webgl-nvidia-rtx-5070-ti-linux"
categories:
  - linux
  - ubuntu
  - chrome
  - nvidia
---

After upgrading to an NVIDIA RTX 5070 Ti (Blackwell architecture, released January 2025), WebGL stopped working in Chrome. Sites like [webglreport.com](https://webglreport.com) showed "This browser supports WebGL 2, but it is disabled or unavailable."

The system-level OpenGL worked fine (`glxinfo` showed full OpenGL 4.6 support), so the issue was Chrome-specific.

## The Problem

Chrome was running with `--use-gl=disabled`. You can check this with:

```bash
ps aux | grep chrome | grep -oE '\-\-use-gl=[^ ]*'
```

The RTX 5070 Ti and driver 580.x are so new that Chrome's GPU blocklist doesn't recognise them, so it defaults to disabling GL entirely.

## The Fix

Edit your Chrome `.desktop` launcher to add flags that bypass the blocklist and use Vulkan:

```bash
cp /usr/share/applications/google-chrome.desktop ~/.local/share/applications/
sed -i 's|Exec=/usr/bin/google-chrome-stable|Exec=/usr/bin/google-chrome-stable --ignore-gpu-blocklist --use-angle=vulkan|g' ~/.local/share/applications/google-chrome.desktop
```

This copies the launcher to your local applications folder (survives Chrome updates) and adds the required flags.

### Available `--use-angle` options

| Value | Description |
|-------|-------------|
| `vulkan` | Vulkan backend (recommended for modern NVIDIA) |
| `gl` | Native OpenGL |
| `swiftshader` | Software rendering (CPU-based, slow) |
| `default` | Let Chrome decide |

Restart Chrome for the changes to take effect.

## Bonus: Restart GNOME Shell without logging out

If your application icons disappear or the launcher gets confused after editing `.desktop` files:

**Press `Alt+F2`, type `r`, press Enter**

This restarts GNOME Shell in place without closing your applications.
