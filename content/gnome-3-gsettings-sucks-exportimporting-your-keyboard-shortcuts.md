---
title: "Gnome 3 / GSettings Sucks: Export/Importing your Keyboard shortcuts"
date: 2012-10-23T10:11:03Z
slug: "gnome-3-gsettings-sucks-exportimporting-your-keyboard-shortcuts"
categories:
  - Uncategorized
comments:
  - {"author":"richard","email":"rkwurth@gmail.com","url":"","date":"2013-02-17T19:59:09Z","content":"Thank you for this.","parent":0}
  - {"author":"pabelmont","email":"pabelmont2007@verizon.net","url":"http://123pab.com","date":"2013-02-23T16:58:52Z","content":"Not too helpful, because\r\ngsettings list-recursively org.gnome.desktop.wm.keybindings \u003e keybindings\r\n\r\n\r\nseems to havfe created a text-file representation of ONLY THE (or SOME?) builtin stuff, not (for instance) my own shortcuts.\r\n\r\nIs there a way to get ALL top-level key-shortcuts (i.e., all those not special to an application)?","parent":0}
  - {"author":"Muhammad Nuzaihan","email":"zaihan@unrealasia.net","url":"","date":"2013-12-22T13:05:06Z","content":"Do you have something similar to using super key to create new tabs, windows similar to Mac OS X?\r\n\r\nIt's faster to switch tabs, create new terminal windows and more and (i think i can put up your files in github for a fork)\r\n\r\nThank you for this blog post.","parent":0}
  - {"author":"Jürgen","email":"juergen@hoetzel.info","url":"http://www.hoetzel.info/","date":"2014-01-02T12:56:35Z","content":"gsettings is just a High-Level-API for application settings. It can use different backends for storage.\r\n\r\nMost likely you use the dconf backend. You can export /import these settings with:\r\n\r\ndconf dump /org/gnome/desktop/wm/keybindings/ \u003ekeybindings-backup.dconf\r\n\r\ndconf load /org/gnome/desktop/wm/keybindings/\r\n\u003ckeybindings-backup.dconf","parent":0}
  - {"author":"confus","email":"con-f-use@gmx.net","url":"","date":"2019-11-01T14:16:21Z","content":"Even worse, the bindings are scattered across many \"Schemas\", e.g.:\r\n\r\n- org.gnome.desktop.wm.keybindings\r\n- org.gnome.settings-daemon.plugins.media-keys\r\n- ...","parent":0}
---

Previous to Gnome 3 I had a nice simple flat file for my keyboard settings, that I just unzipped into .gconf and all was good. Now Gnome is copying the very bad idea of a registry from Windows, you have go through the gsettings API or tool.

Locate the keyboard shortcut schema "org.gnome.desktop.wm.keybindings"

Now lets "export" them

```
gsettings list-recursively org.gnome.desktop.wm.keybindings > keybindings
```

Now unfortunately there does not appear to be a way to import that file, so I just hacked the file. Adding "gsettings set" to the front and quoting the values.  
So

```
org.gnome.desktop.wm.keybindings show-desktop ['<Super>d', '<Primary><Super>d', '<Super>d']
```

becomes

```
gsettings set org.gnome.desktop.wm.keybindings show-desktop "['<Super>d', '<Primary><Super>d', '<Super>d']"
```

I've created [some keybindins](/downloads/keybindings.zip) that use the Windows/Super key for window management, plus this don't conflict with IntellJ