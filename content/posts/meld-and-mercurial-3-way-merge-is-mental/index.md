---
title: "Meld and Mercurial 3 way Merge is Mental"
date: 2012-07-19T08:29:55Z
slug: "meld-and-mercurial-3-way-merge-is-mental"
categories:
  - vcs
comments:
  - {"author":"Omar","email":"blah@blah.com","url":"","date":"2014-09-13T01:45:20Z","content":"THHHHANNNKKK YOOOOUUUU","parent":0}
  - {"author":"XitasoChris","email":"christopher.gross@xitaso.com","url":"","date":"2015-05-11T12:53:18Z","content":"This behavior has now changed with TortioseHg 3.4. You no longer need to make this configuration change as the default configuration for TortoiseHg now includes this.","parent":0}
  - {"author":"Aaron","email":"wright_left@yahoo.com","url":"","date":"2015-10-29T23:54:43Z","content":"Thanks for this. I wasted a lot of time merging stuff to the middle one, and got nothing out of it. /sigh","parent":0}
---

By default [Meld's](http://meldmerge.org/) 3 way merge combined with [Mercurial](http://mercurial.selenic.com/) merges into the left panel. To make it merge into the middle panel add the following to .hgrc file

```

[merge-tools]

meld.executable = meld

meld.args = $local $base $other -o $output

```