---
title: "Error: Invalid or corrupt jarfile"
date: 2014-02-17T16:18:18Z
slug: "error-invalid-or-corrupt-jarfile"
categories:
  - java
  - memory
comments:
  - {"author":"Eric","email":"eric.cooper-6vustr5@yopmail.com","url":"","date":"2014-04-14T16:15:41Z","content":"Actually leading slash is forbidden by .ZIP File Format Specification (http://www.pkware.com/documents/casestudies/APPNOTE.TXT). It's 4.4.17 clause indicate: \r\nThe path stored MUST not contain a drive or\r\ndevice letter, or a leading slash.  All slashes\r\nMUST be forward slashes '/' (...).\r\n\r\nIt's weird if Java API creates such invalid jars.","parent":0}
  - {"author":"Raymond Barlow","email":"rbarlow@raymanoz.com","url":"","date":"2015-02-19T16:08:56Z","content":"Thanks Dan! We had the 65535 jar limit problem and your blog pointed me in the right directions. Ultimately fixed by doing this:\r\n\r\njava -jar app.jar\r\n# becomes\r\njava -cp app.jar app.Main\r\n\r\n(thanks to https://github.com/sbt/sbt/issues/850)","parent":0}
  - {"author":"anonymous","email":"example@example.org","url":"","date":"2015-03-19T19:31:22Z","content":"Thank you! Good to know that it could be because of issues with the Manifest.","parent":0}
  - {"author":"Giacomo","email":"giacomodemartino@gmail.com","url":"","date":"2019-11-11T10:55:39Z","content":"thank you very mych. it worked for me","parent":0}
---

I was creating a Jar via the Java API's and I couldn't get it to run my main class:

```
 $ java -jar foo.jar
Error: Invalid or corrupt jarfile foo.jar
```

Running it via the class path worked fine:

```
 $ java -cp foo.jar Bar
Hello world!
```

So now I knew it was something to do with the manifest file but it wasn't being caused by

- The 65535 file limit ([See Zip64 and Java 7](https://blogs.oracle.com/xuemingshen/entry/zip64_support_for_4g_zipfile) ).
- [The 72 bytes limit per line](http://docs.oracle.com/javase/7/docs/technotes/guides/jar/jar.html#Notes_on_Manifest_and_Signature_Files)
- Missing newline at the end of the Main-Class (Displays "no main manifest attribute, in foo.jar"
- A blank line before Main-Class (Displays "Error: An unexpected error occurred while trying to open file foo.jar"

So after scratching my head for a while I tried comparing a working jar with the failing jar:

```
$ unzip -lv foo.jar 
Archive:  foo.jar
 Length   Method    Size  Cmpr    Date    Time   CRC-32   Name
--------  ------  ------- ---- ---------- ----- --------  ----
      75  Defl:N       75   0% 2014-02-17 16:01 b1eac370  META-INF/MANIFEST.MF
     825  Defl:N      464  44% 2014-02-17 16:01 942f187c  Working.class
```

```
$ unzip -lv foo.jar 
Archive:  foo.jar
 Length   Method    Size  Cmpr    Date    Time   CRC-32   Name
--------  ------  ------- ---- ---------- ----- --------  ----
      75  Defl:N       75   0% 2014-02-17 16:01 b1eac370  /META-INF/MANIFEST.MF
     825  Defl:N      464  44% 2014-02-17 16:01 942f187c  Failing.class
```

So don't prefix the META-INF folder with a slash! Also note it is case sensitive!