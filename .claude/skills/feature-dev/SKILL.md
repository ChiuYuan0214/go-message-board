---
name: feature-dev
description: Feature development workflow for go-message-board. Guides bottom-up implementation (entities → types → services → routes) with requirement analysis, scope evaluation, and step-by-step confirmation. Use this skill whenever the user wants to add a new feature, implement an endpoint, build something new, or extend the system — even if they don't use the word "feature". Trigger on: add, implement, 新增, feature, build, 做一個, create.
---

# Feature Development Workflow

## Before Anything Else — Load Context

Read **`.claude/skills/references/README.md`** first — it points you to what to read next.

The shared references directory is structured as:

```
.claude/skills/references/
├── README.md                             ← start here
├── code-style/
│   ├── go-backend.md                     ← service function patterns, route handlers, GORM, naming
│   └── typescript-frontend.md            ← component pattern, hooks, API layer, styled-jsx
└── project-structure/
    ├── README.md                         ← index of all existing methods (find reusable code here)
    ├── general/entities, types, services, routes/
    ├── security/services, routes/
    ├── chat/types, services/
    └── stream/types, services/
```

**Loading order for a typical backend feature:**
1. `.claude/skills/references/project-structure/README.md` — understand service ownership, find existing functions to reuse
2. `.claude/skills/references/project-structure/<service>/services/<FunctionName>.md` — check exact signature when reusing
3. `.claude/skills/references/code-style/go-backend.md` — apply correct patterns when writing new code

**For frontend work:** also read `.claude/skills/references/code-style/typescript-frontend.md`

Only open actual source files when you need implementation details beyond what the reference docs show. For the full MySQL schema: `mysql/init_script.sql`.

---

## Input Format

```
Feature: <what the feature does, from a user's perspective>
Service: <which backend service(s), or "frontend", or "all"> (optional)
Notes: <any constraints, edge cases, or related features> (optional)
```

If the user doesn't provide a structured input, extract what you can and proceed.

---

## Phase 1 — Scope Analysis

Using `references/project-structure.md` and `references/services/README.md` (already loaded), identify:
- Which existing service functions can be reused — call them directly, don't re-implement
- Which DB tables are relevant; whether a new table is needed
- Which types already exist vs need to be created
- Where in the folder structure the new files belong

Only open individual source files if you need implementation details beyond what the reference docs show.

Then produce:

```
Feature: <one-sentence summary>
Affected service(s): <list>
New DB tables needed: <yes/no — if yes, list columns>
Affected layers:
  - entities: <new or modified structs, or "none">
  - types: <new request/response types needed>
  - services: <new functions>
  - routes: <new endpoints or handler methods>
Frontend changes: <yes/no — if yes, what>
Data flow: <request → service → DB → response>
```

---

## Phase 2 — Clarification

Ask questions needed to resolve ambiguity before writing code. Focus on:

- **Authorization**: who can call this endpoint? Auth required?
- **Ownership**: does the resource belong to a user? Can others access it?
- **Input validation**: what fields are required vs optional?
- **Edge cases**: what if the resource doesn't exist, or the caller isn't the owner?
- **Side-effects**: does this affect other resources (e.g., deleting an article also deletes votes/comments)?

Only ask what actually affects what you'll write.

> Pause here. Wait for answers before proceeding.

---

## Phase 3 — Development Plan

Present a numbered plan in bottom-up dependency order:

```
Plan:
1. <file path> — <what and why>
2. <file path> — <what and why>
...

Order: entities → types → services → routes
```

> Pause here. Confirm the plan before writing any code.

---

## Phase 4 — Incremental Development

Implement **one file at a time**, strictly bottom-up:

**Go backend order:**
1. `entities/` — GORM entity structs (only if a new table is needed)
2. `types/` — request/response type structs
3. `services/` — business logic functions
4. `routes/` — HTTP handlers and route registration

**Frontend order:**
1. `types/` — TypeScript interfaces
2. `api/` — API layer functions
3. component / page

After each file, show the full file content and ask:

```
File: <path>
Changes:
  - <what was added and why>
---
<full file content>
---
OK to continue to <next file>?
```

Wait for explicit approval (`ok`, `好`, `繼續`, `next`, `yes`) before the next file.

---

## Conventions

All patterns are documented with real code examples in the reference files:
- Go backend → `.claude/skills/references/code-style/go-backend.md`
- TypeScript frontend → `.claude/skills/references/code-style/typescript-frontend.md`
