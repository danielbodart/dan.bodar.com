---
title: "No ERB/Ruby? Just use built in envsubst for templates"
date: 2016-02-03T11:31:52Z
slug: "no-erbruby-just-use-built-in-envsubst-for-templates"
categories:
  - Uncategorized
---

template.txt contains:
<pre><code>Hello ${NAME}</code></pre>

Then you could run:

<pre><code>export NAME=Dan
envsubst < template.txt</code></pre>

And the result would be "Hello Dan"