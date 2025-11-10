---
title: "Samsung Series 9 (NP900X3C-A05UK) Ubuntu/Linux: Wake up when I open the lid"
date: 2013-01-13T17:59:38Z
slug: "samsung-series-9-np900x3c-a05uk-ubuntulinux-wake-up-when-i-open-the-lid"
categories:
  - Uncategorized
comments:
  - {"author":"Dan North","email":"dan@dannorth.net","url":"http://dannorth.net","date":"2013-03-16T10:34:08Z","content":"Hi Dan. Have you noticed the lid close state isn't being recognised again since a recent kernel update? Or have I botched something by accident? I'm on exactly the same build (A05UK) as you.","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"http://","date":"2013-03-16T11:07:47Z","content":"I'm not sure if it ever worked, as it works if you open and close the lid with in a few seconds. \nBasically if it doesn't fully suspend it works. \nI've given up at the moment but let me know if you get any where. Maybe 13.04 will fix...","parent":988}
  - {"author":"Wake up on lid open | James n Sheri.comJames n Sheri.com","email":"","url":"http://jamesnsheri.com/wake-up-on-lid-open/","date":"2013-07-08T02:51:22Z","content":"[...] plenty of references on how to prevent a laptop from waking up, but only one on how to set it up here. Basically, it says to find out what the lid event is called using the command:   Code: ls [...]","parent":0}
---

So one of the last things that wasn't perfectly working was coming out of suspend when I opened the lid. First thing you need to so is find the name of the LID switch:

```
ls /proc/acpi/button/lid/
LID0
```

Mine seems to be LID0 or LID1. You can then check it's current state by doing:

```
cat /proc/acpi/button/lid/LID0/state
state:      open
```

Now all you need to do is add this to the wake up list

```
sudo -s
echo LID0 > /proc/acpi/wakeup
```

That's it,