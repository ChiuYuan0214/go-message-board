# Shared References

Shared reference material for all skills (feature-dev, refactor, etc.). Start here.

Before using the paths below, read [`../env.md`](../env.md) for the canonical shared path names.

## Reference Files

| File | Purpose | When to read |
|------|---------|--------------|
| [project-structure/README.md](project-structure/README.md) | Service map, all existing entities / types / service functions / routes | Beginning of every task — find reusable code and understand scope |
| [project-notes.md](project-notes.md) | Repo-specific notes for local verification and workflow decisions | Read when setup, run, or integration testing matters |
| [code-style/go-backend.md](code-style/go-backend.md) | Service function patterns, route handlers, GORM tags, naming, imports | Writing any Go code |
| [code-style/typescript-frontend.md](code-style/typescript-frontend.md) | Component pattern, hooks, API layer, types, styled-jsx | Writing any frontend code |

## How to use

1. **Read `../env.md`** — resolve `PROJECT_STRUCTURE_INDEX`, `BACKEND_CODE_STYLE`, and other canonical paths first
2. **Read `project-structure/README.md`** — understand which service owns the task, find existing functions to reuse, check what DB tables and types already exist
3. **Read `project-notes.md` when testing/setup matters** — prefer repo-native workflows before inventing ad hoc ones
4. **Open the specific method file** (e.g. `project-structure/general/services/GetProfileWithId.md`) — verify exact signature and behaviour before reusing
5. **Read the relevant code-style file** — apply the correct patterns when writing new code
