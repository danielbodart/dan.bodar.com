---
title: "Mercurial settings"
date: 2013-01-25T12:50:56Z
slug: "mercurial-settings"
categories:
  - memory
---

<pre>
[extensions]
hgext.bookmarks =
rebase =
hgext.purge =
color = 

[ui]
username = Daniel Worthington-Bodart <***@*****.com>

[merge-tools]
meld.executable = meld
meld.args = $local $base $other -o $output

[auth]
google.prefix = code.google.com
google.username = **********
google.password = **********
google.schemes = http https

[alias]
pull = pull --rebase
nuke = !$HG revert --all --no-backup ; $HG purge
</pre>