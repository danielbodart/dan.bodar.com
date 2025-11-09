---
title: "Ubuntu file limits 18.04"
date: 2019-07-17T08:04:37Z
slug: "ubuntu-file-limits-18-04"
categories:
  - Uncategorized
---

<p>sudo nano /etc/systemd/user.conf</p>

<pre class="wp-block-syntaxhighlighter-code">DefaultLimitNOFILE=65535</pre>

<p>sudo nano /etc/systemd/system.conf</p>

<pre class="wp-block-syntaxhighlighter-code">DefaultLimitNOFILE=65535</pre>

<p>sudo nano /etc/security/limits.conf</p>

<pre class="wp-block-syntaxhighlighter-code">* hard nofile 65535
* soft nofile 65535</pre>

<p><a href="https://superuser.com/questions/1200539/cannot-increase-open-file-limit-past-4096-ubuntu/1200818#1200818">Source</a></p>

<p></p>
