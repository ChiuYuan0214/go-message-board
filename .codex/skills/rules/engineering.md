# Engineering Rules

These rules apply to all tasks — feature development, refactoring, bug fixes, etc.

---

## File Management

- **Delete empty files.** If a file's content is entirely removed during a task, delete the file. Do not leave behind a file containing only a package declaration or placeholder comment.

---

## Code Style

> See also: `BACKEND_CODE_STYLE` from `.codex/skills/env.md`

- No package-level variables (`var db`, `var cache`, etc.) — use injected struct fields.
- No standalone functions used as service logic — everything is a method on a struct.
- Compile guards (`var _ Interface = (*Impl)(nil)`) at the top of every impl file.
- Named return values in repo methods; bare `return` for brevity.

---

## Architecture

> See also: `DEPIN_ARCHITECTURE` from `.codex/skills/env.md`

- Follow dependency order when wiring backend layers. For data-backed features, prefer `entities` → `types` → `repo` → `service` → `route`.
- HTTP status codes belong in `service`, never in `repo`.
- Each `routes` handler holds exactly one service interface + the `Router`.
- Register all components in `main.go` via `depin.Set` in bottom-up order.
