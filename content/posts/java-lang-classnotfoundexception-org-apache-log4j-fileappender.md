---
title: "java.lang.ClassNotFoundException: org.apache.log4j.FileAppender"
date: 2012-02-28T11:34:16Z
slug: "java-lang-classnotfoundexception-org-apache-log4j-fileappender"
categories:
  - java
  - pain
---

Hopefully this will stop someone from wasting a day of their life...
If you pass in the following when you start your Java application:
<pre>java -Dlog4j.configuration=file:///some/path/log4j.properties
</pre>
And that file contains a class that is not on the class path
<pre>log4j.appender.myAppender=biz.minaret.log4j.DatedFileAppender
log4j.rootLogger=error, myAppender
</pre>
Then due to the log4j static initialisers you will not see an error for the class in question but instead:
<pre>Caused by: java.lang.ClassNotFoundException: org.apache.log4j.FileAppender
at java.net.URLClassLoader$1.run(URLClassLoader.java:202)
at java.security.AccessController.doPrivileged(Native Method)
at java.net.URLClassLoader.findClass(URLClassLoader.java:190)
at java.lang.ClassLoader.loadClass(ClassLoader.java:307)
at java.lang.ClassLoader.loadClass(ClassLoader.java:248)
</pre>
Unfortunately for me this was caused by a transitive dependency changing in Maven. Damn you Maven/Log4J