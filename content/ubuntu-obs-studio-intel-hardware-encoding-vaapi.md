---
title: "Ubuntu + OBS Studio + Intel Hardware Encoding (VAAPI)"
date: 2021-02-27T09:21:07Z
slug: "ubuntu-obs-studio-intel-hardware-encoding-vaapi"
categories:
  - Uncategorized
tags:
  - Ubuntu
  - intel
  - vaapi
  - hardware
---

So I had installed the latest version of OBS Studio (26.x) from the official channels but when I went to the output mode it only listed software encoding. In the logs it mentioned FFMPEG-VAAPI but wasn't using it as any recording was using 30%-50% CPU on a low powered laptop.

In Settings -&gt; Output change Outmode to Advanced (from Simple) then on Streaming -&gt; Encoder change that to FFMPEG VAAPI (Recoding should just be set to use Streaming Encoder which is the default)

But then when I tried to record it

```
[FFMPEG VAAPI encoder] Failed to open VAAPI codec: Invalid argument
```

To fix this I then had to set an environment variable on start

```
 LIBVA_DRIVER_NAME=i965 obs
```

To change the shortcut for OBS I did the following

```
cp /usr/share/applications/com.obsproject.Studio.desktop ~/.local/share/applications/

nano ~/.local/share/applications/com.obsproject.Studio.desktop
```

Change the Exec line as follows

```
Exec=env LIBVA_DRIVER_NAME=i965 obs
```

Now recording only uses 5-10% CPU on the same laptop