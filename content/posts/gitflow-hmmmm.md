---
title: "GitFlow hmmmm"
date: 2016-01-08T15:39:57Z
slug: "gitflow-hmmmm"
categories:
  - pairing
  - vcs
comments:
  - {"author":"Grant","email":"grant.j.sheppard@gmail.com","url":"","date":"2016-01-09T09:20:05Z","content":"The thing I dislike about Gitflow is that you test your changes on a branch but then merge back to master before releasing. Since you effectively now have two commits, you cannot be absolutely sure that what is deployed is what you tested. That makes me sad :-(","parent":0}
  - {"author":"Scott Miller","email":"scott.miller171@gmail.com","url":"","date":"2016-03-22T17:00:06Z","content":"GitFlow is great for waterfall, not for Agile, and especially not for CD","parent":0}
  - {"author":"Tim","email":"timorrusty@gmail.com","url":"","date":"2018-07-21T15:11:43Z","content":"Dan, Thanks for sharing your thoughts on this workflow.  I heard about you from Jez Humble.  I'm curious about your comment in the middle of the post.  \r\n\"GitHub flow is so nearly right except for the last two steps are the wrong way round (Deploy then Merge!):\"\r\nFrom my reading of the GitHub flow, the steps GitHub proposes are in the Deploy/Merge order.  Has that changed in the last two years?  Are you implying that you should Merge to master/then Deploy?\r\nThanks in advance for the clarification.","parent":0}
  - {"author":"dan","email":"dan@bodar.com","url":"","date":"2019-02-10T08:35:57Z","content":"Yes exactly, nothing has changed in the last 2 years. Trunk based development has been around for at least 20 years and is still the reserve of the highest performing teams. See https://trunkbaseddevelopment.com/","parent":0}
---

After reading <a href="https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow" target="_blank">Atlassian</a> worflow comparison and Vincent Driessen <a href="http://nvie.com/posts/a-successful-git-branching-model/" target="_blank">original post</a> about GitFlow I have come to realise a couple of worrying things:
<ul>
 	<li>I incorrectly assumed GitFlow was the same thing as <a href="https://guides.github.com/introduction/flow/index.html">GitHub flow</a> (fork project, do work then pull request)</li>
 	<li>This model appears to be popular but it seems totally archaic to me
<ul>
 	<li>Requires lots of merging especially if you refactor at lot.</li>
 	<li>Doesn't do CD</li>
 	<li>Requires lots of manual work</li>
 	<li>Master and develop seem the wrong way around to me</li>
 	<li>How many branches?</li>
</ul>
</li>
 	<li>I use pull requests but don't use feature branches which the Atlassian article implies requires feature branches</li>
 	<li>Most places I've seen, feature branches do live on origin (unlike the original post)</li>
</ul>
<a href="https://guides.github.com/introduction/flow/index.html">GitHub flow</a> is so nearly right except for the last two steps are the wrong way round (Deploy then Merge!):
<blockquote>Now that your changes have been verified in production, it is time to merge your code into the master branch.</blockquote>
<b>This is my prefered workflow:</b>
<ul>
 	<li>Master is always trunk or head where all new development happens</li>
 	<li>Every single check-in triggers a build (and tag with auto increment minor version) and is expected to be production code
<ul>
 	<li>If possible every build is automatically released but if not then a single click by an authorised user would make that release public</li>
</ul>
</li>
 	<li>If old major versions are supported by team they are on branches (but the rest is the same, i.e every check is a release etc)</li>
 	<li>Hot fixes are just another commit to either master or the branch. i.e nothing special</li>
 	<li>Done means in production and you have monitored it with your own eyes! You don't start new work until you have seen your old work live</li>
 	<li>If you are on the core team:
<ul>
 	<li>If you pair you can commit to master or branch directly</li>
 	<li>If you solo you should get code review or pull request</li>
</ul>
</li>
 	<li>If you are not on core team you pull request from your fork</li>
</ul>
Is it just me that thinks GitFlow doesn't look very "flow" based, more like a lot of manual busy work like the old days. Please report your counter experiences or alternate workflows...