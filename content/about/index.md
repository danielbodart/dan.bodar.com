---
title: "About"
date: 2007-10-06T05:47:51Z
slug: "about"
categories:
  - Uncategorized
comments:
  - {"author":"Stefan Eicher","email":"stefan.eicher@gmail.com","url":"https://github.com/stefaneicher/","date":"2016-12-29T09:35:01Z","content":"Hi Daniel,\r\n\r\nI can't forget your talk https://www.infoq.com/presentations/Crazy-Fast-Build-Times-or-When-10-Seconds-Starts-to-Make-You-Nervous.\r\nI keeps buggimg me that I have to wait some seconds untl my jetty is started.\r\nDou you have an repo or could you send me some code such that I could see how you got those result of some milliseconds to start the server?\r\n\"Tomcat / Jetty take seconds to start, okay you can run Jetty in embedded with helps a lot but nothing, yes nothing is a fast as the the embedded Java 6 HTTP server, 10ms cold JVM, \u003c1ms warm.\"\r\n\r\nLiebe Grüsse \r\nStefan","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"","date":"2017-03-08T11:40:31Z","content":"Basically you want to run Jetty in embedded mode. I do this in one of my libraries:\n\nhttps://github.com/bodar/utterlyidle/blob/master/src/com/googlecode/utterlyidle/jetty/eclipse/RestServer.java\n\nThis take about 10-20ms to start on a desktop machine.","parent":19193}
  - {"author":"Omar Barco","email":"obprado@gmail.com","url":"https://codeforfunandmoney.wordpress.com/","date":"2017-12-15T10:54:30Z","content":"Hi Daniel,\r\n\r\nWe are using Yatspec at my company. We are trying to export the Yatspec output on every build and it would be useful to have the html files being saved somewhere else than 'tmp'. I checked the readme on 'https://github.com/bodar/yatspec' and it says that it can be configured. There is a link apparently pointing to the Configuration.md file (https://github.com/bodar/yatspec/blob/master/Configuration.md) but when trying to follow it I get a cute Github star wars 404 message.\r\n\r\nIs there any place where I can find back this 'Configuration.md'?\r\n\r\nKind regards,\r\nOmar","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"","date":"2017-12-15T11:38:51Z","content":"Ahh that file must have got lost in the migration from Google code to GitHub. Have to see if I can find an old copy and fix it.\n\nIn the mean time if you set the system property \"yatspec.output.dir\" to the location you would like the files to be created.","parent":23021}
  - {"author":"Omar Barco","email":"obprado@gmail.com","url":"https://codeforfunandmoney.wordpress.com/","date":"2017-12-15T11:58:25Z","content":"Thanks!! :D","parent":0}
---

This is Daniel Worthington-Bodart's technical blog, where I will be dumping my current thoughts on all things agile, web related or what ever the current fade is.