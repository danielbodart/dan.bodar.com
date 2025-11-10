---
title: "99 Bottles of Beer in Totallylazy"
date: 2012-01-13T09:18:02Z
slug: "99-bottles-of-beer-in-totallylazy"
categories:
  - Uncategorized
tags:
  - totallylazy functional java beer
---

\[java]import static com.googlecode.totallylazy.Runnables.printLine;  
import static com.googlecode.totallylazy.lambda.Lambdas.n;  
import static com.googlecode.totallylazy.lambda.Lambdas.λ;  
import static com.googlecode.totallylazy.numbers.Numbers.decrement;  
import static com.googlecode.totallylazy.numbers.Numbers.range;

public class BottlesOfBeer {  
// Run with -javaagent:enumerable-java-0.4.0.jar -cp:totallylazy-598.jar:enumerable-java-0.4.0.jar  
public static void main(String\[] args) {  
range(99, 0).map(λ(n, verse(n))).each(printLine("%s of beer on the wall.\\n"));  
}

private static String verse(Number number) {  
if (number.equals(0))  
return "No more bottles of beer on the wall, no more bottles of beer.\\n" +  
"Go to the store and buy some more, 99 bottles";  
return String.format("%s of beer on the wall, %1$s of beer.\\n" +  
"Take one down and pass it around, %s", bottles(number), bottles(decrement(number)));  
}

private static String bottles(Number number) {  
if (number.equals(0)) return "no more bottles";  
if (number.equals(1)) return "1 bottle";  
return number + " bottles";  
}  
}\[/java]