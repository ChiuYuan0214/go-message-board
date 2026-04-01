---
name: linting
description: >
  Linting and validation workflow for go-message-board. Use this skill when the
  user asks to lint, check, validate, review build health, run quality checks,
  or verify a service without doing feature development or refactoring. Trigger
  on keywords such as lint, check, validate, build, test, quality, verify,
  檢查, 驗證, 跑測試, 跑 lint, or 健檢.
---

# Linting Workflow

## Before Anything Else — Load Context

Read these in order:

1. `.codex/skills/env.md`
2. `ENGINEERING_RULES` from `env.md`
3. `PROJECT_STRUCTURE_INDEX` from `env.md`
4. `BACKEND_CODE_STYLE` from `env.md` when checking Go code
5. `.codex/skills/references/project-notes.md` when local run or integration setup matters

Use the index to identify the owning service and the smallest relevant command surface before running checks.

## Goal

Run the smallest useful validation set for the requested scope, report failures clearly, and avoid mixing quality checks with unrelated code changes unless the user explicitly asks for fixes.

## Default Workflow

1. Identify the target scope:
   - repo-wide
   - one backend service
   - one package or file group
2. Prefer repo-native commands over invented ones.
3. Start with the cheapest signal:
   - formatting / lint command if the repo has one
   - `go build ./...` for compile safety
   - `go test ./...` when tests exist and are relevant
4. For cross-service behavior or WebSocket verification, prefer the repo's `docker-compose.yml` flow instead of ad hoc local setup.
5. Summarize:
   - commands run
   - pass/fail result
   - highest-signal errors
   - whether code was changed

## Command Selection

- For a backend service, start from that service directory and prefer targeted checks before repo-wide checks.
- If no dedicated linter is configured, use build and test commands as the baseline validation.
- If a check is too expensive for the request, say what you skipped and why.

## Guardrails

- Do not silently "fix while linting" unless the user asked for fixes.
- When a command fails, report the real blocker instead of masking it with fallback guesses.
- If method signatures changed during the work, confirm the matching `interface.go` in that folder stayed in sync.
- Treat reference docs as navigation aids; confirm source code when a failure depends on business logic, not only structure.
