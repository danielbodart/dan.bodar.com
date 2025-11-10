---
title: "Ubuntu Theme - Arc or Materia"
date: 2018-02-06T14:59:20Z
slug: "ark-theme"
categories:
  - Uncategorized
  - ubuntu
  - memory
  - linux
---

```
# Install Arc soft fork
sudo add-apt-repository ppa:fossfreedom/arc-gtk-theme-daily
sudo apt-get install arc-theme
gsettings set org.gnome.desktop.interface gtk-theme 'Arc-Dark'

# Materia
sudo apt install materia-gtk-theme
gsettings set org.gnome.desktop.interface gtk-theme 'Materia-dark-compact'

# https://github.com/daniruiz/flat-remix
sudo add-apt-repository ppa:daniruiz/flat-remix
sudo apt-get install flat-remix
gsettings set org.gnome.desktop.interface icon-theme 'Flat-Remix-Blue-Dark'
```