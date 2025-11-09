---
title: "Webcam on Razer Stealth Late 2019 with Ubuntu 19.10"
date: 2020-01-01T17:37:47Z
slug: "webcam-on-razer-stealth-late-2019-with-ubuntu-19-10"
categories:
  - Uncategorized
tags:
  - Ubuntu
  - Razor
---

<p>Probably the weakest element on the laptop is the Webcam which can output 720p@30fps. By default the low light performance is not that great but you can improve it considerably by trading frame rate for low light performance. </p>

<p>First install the Video 4 Linux utils:</p>

<pre class="wp-block-code"><code>sudo apt install v4l-utils</code></pre>

<p> Now run the following command (the default is 0)</p>

<pre class="wp-block-code"><code>v4l2-ctl --set-ctrl=exposure_auto_priority=1</code></pre>

<p>Now we have reduced the noise level considerably you can play with the sharpness levels 0-7 (the default is 3)</p>

<pre class="wp-block-code"><code>v4l2-ctl --set-ctrl=sharpness=5</code></pre>

<p></p>
