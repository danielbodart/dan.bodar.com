---
title: "Html Contracts - How semantic html can help your cross functional team"
date: 2009-07-22T10:48:45Z
slug: "html-contracts-how-semantic-html-can-help-your-cross-functional-team"
categories:
  - Uncategorized
  - web
tags:
  - html
  - semantic
---

One of the pain points we see on web projects is the divide between client side and back end development. This pain might show itself in a number of ways:
<ul>
	<li>Small changes in the HTML cause lots of tests to fail</li>
	<li>Small changes to visual layout require large changes to the HTML which then causes the above</li>
	<li>Developers say the work is done but it can't be signed off as it looks terrible or doesn't work in certain browsers</li>
	<li>CSS or QAs want developer to add ID attributes to lots of elements so they can target them more easily</li>
</ul>
Now ideally all your developers should be poly skilled and understand javascript / CSS / HTML just as well as they understand java / C# / ruby but often the reality is not quite so rosy.
So if we are working in a world where we don't have the ideal but still need to get the job done what can we do to reduce the pain?

Well the technique I have used on a number of teams is to come up with a "HTML contract". An example might be as follows:
<ul>
	<li>Everyone must make the HTML as semantic as possible</li>
	<li>IDs are deprecated in favor of class attributes. Just use more specific CSS selectors to target elements</li>
	<li>Developers will liberally add class attributes with semantic names even if they don't need them immediately (QAs and CSS will use them even if you don't)</li>
	<li>CSS / Designers can add class attributes if needed but can not remove class attributes without pairing with a dev/QA.</li>
	<li>Changing the HTML to support visual display (I'm talking about document order, float and clear) is severely frowned upon. If you have to do it consider doing it with javascript instead.</li>
	<li>QAs are to use hand written XPath expressions in tests that match the domain and make extensive use of contains(@class, 'someClassName') and descendant:: rather than IDs or specifying HTML tags</li>
</ul>
Some more general tips that I find help:
<ul>
	<li>Converting tables to divs is really no better apart from download size. Use divs to group things or as the root element for a control, use lists for lists of things, tables for tabular data etc.</li>
	<li>Hacking HTML, CSS, Javascript just because you are fed up with IE6 is not acceptable</li>
	<li>No Cut and Pasting from the web. Or mass import of Javascript / CSS fixes. Don't put it in if you don't understand what it does.</li>
	<li>When you have a CSS issue keep deleting rules until you work out which rule is causing the problem then rebuild the rules up.</li>
</ul>
I post some examples in my next post