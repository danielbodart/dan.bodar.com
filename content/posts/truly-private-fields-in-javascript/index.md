---
title: "Truly private fields in JavaScript"
date: 2009-05-21T02:50:55Z
slug: "truly-private-fields-in-javascript"
categories:
  - Uncategorized
tags:
  - javascript
---

So I was chatting with Christian Blunden about JavaScript, and he asked if it was possible to have private fields in JavaScript.

Now the language doesn't have a key word but I knew that you could use function scoping to achieve the same affect as I had seen the same thing done using the [E programming language](http://www.erights.org/#2levels).

So after 5 minutes here is what we came up with:

```javascript

function Purse(money) {
	this.getMoney = function() {
		return money;
	}
	this.setMoney = function(newMoney) {
		money = newMoney;
	}
}
```

This will create a truely private field that can only be accessed via the methods.

You can still mix your private getters and setters with prototype methods. eg:

```javascript

Purse.prototype = {
	add : function( money ) {
		this.setMoney(this.getMoney() + money);
	}
}

var p = new Purse(2);
p.add( 1 );
p.getMoney(); returns 3
```

If you tried to access the money field directly it would be undefined.