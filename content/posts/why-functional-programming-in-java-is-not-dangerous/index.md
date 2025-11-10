---
title: "Why Functional Programming in Java is NOT Dangerous"
date: 2013-01-23T22:13:24Z
slug: "why-functional-programming-in-java-is-not-dangerous"
categories:
  - Uncategorized
comments:
  - {"author":"Wojtek","email":"wojciech.bulaty@credit-suisse.com","url":"http://test-driven-development.com","date":"2013-01-24T08:40:35Z","content":"Flawless as usual! :)","parent":0}
  - {"author":"Joseph","email":"Dryjins@gmail.com","url":"","date":"2014-08-09T02:42:22Z","content":"Indeed.","parent":0}
---

This is a quick post in response to Elliotte Rusty Harold article titled [Why Functional Programming in Java is Dangerous](http://cafe.elharo.com/programming/java-programming/why-functional-programming-in-java-is-dangerous/).

Lets look at the some of the points made:

- Lazy evaluation
- JIT / JavaC can't optimise
- Recursion

The example that Elliotte uses comes from [Bob Martins article](http://pragprog.com/magazines/2013-01/functional-programming-basics) done in [clojure](http://clojure.org/)

```
(take 25 (squares-of (integers)))
```

Lets show the same things written in Java with [TotallyLazy](http://code.google.com/p/totallylazy/) (Disclaimer: I wrote it) :

```
range(1).map(squared).take(25);
```

You could write the same thing with most functional libraries for Java, as pretty much all of them have lazy Lists or lazy Sequences. In fact the clojure example is doing exactly the same thing: integers returns Seq. If you tried to make it return a PersistentList you would have exactly the same OutOfMemoryError exception.

The article goes on to say "the JIT and javac can't optimize functional constructs as aggressively and efficiently as they can in a real functional language." Obviously clojure is a JVM language that produces byte code and so runs with the same JIT as Java.

The final point is that you can't do recursion with Java, well that's just not true. IBM have been doing tail call optimisation with their Java compiler since I believe version 1.3. Clojure [doesn't even support it implicitly](http://clojure.org/functional_programming#toc6), you need to either explicitly call the recur macro or use an additional library.

For tail call optimisation in Java you could try [JCompilo](http://code.google.com/p/jcompilo/) (Disclaimer: I wrote it). Here is an example:

```
@tailrec
public static int gcd(int x, int y) {
    if (y == 0) return x;
    return gcd(y, x % y);
}
```

If you are reading this article and thinking well Java still doesn't have lambda's till Java 8, then you might want to look at Håkan Råberg's [Enumerable.Java](https://github.com/hraberg/enumerable). Here is the previous example using a lambda:

```
range(1).map(λ( n, n * n)).take(25);
```

Now I'd like to make clear that I think clojure is an amazing language and it inspires me every time I work with it. I'd also like to say that a lot of good stuff is coming in future versions of Java which will make this a much nicer experience but until then it's creating a great space for people like myself and Håkan to innovate. [Java 8 retort here](http://blog.agiledeveloper.com/2013/01/functional-programming-in-java-is-quite.html)