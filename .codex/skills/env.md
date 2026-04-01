# Skill Environment

Use these canonical paths when a skill needs to load shared project guidance.

## Path Roots

- `SKILLS_ROOT`: `.codex/skills`
- `RULES_ROOT`: `.codex/skills/rules`
- `REFERENCES_ROOT`: `.codex/skills/references`

## Shared Files

- `ENGINEERING_RULES`: `.codex/skills/rules/engineering.md`
- `REFERENCES_INDEX`: `.codex/skills/references/README.md`
- `BACKEND_CODE_STYLE`: `.codex/skills/references/code-style/go-backend.md`
- `FRONTEND_CODE_STYLE`: `.codex/skills/references/code-style/typescript-frontend.md`
- `PROJECT_STRUCTURE_INDEX`: `.codex/skills/references/project-structure/README.md`
- `DEPIN_ARCHITECTURE`: `.codex/skills/references/architecture/depin-pattern.md`

## Usage

- Read this file first when a skill refers to shared paths.
- All project-specific skills should load `ENGINEERING_RULES` immediately after this file, before reading other project references.
- Reuse these canonical paths instead of hardcoding duplicated path strings in each skill.
- Resolve all shared references from this file before opening other docs.
