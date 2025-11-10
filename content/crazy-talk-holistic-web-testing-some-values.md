---
title: "Crazy Talk - Holistic Web Testing - Some values..."
date: 2010-06-01T12:20:59Z
slug: "crazy-talk-holistic-web-testing-some-values"
categories:
  - Uncategorized
  - crazy talk
---

(Previously on TWSDEV)

As the "crazy" guy behind the in-memory / out-of-container acceptance testing on a number of java/.net projects, I think it's important I explain to people the "Why" and the forces / constraints I am trying to balance. But first I want to quickly lay down my beliefs and values:

- I believe in testing as much as possible (UI included)
- I believe tests must add more value than they cost (Measure it!)
- I value tests that are fast and are resilient to change more than tests that take a long time to run and are brittle.
- When refactoring a feature I value acceptances tests and integration test over unit tests.
- When designing/exploring a new interface / object interaction I value unit tests over acceptance tests to help guide me.
- I believe that QA's are so much better at finding bugs than DEVs but worse at writing code / abstractions