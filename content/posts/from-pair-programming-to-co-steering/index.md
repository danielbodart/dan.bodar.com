---
title: "From Pair Programming to Co-Steering"
date: 2026-04-27T09:00:00Z
slug: "from-pair-programming-to-co-steering"
series: "Agentic Engineering"
categories:
  - engineering
  - agile
  - agents
tags:
  - agile
  - xp
  - pair-programming
  - mobbing
  - agentic-engineering
  - triptease
---

This is part 2 of a 5-part series on agentic engineering. [Part 1](/2026/04/26/from-agile-engineering-to-agentic-engineering/) made the case that Agile values still matter; the mechanisms are what change. This post takes the most contested of those mechanisms — pair programming — and looks at what happens to it when one of the "pair" doesn't need a keyboard.

I've been a fan of pair programming for a long time. At Triptease we've leaned on it because it improves quality, accelerates learning, spreads context, builds resilience, and quietly does a lot of work for team cohesion. Those outcomes still matter.

What's changed is where the execution happens. The hard part of software has always been **judgment, product understanding, architectural intent, and attention** — translating intent into a working solution that stays correct over time. Typing was never the slow bit. Now that agents can generate, test, refactor, document, and review code quickly, even more weight lands on the human side of that translation.

> Humans mob on intent. Agents mob on execution.

## What pairing gave us

Pair programming delivered a lot of different things at once:

- real-time review / fast feedback
- shared context and ownership
- mentoring and learning
- design discussion
- implementation
- social connection
- accountability
- normalisation of style and approach

If we just say "we do less pair programming now", we lose all of those at once. That's the trap. Instead, the move is to keep each benefit and ask: *what's the best mechanism for it now?*

In an agent-first workflow, my answer is roughly:

- humans collaborate primarily through spoken design and product discussion, and shared steering of agents
- agents handle more of the execution: implementation, tests, refactors, docs, iterative fix loops
- recurring issues get turned into reusable skills, templates, checks, and guardrails — not ongoing manual intervention

That last point matters, and it's harder than it sounds. The temptation when you see something wrong is to just fix it — directly, in the code, and move on. Resist that. The discipline is to look at the problem, work out *how* you'd fix it manually, and then deliberately not do that. Apply the meta-version instead: a check, a constraint, a guardrail, a clearer spec, a better example in the harness. Something that means this class of mistake can't happen again, whether you're watching or not.

Manual fixes don't compound. System fixes do. Every time you do the system fix instead of the easy fix, the harness gets a little smarter and the next class of problem gets a little smaller.

## What we still value (and how we preserve it)

### Learning

Shifts from "how do I write this code?" to:

- how do I specify intent clearly?
- how do I encode constraints and invariants?
- how do I detect weak or risky output?
- how do I design feedback loops and evaluations?

This is a real skill, and not one you pick up by accident.

### Social cohesion

Preserved through **intent-mobbing**: shared decisions, shared rationale, shared ownership. The conversation is still happening, often more of it. It just isn't happening hunched over one keyboard.

### Resilience

Comes from spreading context *and* making knowledge legible and durable: docs, runbooks, standards, acceptance criteria, automated checks. If the only place a piece of context lives is in one engineer's head, it's a single point of failure regardless of whether agents are involved.

### Speed

System throughput goes up when humans spend more time on decisions and quality thresholds, while agents handle iteration. That sounds like a slogan; in practice it's the difference between a 30-minute "is this approach right?" call and an afternoon of co-typing through one approach.

### Quality

Still depends on review — just increasingly enforced by:

- clear acceptance criteria
- strong automated tests
- architecture constraints
- observability and logging standards
- agent-to-agent review and validation loops

## The operating model

**Human work**

- problem framing and trade-offs
- architecture and constraints
- risk judgment
- acceptance criteria
- evolving guardrails and standards
- listening for agent "smells" — the quiet signs the agent is going off the rails

**Agent work**

- implementation
- test generation and execution
- refactoring
- documentation
- review and repeated fix loops

### Practical loop

{{< mermaid >}}
flowchart TD
    A[Humans align on current intent<br/>often spoken, transcribed] --> B[Capture lightweight plan<br/>+ acceptance criteria]
    B --> C[Agents research, implement,<br/>explore, validate]
    C --> D{Risk-based review<br/>or escalation?}
    D -->|low risk| E[Update intent / constraints]
    D -->|high risk or stuck| F[Manual pairing]
    F --> E
    E --> G[Turn recurring failures<br/>into durable rules / checks]
    G --> A
{{< /mermaid >}}

A few things I'll call out about that loop:

- Step 1 is *current* intent, not final design. The whole point is that intent is allowed to evolve.
- Step 3 says "humans should not babysit". This is a real discipline. If you're hovering over the agent watching every token, you're back to co-typing — slowly and badly.
- Step 5 is the harness-engineering move: every recurring failure becomes a system-level fix, not a one-off correction.

## FAQ

A few questions I've fielded enough times to write down.

### Is this "stop pairing"?

No. It's shifting the *default* from co-typing to **co-steering**, while still pairing or mobbing manually when that's the best tool.

### When should we still pair while coding?

- high-risk changes (security, payments, data integrity)
- highly ambiguous behaviour
- complex refactors where the design is evolving mid-flight
- intentional skills-building for newer engineers

### How do we trust agent output?

We don't "trust". We **verify with the system**: tests, constraints, observability, repeatable validation. If you find yourself trusting an agent because it sounded confident, that's a system gap, not a personality flaw.

### Will juniors learn less?

Only if we're careless. We need explicit apprenticeship: guided tasks, intentional manual practice, and strong review/feedback — plus teaching how to steer agents well. The risk isn't agents per se; it's letting "the agent does it" become the answer to questions juniors should still be wrestling with.

### What prevents one person becoming the "AI wizard"?

Make plans, constraints, and decisions **shared artefacts** — specs, docs, checklists, repo-local rules — not private prompts. If the way work gets done lives in one engineer's chat history, you've rebuilt the silo problem with extra steps.

### Who is accountable?

Humans. Delegating execution doesn't delegate responsibility. The agent didn't ship a bug to production; *we* did, with the agent as a tool.

## The short version

We keep the value of pairing. We just stop assuming that two humans at one keyboard is the only mechanism that delivers it. Most days, the most useful thing a senior engineer can do is shape the intent, the constraints, and the system — and let agents handle most of the execution.

Next up in [part 3](/2026/04/28/harness-engineering-building-a-factory-for-code/): the system itself. What does it actually mean to build a *factory* for code?
