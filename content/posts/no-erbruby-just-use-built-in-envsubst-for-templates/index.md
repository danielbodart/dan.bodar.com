---
title: "No ERB/Ruby? Just use built in envsubst for templates"
date: 2016-02-03T11:31:52Z
slug: "no-erbruby-just-use-built-in-envsubst-for-templates"
categories:
  - Uncategorized
---

template.txt contains:

```
Hello ${NAME}
```

Then you could run:

```
export NAME=Dan
envsubst < template.txt
```

And the result would be "Hello Dan"