---
title: "Excel Integration the old school way: Web Queries (IQY) and Data Queries (DQY)"
date: 2013-02-18T15:08:06Z
slug: "excel-integration-the-old-school-way-web-queries-iqy-and-data-queries-dqy"
categories:
  - Uncategorized
comments:
  - {"author":"Roede Orm","email":"roede@orange.fr","url":"","date":"2013-11-29T16:30:35Z","content":"Hi,\r\n\r\nI found your write up to fix the power settings on a Samsung laptop. I could reproduce most of it. Just one thing does not work. Sometimes the keyboard backlight stops working. It's not reproducible. It might happen, when battery power is low and the laptop resumes. But I can't reproduce it. \r\nI have an samsung 900x3d with Ubuntu 12.04.1 running Kernel 3.2.0-56-generic-pae, Do you have the same?\r\nI reset the battery from time to time to get it back working. Any ideas what could go wrong?\r\n\r\nKind regards\r\nRoede","parent":0}
---

An alternative approach to just doing CSV export (which is fast and more cross platform/app) is to use web queries (.iqy files that point to fairly HTML) or data queries (.dqy extension but point to the DB)

<b>Pro</b>
<ul>
<li>Supports refreshing data / live data model</li>
<li>Web based integration</li>
<li>Very easy to do</li>
<li>Works with all versions since excel 97</li>
<li>They support linking back into you app via urls (IQY only) so we had cells that said edit that took you back to the web app.</li>
<li>Auto sorting, pivoting is a doddle</li>
</ul>

<b>Cons</b>
<ul>
<li>No support for different sheets in the same book (you can work around this by having a master workbook that aggregates from multiple web queries sheets.) CSV has the same problem.</li>
<li>There are some methods to write back but they all seemed a bit sucky to me (you can create forms that push the data back from specific cells but it's pretty ugle UI wise). I have always ended up writing a tiny bit of VBA that just does a POST to some url (embed this in your master workbook etc) or just an edit link that goes back to the web app</li>
<li>Some annoying pop ups (See below)</li>
</ul>

<b>Notes</b>
<ul>
<li>With DQY you expose username password and database details to user but you get full SQL support / joins etc.</li>
<li>Make sure you set Cache-Control headers ( I use either "public" or "private, must-revalidate" )</li>
</ul>

<b>IQY Example</b>
<pre>
WEB
1
http://server/money.html?query=blah+blah+blah&sortBy=market_value

</pre>

<b>DQY Example</b>
<pre>
XLODBC
1
DRIVER=SQL Server;SERVER=server\blah;UID=regulatory;PWD=l33t;DATABASE=showMeTheMoney
SELECT * FROM trades ORDER BY market_value DESC 

</pre>

<b>Disable Excel Pop.reg</b>
<pre>
Windows Registry Editor Version 5.00

[HKEY_CURRENT_USER\Software\Microsoft\Office\11.0\Excel\Options]
"QuerySecurity"=dword:00000002
</pre>



