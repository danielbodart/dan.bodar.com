---
title: "How to disable click on a Label with just CSS"
date: 2017-03-08T11:34:17Z
slug: "how-to-disable-click-on-a-label-with-just-css"
categories:
  - web
  - memory
  - css
comments:
  - {"author":"Daniel Prieto","email":"samudex@gmail.com","url":"http://prieto.com.ve","date":"2018-11-29T14:24:04Z","content":"Thanks! Just what I needed.","parent":0}
---

<pre><code>pointer-events:none;
display:block;</code></pre>
Display block is required if the label contains other block level elements.