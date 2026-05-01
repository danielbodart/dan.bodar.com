---
title: "Agentic Architecture"
date: 2026-04-29T09:00:00Z
slug: "agentic-architecture"
series: "Agentic Engineering"
categories:
  - engineering
  - architecture
  - agents
tags:
  - architecture
  - deployment
  - observability
  - version-control
  - trunk-based-development
  - agentic-engineering
  - triptease
---

This is part 4 of a 5-part series on agentic engineering. [Part 1](/2026/04/26/from-agile-engineering-to-agentic-engineering/) set the values-vs-mechanisms frame. [Part 2](/2026/04/27/from-pair-programming-to-co-steering/) covered pairing and learning. [Part 3](/2026/04/28/harness-engineering-building-a-factory-for-code/) made the case for harness engineering. This one is about the architecture *underneath* the harness — the language, runtime, deployment, and version-control choices that make agents safe.

[Part 5](/2026/04/30/picking-a-language-for-agents/) gets concrete about specific languages.

## Why architecture changes with agents

If humans are no longer the primary readers and writers of code, then architecture should no longer be optimised only for human familiarity. Instead, I want to optimise for:

> fast feedback, strong constraints, low ambiguity, safe deployment, and agent legibility.

The question shifts from:

> What stack do we know?

to:

> What stack gives agents the safest, fastest path to correct execution?

That sounds dramatic, but in practice it just means treating the development environment as a system whose job is to make agent output safe, and choosing tech that does that well.

## Core principle

> Choose systems that make the factory better.

Good agentic architecture improves three things at once:

- the agent's ability to make correct changes
- the human's ability to steer, inspect, and recover
- the system's ability to validate itself

The very best choices help **both humans and agents**. Anything that helps one at the expense of the other should make you suspicious.

## Language and runtime choices

Historically, teams chose languages based on the team's expertise. In agentic engineering, that constraint relaxes a lot. Agents can write competently in languages no one on the team has used before. So we get to evaluate languages by their *system properties*:

- strong typing
- fast compile times
- fast test feedback
- simple language surface area
- good standard library
- predictable tooling
- low dependency risk
- strong ecosystem fit for the problem

A few of these deserve their own paragraph.

### Strong typing as a guide

Strong typing is *feedforward*. It constrains the space of possible expressions and gives agents immediate feedback when they're wrong. That makes typed languages much more attractive than they used to be — not because typing is fashionable, but because every type error is one fewer round-trip the agent has to do at runtime.

### Fast feedback loops

Compile time and type-checking speed matter a lot. Agents iterate quickly, but slow loops multiply cost:

- more waiting
- more tokens
- more failed repair loops
- slower convergence

A theoretically powerful language with slow validation may be worse than a simpler language with very fast feedback. That's a re-ranking I would not have predicted ten years ago.

### Simplicity

Agents tend to do better with languages that have a small, clear surface area. A language doesn't need to be widely known on the team if:

- the agent can write it reliably
- the compiler gives strong feedback
- the runtime fits the domain
- the ecosystem is safe and understandable

Don't be afraid of unfamiliar languages where the properties fit. (Concrete suggestions in [part 5](/2026/04/30/picking-a-language-for-agents/).)

## Batteries included over dependency sprawl

Every dependency is:

- a supply-chain risk
- an API surface the agent may misuse
- a maintenance obligation
- another thing to audit

That pushes me toward:

- strong standard libraries
- minimal external dependencies
- careful, deliberate dependency approval
- research-led library selection

We should not "vibe code" dependencies. Agents can research libraries, compare options, and recommend choices, but the actual decision to take on a dependency should remain a deliberate architectural choice — same bar as it always was, just with better research input.

### Library ecosystems are no longer rigid

In the human-centric world, choosing a language often meant committing to its ecosystem. Missing libraries were blockers. In the agentic world, that constraint relaxes:

- agents can extract just the needed functionality from a library
- they can internalise and adapt portions to fit a different language
- they can rapidly translate code across languages

The implications:

- language choice focuses on system properties (typing, speed, simplicity)
- libraries are no longer rigid constraints
- key capabilities can be transplanted across ecosystems

We're moving from being bound by *what's available* to choosing for *what's possible*.

## Runtime and platform choices

Same logic for platforms. I want platforms that give us more operational capability by default:

- observability
- deploy previews
- side-by-side deployments
- fast rollback
- A/B testing
- integrated security
- simple local development
- low configuration burden

This makes higher-level platforms (PaaS, edge-oriented runtimes) more attractive than raw infrastructure when they reduce the operational glue code we'd otherwise have to write and own.

The key question I keep asking:

> Does this platform reduce the amount of custom machinery we need to build safely?

If yes, it's pulling its weight. If no, every extra config file is a place agents (and humans) can get it wrong.

## Deployment architecture

> Prevention won't catch everything. Recovery has to be first-class.

Non-negotiables in my book:

- side-by-side deployments
- instant rollback
- strong observability
- safe preview environments
- A/B testing or controlled exposure
- clear production health signals

If agents can generate more change, the deployment system has to make change safer. That's not optional.

### Side-by-side deployment

This becomes foundational. It lets us:

- compare implementations
- test branches safely
- expose prototypes to internal users
- run controlled experiments
- avoid big-bang replacement

It also changes how I think about branches: not just a code-management tool, but a **deployment and experimentation primitive**. More on that below.

### Observability

Has to be designed in from day one — not "nice to have", core architecture. The bare minimum:

- logs
- metrics
- traces
- business outcome telemetry
- error reporting
- user-behaviour signals
- deploy/version attribution

Without observability, agents produce change faster than humans can understand its impact. That's the whole game lost in one move.

### Rollback

Rollback should be instant and **boring**. If volume and speed of change are going up, rollback is part of the architecture, not an emergency procedure.

A good agentic system makes it trivial to answer:

- what changed?
- where is it running?
- who or what generated it?
- what metrics moved?
- how do we revert safely?

## Architecture for experimentation

Because exploration is cheap, the architecture should *support* exploration:

- branch previews
- feature flags
- isolated test environments
- synthetic data
- safe sandboxing
- multiple implementations
- rapid customer/internal feedback

Ideas should be allowed to expand and then collapse safely into a chosen path.

## Version control: trunk-based, with a wrinkle

Trunk-based development has historically given us:

- reduced cognitive load
- minimal merge complexity
- encouragement of small, incremental changes
- lower risk during integration and release

All of those are still desirable. But some of the original constraints have shifted:

- in a human-driven system: rework is expensive, large branches are risky, merge conflicts are costly to resolve
- in an agentic system: rework is cheap (agents regenerate easily), parallel exploration is encouraged, branching becomes a natural mechanism for experimentation

That creates a tension worth being explicit about.

### What stays true

- **Merging is still a real risk.** Agents can silently introduce regressions during merges; lost changes are harder to detect at scale.
- **Cognitive simplicity still matters.** Humans must be able to reason about system state.
- **Integration safety is critical.** Especially when many parallel explorations are happening at once.

### What changes

Branches become an **exploration primitive**, not just coordination. It becomes feasible to:

- spin up many parallel implementations
- deploy them side-by-side
- evaluate them empirically

### Likely patterns

I don't think there's a settled answer yet, but the patterns I see emerging:

- **Hybrid models** — trunk for convergence, branches for exploration
- **Branch-per-experiment workflows** — each branch deployable independently, evaluated via A/B testing or internal usage
- **Stronger merge validation** — automated diff validation, regression detection via tests + observability, agent-assisted merge verification

### The open question

If merging remains the primary source of risk, do we optimise around **avoiding merges** — or around making them **provably safe**?

I genuinely don't know yet. Both directions look promising. I expect the answer depends on the kind of system you're building, and the cost of getting it wrong if the merge is bad.

## Heuristics

To compress all of the above into something I can actually use in a meeting:

**Prefer:**

- strong types
- fast compile/test loops
- simple languages
- boring runtimes
- batteries-included ecosystems
- minimal dependencies
- deploy previews
- side-by-side releases
- instant rollback
- excellent observability

**Be cautious with:**

- slow compilers
- complex language semantics
- dependency-heavy ecosystems
- custom infrastructure glue
- manual release processes
- systems that are hard to test locally
- architectures where rollback is difficult

## Open questions

A few I'm still chewing on:

- Should we prefer branches more in an agentic world because they enable parallel exploration?
- How do we make merges provably safe?
- What coding standards are best for *LLM* comprehension rather than human comprehension?
- Which languages produce the lowest agent failure rate per unit of useful output?
- How do we measure token cost, retry rate, and feedback-loop speed as architectural concerns?
- How much should we optimise for standard libraries over ecosystem breadth?
- Which platforms give us the best default safety rails?

## Final thought

Agentic architecture isn't about chasing new technology. It's about choosing systems that make safe automated execution easier.

> In the old world, architecture helped humans build software. In the agentic world, architecture helps humans build systems that safely generate software.

In [part 5](/2026/04/30/picking-a-language-for-agents/) I get specific about which languages I'd reach for, and why.
