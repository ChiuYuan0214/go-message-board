# Engineering Rules

These rules apply to all tasks — feature development, refactoring, bug fixes, etc.

---

## Terminology

- When the user says `文件` or `doc`, treat that as `references/project-structure/` unless the user clearly points to some other documentation target.

---

## File Management

- **Delete empty files.** If a file's content is entirely removed during a task, delete the file. Do not leave behind a file containing only a package declaration or placeholder comment.

---

## Code Style

> See also: `BACKEND_CODE_STYLE` from `env.md`

- No package-level variables (`var db`, `var cache`, etc.) — use injected struct fields.
- No standalone functions used as service logic — everything is a method on a struct.
- Compile guards (`var _ Interface = (*Impl)(nil)`) at the top of every impl file.
- Named return values in repo methods; bare `return` for brevity.

---

## Architecture

> See also: `DEPIN_ARCHITECTURE` from `env.md`

- Follow dependency order when wiring backend layers. For data-backed features, prefer `entities` → `types` → `repo` → `service` → `route`.
- HTTP status codes belong in `service`, never in `repo`.
- Route handlers should hold the `Router` plus only the service interfaces they actually need. HTTP handlers often need one service; WebSocket handlers may need a small set of coordinating services.
- Register all components in `main.go` via `depin.Set` in bottom-up order.
- If a feature or refactor changes schema, request/response args, API surface, architecture, layering, ownership, or service/repo/route responsibilities, update `references/project-structure/` in the same task.
- Treat `references/project-structure` as a discovery index, not proof of business correctness. Before reusing an existing method for new work, confirm the actual source code still matches the intended business logic.
