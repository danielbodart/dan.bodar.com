---
title: "Wat? Scala"
date: 2013-12-04T22:13:47Z
slug: "wat-scala"
categories:
  - Uncategorized
comments:
  - {"author":"adsf","email":"asdf@asdf.dfg","url":"","date":"2013-12-04T23:09:12Z","content":"Not an issue since the type safety of the language will save you.","parent":0}
  - {"author":"Mateusz","email":"dymczyk@gmail.com","url":"","date":"2013-12-05T04:10:15Z","content":"Not sure what type inferencing bug you are referring to, but for me (2.9.2) this returns a Set[Int]:\r\n\r\nList(1,2,3).toSet\r\nres3: scala.collection.immutable.Set[Int] = Set(1, 2, 3)\r\n\r\nAlso List(1,2,3).toSet() isn't a bug. \"toSet\" is defined as \"def toSet\" since it does not have any side effects. It's not the clearest feature but I don't think it's a bug.","parent":0}
  - {"author":"daaandi","email":"andi@neumann.biz","url":"http://www.neumann.biz","date":"2013-12-05T10:57:31Z","content":"The idiomatic call works just fine:\r\nscala\u003e List(1,2,3).toSet\r\nres2: scala.collection.immutable.Set[Int] = Set(1, 2, 3)\r\n\r\njust don't clutter your code with unnecessary parantheses.\r\n\r\nWhat has happend here?\r\n\r\n1. there is no method \"toSet()\" defined on List just \"toSet\"\r\n2. What you did was call\r\nList(1,2,3).toSeq.apply()\r\n3. The apply method of a set returns true if the element you call for exists i the set\r\n4. () can't exist","parent":0}
  - {"author":"Ionuț G. Stan","email":"de.sacrificat@gmail.com","url":"http://igstan.ro","date":"2013-12-05T13:24:12Z","content":"Hideous behaviour, I don't want to make excuses for it. Just wanted to point out that adding the \u003ccode\u003e-Yno-adapted-args\u003c/code\u003e to the scalacOptions property will prevent this from happening.","parent":0}
  - {"author":"Alec Zorab","email":"alec.zorab@gmail.com","url":"","date":"2013-12-05T14:07:17Z","content":"Welcome to Scala version 2.10.3 (Java HotSpot(TM) 64-Bit Server VM, Java 1.7.0_21).\r\nType in expressions to have them evaluated.\r\nType :help for more information.\r\n\r\nscala\u003e List(1,2,3).toSet()\r\n:8: warning: Adapting argument list by inserting (): this is unlikely to be what you want.\r\n        signature: GenSetLike.apply(elem: A): Boolean\r\n  given arguments: \r\n after adaptation: GenSetLike((): Unit)\r\n              List(1,2,3).toSet()\r\n                               ^\r\nres0: Boolean = false\r\n\r\n\r\nIt's not like it doesn't warn you that bad things are happening.","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"http://","date":"2013-12-05T14:15:07Z","content":"-Yno-adapted-args\r\n\r\nThanks that's really good to know, hopefully they will remove this behaviour completely real soon.","parent":0}
  - {"author":"Simon","email":"smn@oxnrtr.de","url":"","date":"2013-12-05T14:36:47Z","content":"Hi Dan,\r\n\r\nnice post! It's an issue with auto-tupling, where the compiler tries to be too helpful.\r\n\r\nThe toSet method doesn't take any parameters, so the compiler assumes that you must have meant to call the apply method of Set, and that method returns a Boolean.\r\n\r\nSo what the compiler understands looks more like\r\n  List(1,2,3).toSet.apply(())\r\n\r\nSome discussions about this on the mailing list: https://groups.google.com/d/topic/scala-internals/4RMEZGObPm4/discussion https://groups.google.com/d/topic/scala-debate/zwG8o2YzCWs/discussion\r\n\r\nIf everything works out, we can hopefully get rid of it for 2.11. :-)\r\n\r\nBye,\r\n\r\nSimon","parent":0}
  - {"author":"Simon","email":"smn@oxnrtr.de","url":"","date":"2013-12-05T18:17:08Z","content":"Hi Dan,\r\n\r\nautomatic () insertion will be deprecated for 2.11 and most likely removed in 2.12.\r\n\r\nThe corresponding tickets in the bug tracker are:\r\n\r\nhttps://issues.scala-lang.org/browse/SI-8035\r\nhttps://issues.scala-lang.org/browse/SI-8036\r\n\r\nThanks,\r\n\r\nSimon","parent":0}
  - {"author":"George","email":"acetysal@yahoo.com","url":"","date":"2013-12-05T23:22:43Z","content":"x and x() are different signatures in scala, and the convention is that you use x() for things that have side effects. but that's not even massively relevant here, just rtfm.","parent":0}
  - {"author":"Toby","email":"secret@squirel.com","url":"http://baddotrobot.com","date":"2014-05-29T07:29:52Z","content":"The link to Pete's gist is broken, the new link is https://gist.github.com/petekneller/7803974 \r\n\r\nttfn","parent":0}
  - {"author":"Steven Shaw","email":"steven@steshaw.org","url":"http://steshaw.org/","date":"2015-06-19T11:58:56Z","content":"Seems you only need -deprecation to avoid the issue. Even without the -deprecation, you get a warning which is nice.\r\n\r\nThe 1-star blog you point to is a ridiculous rant from someone who used Scala for a couple of weeks and wasn't taken by the syntax use or use of types. Obviously coming from a Ruby or JS type of background.\r\n\r\nHow about yourself, what programming languages are you into at the moment?","parent":0}
---

As I got quoted recently in ["Scala — 1★ Would Not Program Again"](http://overwatering.org/blog/2013/12/scala-1-star-would-not-program-again/) I though I finally should write up a little [Wat](https://www.destroyallsoftware.com/talks/wat) moment we had recently:

So does anyone know "wat" the following Scala code returns? (Value and Type)

```

List(1,2,3).toSet()

```

A Set&lt;Int&gt; containing 1,2,3?

Nope how about I give you a clue, there are 2 bugs in this one line:

1. A type inferencing bug where it chooses Set&lt;Any&gt;
2. A bug where the brackets are used for both calling the Set.apply method and constructing Unit, notice there no space between the "toSet" and "()"

Yup you guessed it, it returns:

```
false
```

Wat? Try it in your repl and for even more fun check the bytecode out.

**UPDATE:**  
Looks like ("-Yno-adapted-args", "Do not adapt an argument list (either by inserting () or creating a tuple) to match the receiver.") is your friend

Pete Kneller has done some [really good analysis](https://gist.github.com/petekneller/7803974) so you can see all the different weird combinations