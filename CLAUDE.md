# go-message-board

Four Go microservices + Next.js 14 frontend.

| Service | Port | Stack | Role |
|---------|------|-------|------|
| `backend/general/` | 8080 | Gin + GORM | Articles, comments, votes |
| `backend/security/` | 7080 | net/http | Auth, profiles |
| `backend/chat/` | 9080 | Gorilla WebSocket | Real-time chat |
| `backend/stream/` | 5000 | Gin | Live streaming |
| `frontend/` | 3000 | Next.js 14 / TypeScript | UI |

**DB:** MySQL · MongoDB (chat) · Redis (cache)  
**Run:** `docker-compose up` — env files in `/env/`

---

## Skills

> Read the relevant skill file before starting any of the following tasks.

| Task | Trigger | Skill file |
|------|---------|-----------|
| Writing / editing code | _(any code task)_ | `.claude/skills/code-style.md` |
| Bug fix | `fix` `bug` `debug` `error` `crash` `修` `問題` `報錯` | `.claude/skills/bug-fix.md` |
| Feature development | `add` `implement` `新增` `feature` `build` `做一個` | `.claude/skills/feature-dev/SKILL.md` |
| Refactor | `refactor` `重構` `cleanup` `整理` `optimize` `優化` | `.claude/skills/refactor/SKILL.md` |
