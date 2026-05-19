---
title: "Git Config"
date: 2014-01-30T09:03:09Z
slug: "git-config"
categories:
  - vcs
  - memory
---

My .gitconfig contains

```
[user]
        email = dan@bodar.com
        name = Daniel Bodart
[color]
        ui = true
[alias]
        ci = commit
        co = checkout
        st = status -sb
        nuke = !git checkout -f && git clean -xf -d
[push]
        default = current
[filter "lfs"]
        clean = git-lfs clean -- %f
        smudge = git-lfs smudge -- %f
        process = git-lfs filter-process
        required = true

```
