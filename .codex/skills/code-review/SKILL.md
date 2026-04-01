---
name: code-review
description: >
  Code review workflow for go-message-board. Use this skill when the user asks
  for a review, code review, PR review, diff review, or wants a strict pass
  over existing changes without primarily asking for implementation. Trigger on
  keywords such as review, code review, PR review, diff review, 審查, review 一下,
  看看這次變更, or 幫我抓問題.
---

# Code Review Workflow

## Before Anything Else — Load Context

Read these in order:

1. `.codex/skills/env.md`
2. `ENGINEERING_RULES` from `env.md`
3. `PROJECT_STRUCTURE_INDEX` from `env.md`
4. `BACKEND_CODE_STYLE` or `FRONTEND_CODE_STYLE` from `env.md`, depending on the files under review
5. `PROJECT_NOTES` from `env.md` if setup, docker, or local verification matters

Treat `project-structure` as a discovery index only. Confirm real source code before concluding that an existing method or flow is correct.

## Review Standard

Be strict by default.

The goal of review is to prevent the codebase from getting worse over time while still allowing progress. Prefer surfacing a smaller number of high-confidence, high-impact findings over flooding the user with speculative comments.

If multiple implementations are acceptable and the change does not reduce code health, prefer the author's choice.

## Primary Review Order

Review in this order:

1. correctness and behavioral regressions
2. security and auth boundaries
3. data loss, consistency, and persistence risks
4. concurrency, lifecycle, and resource handling
5. API, schema, and compatibility impact
6. observability and operability gaps
7. performance issues that are likely to matter
8. maintainability and readability

Style-only comments should be lower priority unless they violate project rules or make future mistakes significantly more likely.

## What To Look For

### Correctness

- Does the change actually do what the code claims?
- Are edge cases handled: nil, empty, auth failure, missing records, duplicate events, retries, partial failure?
- Are old call paths, jobs, or handlers still pointing at stale behavior?

### Security

- Is auth enforced at the right boundary?
- Can user-controlled input cross trust boundaries unsafely?
- Are tokens, IDs, permissions, or ownership assumptions validated correctly?

### Data and State

- Could this corrupt data, lose messages, double-write, or leave cache and persistence out of sync?
- If the change touches schema, args, or API shape, were docs updated too?

### Concurrency and Lifecycle

- For WebSocket and goroutine code, check cancellation, channel ownership, connection close paths, shared mutable state, and background loops.
- Watch for races around maps, connection objects, caches, and periodic jobs.

### Compatibility

- Does this break existing request/response contracts, DB assumptions, or caller expectations?
- Is any migration or deployment ordering required?

### Tests and Validation

- Are the checks appropriate for the risk level?
- If tests are missing, say what should have been validated.
- Prefer running focused verification when practical instead of guessing.

## Review Output

Findings come first.

For each finding:
- include severity ordering, highest first
- cite exact file paths and line numbers when possible
- explain the user-visible or maintenance impact
- suggest the safer expectation, not a vague preference

After findings, optionally include:
- open questions or assumptions
- a short summary of what you reviewed
- residual risks if no findings were discovered

If there are no findings, say that explicitly, then note any verification gaps.

## Severity Guide

- `P0`: must fix before merge; severe outage, security, or irreversible data risk
- `P1`: likely bug or regression with meaningful impact
- `P2`: real maintainability or correctness risk, but narrower in scope
- `P3`: minor improvement or polish; usually non-blocking

Do not inflate severity. Be willing to be strict, but keep findings evidence-based.

## Review Tactics

- Read the changed code closely; do not trust file names, method names, or summaries.
- For larger changes, inspect the call chain above and below the modified code.
- If a reused helper looks suspicious, open the helper source.
- For multi-layer backend changes, check `repo -> service -> route` boundaries explicitly.
- For this repo, when chat code changes, pay extra attention to WebSocket handshake, token validation, event dispatch, cache/history sync, and background jobs.

## Verification Checklist

- Did I identify the highest-risk behavior first?
- Did I verify actual source behavior instead of inferring from names?
- Did I avoid burying the main issues under style comments?
- Did I note missing tests or unrun checks where they matter?
- If docs should have changed, did I call that out?
