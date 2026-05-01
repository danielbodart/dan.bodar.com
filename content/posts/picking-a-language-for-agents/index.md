---
title: "Picking a Language for Agents"
date: 2026-04-30T09:00:00Z
slug: "picking-a-language-for-agents"
series: "Agentic Engineering"
categories:
  - engineering
  - languages
  - agents
tags:
  - typescript
  - kotlin
  - python
  - go
  - elixir
  - zig
  - rust
  - languages
  - agentic-engineering
  - triptease
---

This is the final part of a 5-part series on agentic engineering. [Part 4](/2026/04/29/agentic-architecture/) laid out architectural principles for choosing languages and runtimes. This post puts a handful of worked examples through those principles, just to show what the analysis looks like in practice.

This isn't a list of "the languages I use", and it isn't exhaustive. It's a deliberate move *away* from personal taste — what I happen to know, what's fashionable, what the team is already comfortable with — and toward something more objective: which language properties best fit **the agent** *and* **the problem domain** in front of you. Different problems will pull you toward different answers, and the examples below are picked specifically to make that point.

The heuristics from the previous post still apply: strong typing, fast compile/test loops, simple language surface area, low dependency risk, batteries-included ecosystem, good fit with the harness. With those in mind, here are some illustrative examples.

## TypeScript

Still a strong default, especially on the client.

**Pros**

- widely represented in training data — agents are very fluent here
- strong typing
- huge ecosystem
- natural fit for web products

**Risks**

- high supply-chain risk (using [Bun](https://bun.sh/) or [Deno](https://deno.com/) helps, but doesn't eliminate this)
- fragmented tooling — Bun or Deno cuts through most of this
- single-threaded

Worth noting: Claude Code itself uses Bun + TypeScript. Strong typing, fast feedback, and the largest training corpus in existence is a meaningful combination for the agent — particularly for client-side or full-stack web problems where the domain happens to align with that ecosystem.

## Kotlin

Solid for service work where the JVM is the right runtime.

**Pros**

- strong typing
- mature JVM ecosystem
- good for service architecture
- excellent testing and HTTP support

**Risks**

- slower feedback loops than I'd like
- more complex types == more complex error messages (and more places agents can get tangled)
- JVM overhead
- medium supply-chain risk (better than JS, not great)

Kotlin fits well when JVM context is already a hard requirement and the problem benefits from a typed, expressive language with mature libraries. From a purely agent-first perspective, the feedback-loop story is the weakest part of the picture — agents pay for slow compile/test cycles in retries and tokens.

## Python (with a type checker)

Highly effective for AI/ML, data, and evaluation-heavy systems.

**Pros**

- dominant ecosystem for ML, AI, and data workflows
- fast iteration, low ceremony
- excellent for experimentation, evaluation, and internal tooling
- widely represented in training data

**Considerations**

- requires real discipline to avoid runtime issues
- supply-chain risks similar to JS — large transitive dependency graphs are common

Python is the natural fit where the system interacts closely with models, data, or evaluation loops — the ecosystem is doing real work for you. In an agentic setup it should be paired with **mandatory type checking** (`mypy`, `pyright`). Untyped Python in an agent loop is asking for trouble — every type bug becomes a round-trip cost.

## Go

Strong agent-fit for server-side problems with a simple shape.

**Pros**

- very fast compile times
- simple language
- strong typing
- fast runtime
- excellent standard library
- low cognitive overhead
- low supply-chain risk (few dependencies is the norm because the standard library is so good)

**Risks**

- less expressive
- verbose error handling
- which leads to higher token usage

Go ticks an enormous number of agentic boxes — fast feedback, simple surface area, batteries included. The main cost is verbosity, and verbosity in an LLM world is paid in tokens. The trade-off depends on how chatty the problem domain is.

## Elixir

Strong choice for distributed, concurrent, and fault-tolerant systems.

**Pros**

- excellent model for many small, independent processes
- built-in supervision and fault tolerance ([BEAM](https://www.erlang.org/blog/a-brief-beam-primer/))
- well suited for messaging, orchestration, and real-time systems
- encourages simple, isolated components
- low/medium supply-chain risk (fewer deps than JS/Python/Rust)

**Considerations**

- typing story is improving, but still less strict than the alternatives
- smaller ecosystem

Elixir is interesting because the runtime *resembles an agentic model itself* — many workers, message passing, retries, supervision. For problem domains that match that shape, it reduces the amount of custom orchestration and resilience machinery you'd otherwise have to build. (Reportedly part of why OpenAI ended up using it.)

## Zig

Interesting where low-level control or performance matters — WebAssembly, hardware, native apps.

**Pros**

- explicit and simple language
- good compile times
- exceptional multi-platform support (including WebAssembly)
- extremely low supply-chain risk — dependencies are explicit and there are no transitives
- exceptional runtime performance

**Risks**

- pre-1.0
- small ecosystem (offset by the rich C ecosystem it can use directly)
- not as strong guarantees as Rust
- terrible fit for web services (great for WebAssembly or native apps)

There's a reason Bun chose Zig over Rust even with its pre-1.0 status. I've also used it on a [recent personal project](/2026/02/19/capsper-push-to-talk-voice-dictation-for-linux/) — for a problem in its sweet spot, fast compile times and minimal runtime dependencies do real work for both the agent and the human.

## Rust

Powerful, but potentially expensive for agents.

**Pros**

- extremely strong guarantees
- memory safety
- good WebAssembly support
- excellent C++ support
- exceptional runtime performance

**Risks**

- extremely slow compile times
- complex language model
- higher failure/retry rate in LLM code generation
- higher token and iteration cost
- real supply-chain risk — massive transitive dependencies are the norm, like JS

Rust has real strengths and real costs from an agentic perspective: every retry is expensive in compile time, the type system is rich enough that errors can spiral, and the dependency graphs grow alarmingly fast. The case for Rust gets stronger the more the problem domain demands what it's actually good at — correctness, performance, FFI — and weaker the further you drift from that.

## The shape of the evaluation

The point of going through these examples isn't to land on a list of approved languages. It's to make the evaluation explicit. For any new piece of work, the question I want to ask out loud is roughly:

- **What does the problem actually need?** Concurrency model, performance, ecosystem fit, deployment shape, integration constraints.
- **What are the agent's strengths and failure modes here?** Compile speed, type system, training-data fluency, retry cost, error-message quality.
- **Where do those two answers overlap?**

Whatever sits in the overlap is what you should be reaching for, even if it's a language nobody on the team has used before. Whatever doesn't sit in the overlap should make you uncomfortable, even if it's the language you've used for ten years.

There's a slightly mischievous version of this worth saying out loud: deliberately picking a language you don't know well might actually be useful. The whole shift in the series is humans doing less reading and writing of code and more steering of intent and constraints. If you choose a language you're already fluent in, the temptation to drop down into the editor and "just fix it yourself" is enormous. Choose one you can't fluently read, and that temptation evaporates. You're forced to stay in the steering seat. That's not a reason to pick badly — agent fit and problem fit still come first — but among otherwise-equal options, "I won't be able to resist meddling in this one" is a real argument *against* a familiar language.

The honest version of all this is the part I find hardest. Most language choices in my career have been driven by personal taste or team familiarity. Agentic engineering forces a more boring question: *given this agent and this problem, what's the right tool?* Sometimes that answer is the one I'd have picked anyway. Sometimes it isn't, and that's the more interesting case.

## A note on what I'm *not* optimising for

A few things I deliberately don't put much weight on any more:

- **Team familiarity.** Less of a constraint than it used to be. Agents can write competently in unfamiliar languages; humans can read them with help.
- **Library count.** What matters is whether the standard library and a small set of vetted dependencies cover the problem. Big ecosystems are often a *liability* in an agentic context.
- **Trendiness.** Boring runtimes are good. Boring runtimes have predictable failure modes.

## Wrapping up the series

If you've read all five posts, the core argument is fairly simple:

1. Keep the [Agile values](/2026/04/26/from-agile-engineering-to-agentic-engineering/). Change the mechanisms.
2. Pairing becomes [co-steering](/2026/04/27/from-pair-programming-to-co-steering/). Humans mob on intent, agents mob on execution.
3. Quality emerges from the [harness](/2026/04/28/harness-engineering-building-a-factory-for-code/), not from human discipline alone. Improve the system, not the output.
4. [Architecture](/2026/04/29/agentic-architecture/) gets re-evaluated against agent legibility, fast feedback, and safe deployment.
5. Languages get re-evaluated against the same criteria — chosen for fit between the agent and the problem domain, not for taste.

The throughline: **we're building systems that build software**. The interesting craft has moved up a level. I'm still learning how to do it well, and I fully expect to be wrong about a lot of this by tomorrow, or next week, or this time next year. That's the name of the blog after all.
