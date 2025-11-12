---
title: "What I have been doing for the last few months"
date: 2007-10-06T06:20:32Z
slug: "3"
categories:
  - web
comments:
  - {"author":"Jim Webber","email":"jim@webber.name","url":"http://jim.webber.name","date":"2007-10-09T06:23:32Z","content":"Hey Dan,\r\n\r\nWhat's going on inside the GET action that requires the use of transactions?\r\n\r\nGET is meant to be safe and idempotent, but transactions are all about (consistent) state changes which seems to be the antithesis of GET.\r\n\r\nSo I'm intrigued :-)\r\n\r\nJim\r\n\r\nJim","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"http://","date":"2007-10-09T11:25:34Z","content":"Good point, this is because Hibernate needs a transaction to operate queries!!!, with straight JDBC you don't need one but if you wish to enforce the idempotency then make sure the Connection is set to read only.","parent":0}
  - {"author":"Sarah Taraporewalla\u0026#8217;s Technical Ramblings \u0026raquo; Picking the right domain to model","email":"","url":"http://sarahtaraporewalla.com/thoughts/?p=6","date":"2008-09-22T12:21:59Z","content":"[...] booking registration. The registration spans many pages (sections), has no session state (see Dan Bodart’s blog under heading No Session State just persistent documents) and has a similar usability concept as [...]","parent":0}
---

So I thought I would start my first post with what I have been working on for the last couple of months. Naturally I have removed all references to my current client.

**What we have built already:**

**RESTful Web Site and Web Components**

So just like REST exposed your web service api/end points on to url and embraced simple message / state transfer we are going a step further and making all our web components (think html widgets) be exposed by an addressable url. Composition of components can then be done done either server side or client side. And is therefore not language specific so one could compose a widget that comprised of other widgets written in ruby, python and C# etc. In fact the widgets don't even need to be on our server or written by us - think mashups on steroids.

**Post/Redirect/Get pattern**  
http://en.wikipedia.org/wiki/Post/Redirect/Get

We have built a simple interface contract that enforces that all POSTs return a redirect so that the Back button will always work. This combined with the transaction boundaries ensures that all GET requests are idempotent.

**Simple transaction boundaries**

Transaction boundaries have been enforced so that developers do not need to worry about them at all in production code. GET requests run in a transaction that will always be rolled back, POST requests will always commit if successful and always roll back if an error is thrown. Some serious work went into taming Hibernate so that it did not auto commit, had no mutable static state and was completely encapsulated.

**Extended SiteMesh**  
http://www.opensymphony.com/sitemesh/

We are using SiteMesh not only for decoration of the site but for the composition of web components and the ability to extract any content from a page so that when combined with AHAH only the smallest payload is returned to the client keeping response time down.

**AHAH (Asynchronous HTML and HTTP)[](http://microformats.org/wiki/rest/ahah)**http://microformats.org/wiki/rest/ahah[](http://microformats.org/wiki/rest/ahah)

This enables much simpler JavaScript to be written (think JavaScript that never needs to replicate the domain model or business rules on the client) and allows for complete reuse of server side logic. This combined with behaviour CSS bindings ( http://www.bennolan.com/behaviour/ ) leads to NO in-line JavaScript nastiness and more semantic html.

**No Session State just persistent documents**

We have NO session state, all state changes to documents are persisted which leads to a number of advantages: Users never loose data they have filled in, marketing can see exactly how far a user got before balling out of a work flow. Users can fill in data in any order they want and only when they are ready do we action their request. Domain objects are only updated once the document has been validated.

**In memory web acceptance testing**

Think super fast builds, no deployment (in fact no need for a web server), pure Java acceptance tests (i.e. refactorable) and the 80/20 rule on what is good enough to give you confidence that the system works.

**Progressive Enhancement and Accessibility**  
http://en.wikipedia.org/wiki/Progressive_Enhancement

All stories are played vanilla html version first and a second story for JavaScript enhancements. This leads to cleaner simpler semantic html and also allows feedback from the first story to be tacked onto the second story.

**NO Logic in the View**  
http://www.stringtemplate.org/

We are using StringTemplate to enforce that NO logic can be written in the view, plus it's super fast, has no "for loops" (think ruby each blocks).