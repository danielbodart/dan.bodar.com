---
title: "Using a fullscreen editor to update my blog"
date: 2009-08-31T03:36:25Z
slug: "using-a-fullscreen-editor-to-update-my-blog"
categories:
  - Uncategorized
---

<p>So after rediscovering the distraction free editors to write reports, it seemed mad not to use the same environment to update my blog. I really like the fact that the editor only supports plain text and didn't want to reduce the readability of the plain text by having lots of unneeded markup all over the place.</p>
<p>What I needed was a simple tool to convert my semi structured text into html and then I could just upload that to my blog. Ideally the tool would be a command line tool so I could easily automate the two steps.</p>
<p>I looked at a few different tools but in the end chose to use <a href="http://www.sange.fi/~atehwa/cgi-bin/piki.cgi/stx2any">stx2any</a> as it seemed to be a fairly close fit to the type of plain text I normally use, and also produced the cleanest html.</p>
<p>Then I started to look at tools to upload html into my Wordpress blog, I was expecting there to be quite a few but can honestly say I didn't even find one. I find this a bit wierd as blogs where originally just html and before that text (remember finger). So I started to look at Python scripts that interacted with the <a href="http://www.xmlrpc.com/metaWeblogApi">MetaWeblogAPI</a> but as my Python is fairly limited I thought I could probably do it quicker in <a href="http://www.scala-lang.org/">Scala</a> (as thats what I've been mainly working in lately).</p>
<p>And so an hour or two later <a href="http://code.google.com/p/html2blog/">html2blog</a> was born. And here is how I tied it all together:</p>
<pre>
stx2any --link-abbrevs --make-title off -T html $1 |
 tidy -asxhtml -qc -w 0 | java -jar html2blog.jar
</pre>
<p>Currently html2blog is very limited in that it always creates a new draft, so the next step will be to make it update an existing entry. I'd also like to get images working at some point.</p>
