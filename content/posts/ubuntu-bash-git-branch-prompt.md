---
title: "Ubuntu Bash Git Branch Prompt!"
date: 2018-01-31T09:04:22Z
slug: "ubuntu-bash-git-branch-prompt"
categories:
  - Uncategorized
---

<pre><code># Git branch in prompt.
parse_git_branch() {
    git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/'
}

export PS1="${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\W\[\033[00m\]:\[\033[32m\]\$(parse_git_branch)\[\033[00m\]\$ "
</code></pre>

This is compatible with the default colour prompt in Ubuntu 16.04, put this in you ~/.bashrc.

<a href="https://martinfitzpatrick.name/article/add-git-branch-name-to-terminal-prompt-mac/">Original version</a> (non Ubuntu specific)
