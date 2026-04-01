# Go Backend Code Style

Applies to all four backend services. Patterns extracted from the existing codebase.

---

## Service Functions

**Return signature:** `(data, int)` where int is HTTP status (`0` = success), or `(string, int)` for error message + status, or a single value like `bool` / `uint64` when simpler.

```go
// Returns data + status
func GetArticle(userId uint64, articleId string) (*types.Article, int) {
    var article types.Article
    err := db.Raw(`...`, userId, userId, articleId).Scan(&article).Error
    if err != nil {
        log.Println(err)
        return nil, http.StatusInternalServerError
    }
    return &article, 0
}

// Returns error message + status
func UpdateArticle(userId uint64, articleId uint64, data *types.UpdateArticleData) (string, int) {
    var article types.Article
    if err := db.Where("article_id = ?", articleId).First(&article).Error; err != nil {
        return "article not exist.", http.StatusBadRequest
    }
    if userId != article.UserId {
        return "user incorrect.", http.StatusBadRequest
    }
    // ...
    return "", 0
}

// Returns single value (0 = error)
func InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) uint64 {
    newArticle := types.Article{ UserId: userId, Title: article.Title, ... }
    err := db.Create(&newArticle).Error
    if err != nil {
        log.Println(err)
        return 0
    }
    return newArticle.ArticleId
}
```

**Error handling rule:** always `log.Println(err)` immediately on error, then return. Never bubble `error` up to the route layer.

---

## Route Handlers (Gin)

```go
// 1. Init function registers routes
func initVote(router *gin.Engine) {
    vh := VoteHandler{}
    router.POST("/vote", middleware.Auth(), vh.add)
    router.PUT("/vote", middleware.Auth(), vh.update)
}

// 2. Empty handler struct
type VoteHandler struct{}

// 3. Lowercase receiver methods
func (vh *VoteHandler) add(c *gin.Context) {
    // 4. Extract userId from context (set by Auth middleware)
    val, _ := c.Get("userId")
    userId := val.(uint64)

    // 5. Parse body
    data := &NewVoteData{}
    message, status := utils.ParseBody(c.Request.Body, data)
    if message != "" {
        c.JSON(status, gin.H{"status": "fail", "message": message})
        return
    }

    // 6. Input validation
    if data.SourceId == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "sourceId cannot be empty."})
        return
    }

    // 7. Call service
    message, voteId := services.Vote(userId, data.SourceId, data.Score, &data.VoteType)
    if message != "" {
        c.JSON(int(voteId), gin.H{"status": "fail", "message": message})
        return
    }

    // 8. Success response
    c.JSON(http.StatusOK, gin.H{"status": "success", "id": voteId})
}
```

**Response shape:**
- `gin.H{"status": "success", "data": item}` — single item
- `gin.H{"status": "success", "list": items}` — array
- `gin.H{"status": "success", "id": id}` — newly created ID
- `gin.H{"status": "success"}` — no payload
- `gin.H{"status": "fail", "message": "reason."}` — error (message ends with `.`)

**`isErr` helper** (defined in `routes/utils.go`):
```go
func isErr(err error) bool {
    if err != nil {
        log.Println(err)
        return true
    }
    return false
}
// Usage:
if isErr(err) {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "body format was wrong."})
    return
}
```

---

## GORM Entities

```go
// Combined json + gorm tags on every field
type Comment struct {
    CommentId  uint64    `json:"commentId"  gorm:"primaryKey"`
    UserId     uint64    `json:"userId"     gorm:"column:user_id"`
    ArticleId  uint64    `json:"articleId"  gorm:"column:article_id"`
    Title      string    `json:"title"      gorm:"column:title"`
    Edited     bool      `json:"edited"     gorm:"column:edited"`
    UpdateTime time.Time `json:"updateTime" gorm:"column:update_time"`
}

// Computed / joined fields that don't map to a column use gorm:"-"
type Article struct {
    ArticleId uint64 `json:"articleId" gorm:"primaryKey"`
    Author    string `json:"author"    gorm:"-"`   // joined, not a real column
    VoteUp    int32  `json:"voteUp"    gorm:"-"`   // aggregated via subquery
}
```

No validation tags. No doc comments.

---

## Naming Conventions

| Thing | Pattern | Example |
|-------|---------|---------|
| Service function | Verb-first PascalCase | `GetArticle`, `InsertArticle`, `DeleteArticle` |
| Route handler struct | `FooHandler` | `ArticleHandler`, `VoteHandler` |
| Handler methods | lowercase | `(ah *ArticleHandler) get(c *gin.Context)` |
| Local variables | camelCase | `userId`, `articleId`, `targetUserId` |
| Error messages | lowercase, end with `.` | `"body format was wrong."` |

---

## Imports

Three groups, blank-line separated:

```go
import (
    "log"           // 1. stdlib
    "net/http"

    "github.com/gin-gonic/gin"  // 2. third-party
    "gorm.io/gorm"

    "general/services"          // 3. local module
    "general/types"
    "general/utils"
)
```

---

## Package-level DB and Cache

Service files use the package-level `db` and `cache` declared in `services/base.go` — never redeclare them:

```go
// services/base.go — already exists, don't touch
var db *gorm.DB
var cache *types.RedisCache

// In your new service file:
package services

import "gorm.io/gorm"

func GetFoo(id uint64) (*types.Foo, int) {
    // just use db directly
    err := db.Where("foo_id = ?", id).First(&foo).Error
    ...
}
```

---

## Route Registration

Always add new route groups to `InitRouter` in `routes/base.go`:

```go
func InitRouter(router *gin.Engine) {
    router.Use(middleware.Cors)
    initArticle(router)
    initVote(router)
    initFoo(router)   // ← add your new group here
}
```
