---
title: "More Ubuntu 7.10 notes for Dell D630"
date: 2007-12-29T12:03:06Z
slug: "more-ubuntu-710-notes-for-dell-d630"
categories:
  - ubuntu
comments:
  - {"author":"Boss, Manager, or Leader? \u0026raquo; MJK","email":"","url":"http://www.klynstra.com/blog/?p=81","date":"2007-12-30T12:24:48Z","content":"[...] personal blog with the belief that in a ThoughtWorks syndicated stream of posts on Visual Studio, Ubuntu, and the role of software developer as problem solver, my posts on writing, design, photography and [...]","parent":0}
  - {"author":"sp8472","email":"fake@mailinator.com","url":"","date":"2008-01-05T06:16:10Z","content":"For me the iwl3945 module gives similar problems. Especially when going in hibernation, it will always lock up. And unlike the ipw3945 module that sometimes allows restarting, it seems it never be unloaded or restarted...","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"http://","date":"2008-01-05T13:16:23Z","content":"\u003cp\u003eSo I've just checked: hibernate and suspend are working perfectly for me with iwl3945 module so you will have to try to see if it works for you.\u003c/p\u003e\r\n","parent":0}
  - {"author":"sp8472","email":"bodar@lkshome.cjb.net","url":"","date":"2008-01-07T04:15:26Z","content":"Well, for some reason it doesn't work for me, not with the standard power manager suspend of Ubuntu 7.10 and not with uswsusp...","parent":0}
---

So the wifi seemed to work pretty well out of the box but I noticed after prolonged use it would just suddenly freeze and the only way to make it come back was to reboot. (You couldn't even reload the module or restart networking stack).

You can tell if this will be a problem for you by running iwconfig: if you have lots of Invalid misc errors and the signal and noise levels are fixed at -60dBm then you will want to switch from the ipw3945 module to the iwl3945.

To try it out:
<code>
sudo modprobe -r ipw3945
sudo modprobe -r ieee80211
sudo modprobe -r ieee80211_crypt
sudo modprobe -r mac80211
sudo modprobe iwlwifi_mac80211
sudo modprobe iwl3945
</code>

Now if you run iwconfig you should see wlan0 (plus <a href="http://linuxwireless.org/en/developers/Documentation/mac80211#Themasterdevicewmaster0">wmaster0</a> which we'll just ignore). If it's called wlan0_rename then you run
<code>
sudo nano /etc/udev/rules.d/70-persistent-net.rules
</code>
Comment out the line reserving eth1 for ipw3945 module.

To make this all permanent you need to add the iwl3945 and iwlwifi_mac80211 modules to /etc/modules

<code>sudo echo iwlwifi_mac80211 >> /etc/modules
sudo echo iwl3945 >> /etc/modules
</code>

And now stop the ipw3945 and dependencies from loading by blacklisting them
<code>
sudo echo blacklist ipw3945 >> /etc/modprobe.d/blacklist
sudo echo blacklist ieee80211 >> /etc/modprobe.d/blacklist
sudo echo blacklist ieee80211_crypt >> /etc/modprobe.d/blacklist
sudo echo blacklist mac80211 >> /etc/modprobe.d/blacklist</code>