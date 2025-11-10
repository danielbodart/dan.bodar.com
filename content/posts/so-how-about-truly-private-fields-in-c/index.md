---
title: "So how about truly private fields in C#?"
date: 2009-05-21T02:58:55Z
slug: "so-how-about-truly-private-fields-in-c"
categories:
  - Uncategorized
comments:
  - {"author":"Jim Arnold","email":"jim@thoughtworks.com","url":"","date":"2009-05-21T06:52:38Z","content":"var p = new Purse(2);\r\np.set(p.get() + 1);\r\n\r\nint money = (int)p.get.Target.GetType().GetField(\"money\").GetValue(p.get.Target);\r\n\r\nAssert.AreEqual(3, money);\r\n\r\n\r\nOr am I missing the point?","parent":0}
---

UPDATE: Jim pointed out that you can access the field via reflecting over the delegate.  (See comment) Damn this is a bit like how java does anonymous access to private fields of the parent class. I wonder if you could use this for some nasty security violations as people tend to think local variables are safe from reflection.

After the crazy !@$%  with JavaScript yesterday I said to Christian, I bet we can do this with C# lambda. So the challenge was set....

```csharp


class Purse
{
    public Func<int> get;
    public Action<int> set;

    public Purse(int money)
    {
        get = () => { return money; };
        set = (newMoney) => { money = newMoney ; };
    }
}

```

And here is the test ...

```csharp


var p = new Purse(2);
p.set(p.get() + 1);
Assert.AreEqual(3, p.get());

```

If you tried to use reflection, as expected there is no field to inspect.