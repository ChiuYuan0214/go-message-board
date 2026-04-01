---
name: refactor
description: >
  General refactor skill for go-message-board. Use this skill whenever the user
  wants to refactor, restructure, optimize, clean up, migrate, or standardize
  existing code. Trigger on keywords such as refactor, 重構, cleanup, 整理,
  optimize, 優化, migrate, standardize, continue refactor, or finish refactor.
---

# Refactor Skill

## Before Anything Else — Load Context

Read these before writing any code:

1. **`.claude/skills/env.md`** — canonical shared paths
2. **`ENGINEERING_RULES`** from `env.md` — rules that apply to every task
3. **`REFERENCES_INDEX`** from `env.md` — service map and reusable references
4. **`BACKEND_CODE_STYLE`** from `env.md` when changing Go code
5. **`FRONTEND_CODE_STYLE`** from `env.md` when changing frontend code

Always read the current state of the relevant files before writing anything.

---

## Goal

Refactor existing code without changing intended behavior unless the user explicitly asks for a behavioral change.

Prefer small, reviewable steps that improve one of these areas:
- architecture consistency
- naming clarity
- duplication removal
- dependency direction
- data access boundaries
- testability
- dead code removal

---

## Scope Analysis

Before editing, identify:
- What behavior must stay the same
- Which files own the current behavior
- Whether reusable abstractions already exist
- Whether the refactor affects API shape, persistence, or side-effects
- What verification is needed to prove the refactor is safe

Then summarize:

```text
Refactor target: <what is being improved>
Behavior preserved: <what must not change>
Affected files: <list>
Main risks: <list>
Verification: <tests/build/manual checks>
```

If the requested refactor is ambiguous, ask only the questions needed to avoid risky assumptions.

---

## Planning

Present a short numbered plan before editing.

Use dependency-aware ordering. For backend structural work, prefer:

```text
entities → types → repo → service → route
```

For other refactors, order changes from lowest-level dependency to highest-level call site.

Keep each step narrow enough to validate independently.

---

## Refactor Rules

- Preserve behavior by default.
- Do not mix refactoring with unrelated feature work.
- Prefer moving logic behind existing abstractions over inventing parallel patterns.
- Keep HTTP concerns in handlers/services, not in repo/data-access layers.
- Avoid package-level mutable state.
- Prefer struct methods and injected dependencies over global access patterns.
- Add compile guards where the codebase already uses them.
- Delete empty files if a refactor removes their last meaningful content.

---

## Implementation Workflow

1. Read the current implementation and identify the boundary to improve.
2. Update low-level contracts first.
3. Move or simplify implementations in small steps.
4. Update callers after the new abstraction is ready.
5. Remove obsolete code paths only after all references are moved.
6. Verify there are no stale references, duplicate implementations, or dead imports.

If the refactor spans multiple files, explain progress in small checkpoints instead of making the user infer what changed.

---

## Validation Checklist

- [ ] Behavior-preserving assumptions are still true
- [ ] No stale call sites remain
- [ ] No duplicate old/new paths remain
- [ ] Imports and interfaces are consistent
- [ ] Relevant build/tests/checks pass, or any unrun checks are clearly noted
