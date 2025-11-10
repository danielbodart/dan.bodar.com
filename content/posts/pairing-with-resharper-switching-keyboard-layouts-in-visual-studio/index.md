---
title: "Pairing with Resharper (Switching Keyboard Layouts in Visual Studio)"
date: 2008-12-02T05:01:53Z
slug: "pairing-with-resharper-switching-keyboard-layouts-in-visual-studio"
categories:
  - pairing
  - resharper
comments:
  - {"author":"keyboard macro | Digg hot tags","email":"","url":"http://diggwow.info/tags/102/200812/keyboard-macro.html","date":"2008-12-03T01:56:41Z","content":"[...] Vote  Pairing with Resharper (Switching Keyboard Layouts in Visual Studio) [...]","parent":0}
---

So at my current client we have a bunch of Devs that are fairly recent converts to Resharper and a bunch of old hats who know IntelliJ or have used Resharper since Version 1. We are pairing but we know different keyboard layouts and don't want to decrease productivity by making one lot relearn the other layout. So with some searching around I found out how to create new layouts and using a simple macro switch between them. 

Please see [attached zip](keyboardlayout.zip)

**References:**

[An Example Visual Studio Keyboard settings file](http://blogs.msdn.com/jim_glass/archive/2005/02/18/376113.aspx)

[Programatically changing the keyboard thread](http://social.msdn.microsoft.com/forums/en-US/vsx/thread/dde425b4-ba36-4a50-a0a7-47a16d2b921d/)

[MSDN Reference for binding keys in Visual Studio](http://msdn.microsoft.com/en-us/library/envdte.command.bindings.aspx)

 

**What does it do?**

Allow you to switch between "Resharper + Visual Studio" and IntelliJ shortcuts

 

**Why would you do that?**

You are pairing with someone who knows the other set of shortcuts

 

**How to install?**

- Close Visual Studio
- Run install.cmd from a drive (mapped if a network share) to copy the files to the default locations
- Open Visual Studio
- Tools - Macros - Load Macros Project... Select the ThoughtWorks folder Select the ThoughtWorks Project
- Tools -  Options - Environment - Keyboard: Select the Resharper or IntelliJ

 

**How to switch?**

- To switch to IntelliJ Layout press Ctrl-Shift-Alt-I
- To switch to Resharper 4 Layout press Ctrl-Shift-Alt-R

**It still does not work...**

- Tools - Options - Environment - Keyboard: Select the IntelliJ
- Tools - Import and Export Settings... - Import selected environment settings - No, just import new settings - Browse - Select Switch.Shortcuts
- Tools - Options - Environment - Keyboard: Select the Resharper
- Tools - Import and Export Settings... - Import selected environment settings - No, just import new settings - Browse - Select Switch.Shortcuts

 

**Can't you create a decent installer?**

I plan to but it looks like I might need to get my code signed by MS!