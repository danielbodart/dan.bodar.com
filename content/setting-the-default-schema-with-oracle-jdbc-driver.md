---
title: "Setting the Default Schema with Oracle JDBC Driver"
date: 2013-01-24T11:19:24Z
slug: "setting-the-default-schema-with-oracle-jdbc-driver"
categories:
  - Uncategorized
comments:
  - {"author":"Markus W.","email":"markus@nomail.de","url":"http://noweb","date":"2013-11-19T15:12:56Z","content":"You saved my day! Thanks","parent":0}
---

If you use C3PO you can make it do it when it checks the connection out.

As properties:
<pre>c3p0.preferredTestQuery=alter session set current_schema=animals
c3p0.testConnectionOnCheckout=true
</pre>
As Java code:
<pre>ComboPooledDataSource dataSource = new ComboPooledDataSource();
dataSource.setPreferredTestQuery("alter session set current_schema=animals");
dataSource.setTestConnectionOnCheckout(true);
</pre>
Downside is this will happen every time the connection is taken out of the pool

If you are using a JDBC connection yourself you could just do:
<pre>Class.forName("oracle.jdbc.driver.OracleDriver");
Connection connection = getConnection("jdbc:oracle:thin:@//server:1521/instance", "username", "password");
connection.createStatement().execute("alter session set current_schema=animals"));
</pre>
I also posted it to <a href="http://stackoverflow.com/questions/2353594/default-schema-in-oracle-connection-url">StackOverflow</a>