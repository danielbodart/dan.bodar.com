---
title: "Harness Engineering: Building a Factory for Code"
date: 2026-04-28T09:00:00Z
slug: "harness-engineering-building-a-factory-for-code"
series: "Agentic Engineering"
categories:
  - engineering
  - agile
  - agents
tags:
  - tdd
  - ci
  - refactoring
  - simplicity
  - agentic-engineering
  - triptease
---

This is part 3 of a 5-part series on agentic engineering. [Part 1](/2026/04/26/from-agile-engineering-to-agentic-engineering/) made the case that we keep Agile values and change the mechanisms. [Part 2](/2026/04/27/from-pair-programming-to-co-steering/) walked through what that looks like for pairing and learning. This post is about the thing that does most of the work once agents are doing the execution: **the harness**.

## The reframe

> We are no longer primarily writing code. We are building a factory for generating code safely.

That sentence is doing a lot of work, so let me unpack it.

In the old model, "the system" was mostly the production system — the thing the customer interacts with. The development environment was scaffolding around it: an IDE, a test runner, a CI pipeline, the occasional linter someone added in 2019.

In the new model, that scaffolding *is* the product as far as quality is concerned. It's what makes agent output safe. If a class of mistake gets through, the question isn't "why did the agent do that?" — it's "why didn't the system catch it?"

I think of the harness as having two concerns:

- **Feedforward (guides):** specs, types, constraints, standards. Things that tell the agent what "correct" looks like *before* it writes the code.
- **Feedback (sensors and guardrails):** tests, metrics, logs, validation systems, limits, failure detection. Most of this feeds straight back to the agent so it can self-correct. Some of it — the guardrail end of the spectrum — escalates straight to humans instead, stop-the-line style, when something is too risky for the agent to keep iterating on.

{{< mermaid >}}
flowchart LR
    G[Feedforward<br/>specs, types,<br/>constraints, standards] --> A[Agent execution]
    A --> S[Feedback<br/>tests, metrics, logs,<br/>validation, limits,<br/>failure detection]
    S --> A
    S -. escalate .-> H[Human]
    H --> G
{{< /mermaid >}}

Once you see code-generation as a process running inside a control system, a lot of XP practices start to make sense in a slightly different way. Let me walk through them.

## Practices, reinterpreted

### TDD

Still fundamental, but rebalanced.

- Humans define **acceptance criteria** and **specifications**.
- Agents write tests first, then implement to satisfy them.

The new emphasis is on:

- high-quality in-memory test doubles
- massive test coverage (which is genuinely cheap now)
- testing at scale and volume

The bigger shift is what becomes possible. We've always known testing is good. We haven't always had the energy or the discipline to write *as much* of it as we'd like, especially the kinds that take real effort to set up. Unit tests, integration tests, acceptance tests, smoke tests, performance tests, property-based tests, mutation tests, fuzz tests — every one of those is now a single prompt away.

So the question changes. Instead of "is it worth writing this kind of test?", it becomes: "what kind of testing would actually help me here?" Spot the gap, ask for it, and the harness gets a little better. Tests become the *primary feedback signal* the agent runs against, and the bar for "enough testing" can finally move where it always should have been.

### Continuous Integration

More critical than ever. CI is the central nervous system of the factory.

- constant execution loops
- fast feedback cycles
- a mix of deterministic and non-deterministic checks (yes, some of those checks may now be other agents)

If CI is slow or flaky, every cycle of agent iteration multiplies the cost. Every minute on the build is a minute of agent waiting plus tokens spent re-trying.

### Refactoring

Becomes:

- continuous
- agent-driven
- often automatic

But it has to be guided by **simplicity constraints**, **architectural rules**, and **regression safety nets**. Otherwise you get refactoring that makes things "different" without making them better — a rearrangement without a reduction.

### Release early, release often

Still important, but expect more of it. Code is cheap, releases get more frequent, batch sizes can sometimes get *bigger* simply because there's more to ship. That makes side-by-side deployment and instant rollback non-negotiable. (More on that in [part 4](/2026/04/29/agentic-architecture/).)

### Coding standards

Still required, but they evolve. Some are now optimised for **LLM understanding**, not just human readability. There are real trade-offs to feel out:

- verbosity vs. token cost
- naming clarity vs. length
- "idiomatic" vs. "easy for the agent to get right first time"

I don't think there are clean answers here yet, and a lot of it will be experimentation.

### Collective ownership

Strengthened through **shared specs, shared constraints, shared systems**. What we want to avoid is:

- prompt silos (private chats that produce code only one person can extend)
- AI wizards (one engineer who is mysteriously much faster because of bespoke tooling nobody else can use)

The cure for both is the same: make it a repo-local artefact, not a private one.

### Work in Progress

This is the practice that I think gets reframed *most*.

The old model said: limit unfinished work.

The new model says: **limit cognitive load**.

We should expect:

- bursts of parallel exploration
- collapse into sequential convergence

That's a different shape of WIP. You might have five branches running in parallel for a few hours and then collapse to one — the unfinished work *was the point*. The discipline is doing intentional expansion and contraction cycles, and actively garbage-collecting the abandoned branches once you've converged. Otherwise the cost of "kept around just in case" creeps back in.

### Customer feedback

Becomes more important and more frequent. More prototypes to evaluate. More direct interaction. Cheap exploration is only useful if you actually expose the explorations to real signal.

### Retrospectives

Mostly unchanged, structurally. They're a human process about a human team. The only real change is that transcription tools mean you don't lose the bits people said off the cuff.

### Self-organising teams

Still true, with new responsibilities. Teams must:

- design feedback loops
- detect agent struggle
- build escalation paths

The hard part:

> Humans no longer feel the problem directly — they have to sense it indirectly.

That's a real shift in skill. When you're hand-typing every line, you feel the friction immediately. When agents are typing, the friction shows up as token cost, retry rate, weird diffs, buggy merges, slow tests, vague commits. Learning to read those signals is part of the new craft.

## The new rule

If I had to compress harness engineering down to one rule, it would be:

> If something breaks, we don't fix the output. We improve the system.

That's the move from craftsperson to factory operator. The output is a symptom. The interesting question is always "what's missing in the harness that let this happen?"

It also means *any* hour you spend on the harness compounds. Every guardrail you add catches a class of mistake forever. Every spec you tighten removes a class of ambiguity forever. Manual fixes don't compound. System improvements do.

## Where this leads

If the harness is the new primary product of an engineering team, then the *architecture choices around it* — language, runtime, deployment platform, version control, observability — start to matter for completely different reasons than they used to. That's [part 4](/2026/04/29/agentic-architecture/).
