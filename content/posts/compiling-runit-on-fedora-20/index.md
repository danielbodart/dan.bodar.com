---
title: "Compiling RunIt on Fedora 20"
date: 2014-04-30T10:40:53Z
slug: "compiling-runit-on-fedora-20"
categories:
  - Uncategorized
comments:
  - {"author":"Joshua Gies","email":"jsgies@gmail.com","url":"","date":"2017-12-03T18:28:29Z","content":"Thanks a bunch! This is the same error I was facing, and your fix worked for me as well.","parent":0}
---

I had dowloaded [RunIt](http://smarden.org/runit/runit-2.1.1.tar.gz). Then ran

```
package/compile
```

It errored with:

```
./compile runit.c
./load runit unix.a byte.a -static
/usr/bin/ld: cannot find -lc
collect2: error: ld returned 1 exit status
make: *** [runit] Error 1
```

To fix I needed to run

```
sudo yum install glibc-static
```