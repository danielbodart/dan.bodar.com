---
title: "Intel Graphics Installer on Ubuntu 14.10"
date: 2015-01-13T09:52:26Z
slug: "intel-graphics-installer-on-ubuntu-14-10"
categories:
  - ubuntu
comments:
  - {"author":"shawn","email":"plus.shawn@gmail.com","url":"","date":"2015-03-01T07:40:09Z","content":"all iz well . but........when i press install its showing a warning notice......that looks like: \r\n\r\nW:Failed to fetch http://ppa.launchpad.net/tualatrix/ppa/ubuntu/dists/utopic/main/binary-amd64/Packages  404  Not Found\r\n, W:Failed to fetch http://ppa.launchpad.net/tualatrix/ppa/ubuntu/dists/utopic/main/binary-i386/Packages  404  Not Found\r\n, W:Failed to fetch http://ppa.launchpad.net/ubuntu-x-swat/x-updates/ubuntu/dists/utopic/main/binary-amd64/Packages  404  Not Found\r\n, W:Failed to fetch http://ppa.launchpad.net/ubuntu-x-swat/x-updates/ubuntu/dists/utopic/main/binary-i386/Packages  404  Not Found\r\n, E:Some index files failed to download. They have been ignored, or old ones used instead.\r\n\r\n\r\nwhat can i do?","parent":0}
  - {"author":"sai","email":"sairsi_2k2@yahoo.com","url":"","date":"2015-03-10T02:38:52Z","content":"Excellent. you made my day..","parent":0}
---

So if you just downloaded the latest [Intel Graphics Installer](https://01.org/linuxgraphics/) and just found that it doesn't support Ubuntu 14.10.

Fear not you can trick it into installing by doing the following:

First backup

```
sudo cp /etc/lsb-release /etc/lsb-release.backup
```

Then edit the file

```
sudo nano /etc/lsb-release
```

And put the following in there

```
DISTRIB_RELEASE=14.04

DISTRIB_CODENAME=trusty
```

This will allow the installer to proceed but you will also want to add the public key so updates work correctly:

```
wget --no-check-certificate https://download.01.org/gfx/RPM-GPG-KEY-ilg -O - | sudo apt-key add -

wget --no-check-certificate https://download.01.org/gfx/RPM-GPG-KEY-ilg-2 -O - | sudo apt-key add - 
```

Original articles:  
[“Distribution not supported” when trying to install Intel Graphics Installer in 14.10](http://askubuntu.com/questions/553581/distribution-not-supported-when-trying-to-install-intel-graphics-installer-in)  
[Intel Linux Graphic Drivers](http://askubuntu.com/questions/339476/intel-linux-graphic-drivers)