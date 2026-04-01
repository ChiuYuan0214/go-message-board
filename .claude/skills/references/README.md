# Shared References

Shared reference material for all skills (feature-dev, refactor, etc.). Start here.

## Reference Files

| File | Purpose | When to read |
|------|---------|--------------|
| [project-structure/README.md](project-structure/README.md) | Service map, all existing entities / types / service functions / routes | Beginning of every task — find reusable code and understand scope |
| [code-style/go-backend.md](code-style/go-backend.md) | Service function patterns, route handlers, GORM tags, naming, imports | Writing any Go code |
| [code-style/typescript-frontend.md](code-style/typescript-frontend.md) | Component pattern, hooks, API layer, types, styled-jsx | Writing any frontend code |

## How to use

1. **Read `project-structure/README.md`** — understand which service owns the task, find existing functions to reuse, check what DB tables and types already exist
2. **Open the specific method file** (e.g. `project-structure/general/services/GetProfileWithId.md`) — verify exact signature and behaviour before reusing
3. **Read the relevant code-style file** — apply the correct patterns when writing new code
