---
title: "Slow TrackPoint on Dell Precision M6500 under Ubuntu 10.10"
date: 2011-03-03T00:12:32Z
slug: "slow-trackpoint-on-dell-precision-m6500-under-ubuntu-10-10"
categories:
  - Uncategorized
comments:
  - {"author":"Derek","email":"derek_kingston@hotmail.com","url":"","date":"2011-04-16T15:53:55Z","content":"Thanks so much for this! The slow trackpoint was driving me crazy. Have you figured out yet how to get Ubuntu to recognize it as a true trackpoint?","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"http://","date":"2011-04-18T02:49:09Z","content":"Not yet, I've also noticed that after some use the trackpoint seems slow down again but then returns to normal if I stop using for a little while.","parent":0}
---

If you have a Dell Precision M6500 with a Synaptic TouchPad + TrackPoint) (rather than an ALPS version) and you are running Ubuntu 10.10 you may find the TrackPoint / TrackStick / Nipple is being detected as a standard PS/2 Generic Mouse and this is causing the movement to be very slow.

You can check this by typing:

<pre>xinput list</pre>

If you see something like this:

<pre>
Virtual core pointer                    	id=2	[master pointer  (3)]
↳ Virtual core XTEST pointer              	id=4	[slave  pointer  (2)]
↳ SynPS/2 Synaptics TouchPad              	id=11	[slave  pointer  (2)]
↳ PS/2 Generic Mouse                      	id=12	[slave  pointer  (2)]
</pre>

As you can see the TrackPoint is not listed and as such you wont see it in any of the configuration screens.

As I haven't worked out how to get it to detect it as an actual TrackPoint the best solution I have is to change the acceleration mode to more closely represent what you would expect:

<pre>xinput set-prop "PS/2 Generic Mouse" "Device Accel Profile" 6</pre>

I then just add this to the startup applications for my profile.

