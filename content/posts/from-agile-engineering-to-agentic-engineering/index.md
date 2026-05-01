---
title: "From Agile Engineering to Agentic Engineering"
date: 2026-04-26T09:00:00Z
slug: "from-agile-engineering-to-agentic-engineering"
series: "Agentic Engineering"
categories:
  - engineering
  - agile
  - agents
tags:
  - agile
  - xp
  - agentic-engineering
  - triptease
---

For most of my career, "doing engineering well" has meant doing Agile well. Fast feedback, sustainable pace, simple design, close customer collaboration, all the values from the [Agile Manifesto](https://agilemanifesto.org/) that I've spent twenty-odd years arguing about over coffee. At [Triptease](https://www.triptease.com/) those values still pay rent: they help us deliver, adapt, collaborate, and not burn out.

What's changing is *who is doing the work*.

The hard part of software has never been the typing. It's always been translating fuzzy intent into a working solution: understanding the problem, choosing the design, working out what "correct" actually means, and keeping all of that consistent as things change. That's where the real time has always gone.

What's changed is that agents can now help carry a lot of that load — generating, testing, refactoring, documenting, and reviewing at high speed. The scarce resource moves further upstream: deciding *what to build*, *how to constrain it*, and *how to know it's right*.

> Humans own intent, judgment, and systems. Agents own execution and iteration.

That's the working hypothesis. This post is the first in a series of five where I work through what it actually means in practice.

## Series map

1. **From Agile to Agentic Engineering** *(this one)* — the values vs. mechanisms split, and what changes when execution gets cheap.
2. [From Pair Programming to Co-Steering](/2026/04/27/from-pair-programming-to-co-steering/) — what happens to pairing, mobbing, and learning when agents handle most of the execution.
3. [Harness Engineering: Building a Factory for Code](/2026/04/28/harness-engineering-building-a-factory-for-code/) — TDD, CI, refactoring, and simplicity reinterpreted around the system, not the human.
4. [Agentic Architecture](/2026/04/29/agentic-architecture/) — language, runtime, deployment, and version-control choices that make agents safe.
5. [Picking a Language for Agents](/2026/04/30/picking-a-language-for-agents/) — a concrete tour of TypeScript, Kotlin, Python, Go, Elixir, Zig, and Rust through an agentic lens.

## Separate the values from the practices

Agile named the values clearly, and the practices delivered them well. Each practice — pair programming, TDD, stand-ups, retros, CI — typically delivered several values at once. Pair programming, for example, gave you fast feedback, shared context, mentoring, design discussion, accountability, and social cohesion all at the same time, all through the same mechanism: two humans at one keyboard. That worked precisely *because* the mechanism was constant.

When the mechanism changes, you have to be careful not to throw out the values it was carrying. You have to look at each one separately and ask: what's the best way to preserve this now?

Agile values I still want to keep:

- fast feedback
- shared understanding
- adaptability
- quality
- sustainable pace
- simplicity
- close customer collaboration

What changes is *how we deliver each one*. Humans collaborate on intent, constraints, and decisions. Agents execute, explore, and iterate at scale. The system — tests, specs, pipelines, guardrails — replaces a lot of the manual discipline that used to live in people's heads and habits.

## What still matters (and how it evolves)

### Customer value

Still number one. But everything around it has changed:

- cost of experimentation → near zero
- prototypes → cheap and abundant
- multiple implementations → normal

That unlocks parallel exploration, A/B testing as a default, and "vote with your feet" validation. The shift isn't *more* incremental delivery of a single path — it's rapid exploration of many paths, then convergence on one.

### Sustainable pace

More important than ever, because agentic development is fast, addictive, and high-dopamine. The risk isn't writing too little code, it's overwork disguised as productivity, constant interruption from agent outputs, and cognitive overload.

> Sustainable pace is now about protecting human cognition, not limiting output.

In practice that means controlled context-switching (intentional, not interrupt-driven), explicit downtime between steering cycles, team responsibility for burnout prevention, and agents that don't demand attention continuously.

### Collaboration

Still essential, just one level up. Less co-typing, more co-steering. More spoken alignment, often transcribed. Shared intent becomes the primary artefact, not shared keystrokes. We're communicating about *intent and constraints* rather than the lines of code that implement them.

### Learning

Learning moves up the abstraction layer. From "how do I write this?" to "how do I specify intent clearly, design constraints and invariants, detect weak output, and build feedback loops?" That's a different skill set, and we have to deliberately create apprenticeship paths so juniors actually pick it up.

### Quality

Quality is no longer enforced primarily by human review. It emerges from strong acceptance criteria, comprehensive automated tests, architectural constraints, observability, and agent-to-agent validation loops.

> We do not trust output. We verify with systems.

### Simplicity (harder, not easier)

This is the one that worries me most. Agents tend to add `if/else` spaghetti, ignore separation of concerns, and do the easiest thing in front of them. Simplicity has to be *enforced* through explicit constraints, automated complexity checks, and review loops focused on reduction. It won't happen by accident.

## Who does what

A useful way to think about responsibilities:

**Humans**

- problem framing
- trade-offs and prioritisation
- architecture and constraints
- acceptance criteria and specifications
- defining guardrails, guides, and sensors
- reviewing based on risk (not exhaustively)
- improving the system, not patching outputs
- sensing when agents are struggling

**Agents**

- implementation
- test generation and execution
- refactoring
- documentation
- exploration of multiple approaches
- iterative fix loops
- code review (initial pass)
- running pipelines and validations

## The new discipline: harness engineering

Here's the line that crystallised it for me:

> We are no longer primarily writing code. We are building a factory for generating code safely.

That factory has three core concepts:

- **Feedforward (guides):** specs, types, constraints, standards
- **Feedback (sensors):** tests, metrics, logs, validation systems
- **Guardrails:** escalation points, limits, failure detection

And one rule:

> If something breaks, we don't fix the output — we improve the system.

I'll go deep on harness engineering in [post 3](/2026/04/28/harness-engineering-building-a-factory-for-code/).

## The operating loop

{{< mermaid >}}
flowchart TD
    A[Humans align on intent] --> B[Capture lightweight specs<br/>and acceptance criteria]
    B --> C[Agents explore, implement, validate]
    C --> D[Humans review selectively,<br/>adjust constraints]
    D --> E[System improvements<br/>replace manual fixes]
    E --> A
{{< /mermaid >}}

That's it. Not a big methodology. A loop where humans steer and the system catches what humans no longer see directly.

## Guiding principles

A few that I've found myself coming back to:

- humans don't read or write code directly — they shape what does
- systems scale, individuals don't
- verification replaces trust
- cognitive load is the limiting factor
- exploration is cheap; convergence is valuable
- anything that improves both humans *and* agents is the best kind of choice

## So is this the end of Agile?

No. It's the continuation.

We keep the values. We change the mechanisms. We elevate the discipline.

Agile taught us how to build software with humans. Agentic engineering teaches us how to build *systems that build software*.

The next four posts get specific about what that looks like in practice. First up: what happens to pair programming when one of the pair doesn't need a keyboard? See you in [part 2](/2026/04/27/from-pair-programming-to-co-steering/).
