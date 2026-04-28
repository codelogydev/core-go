# core-go

Reusable core library for Go backend services.

## 🚀 Features

- JWT Authentication (generate & validate)
- Auth Middleware (Gin)
- Logger Middleware (Zap structured logging)
- Recovery Middleware (panic handler)
- Standard API Response (success, error, helpers)
- Structured Logger (Zap)
- Environment Config Helper

---

## 📦 Installation

```bash
go get github.com/codelogydev/core-go@v1.0.3
```

## 🛠 Usage

### Initialize Logger

```go
import "github.com/codelogydev/core-go/logger"

func main() {
    logger.Init()
    defer logger.Log.Sync()
}
```

### Setup Server

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/codelogydev/core-go/middleware"
)

r := gin.New()
r.Use(middleware.Recovery())
r.Use(middleware.Logger())
```

### JWT Generate & Validate

```go
import "github.com/codelogydev/core-go/auth"

token, err := auth.GenerateToken(1)

userID, err := auth.ValidateToken(token)
```

### Auth Middleware

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/codelogydev/core-go/auth"
    "github.com/codelogydev/core-go/middleware"
)

protected := r.Group("/api")
protected.Use(middleware.AuthMiddleware())
{
    protected.GET("/me", func(c *gin.Context) {
        userID := auth.GetUserID(c)
        response.Success(c, gin.H{"user_id": userID})
    })
}
```

### Standard Response

```go
import "github.com/codelogydev/core-go/response"

response.Success(c, data)

response.Error(c, 500, "internal server error")
response.BadRequest(c, "invalid request")
response.Unauthorized(c, "missing token")
response.Forbidden(c, "access denied")
response.NotFound(c, "resource not found")
```

### Cache (Redis)

```go
import (
    "context"
    "time"
    "github.com/codelogydev/core-go/cache"
)

if err := cache.Init(os.Getenv("REDIS_URL")); err != nil {
    logger.Log.Warn("redis unavailable", zap.Error(err))
}

cache.Set(ctx, "key", "value", 10*time.Minute)

val, err := cache.Get(ctx, "key")

cache.Delete(ctx, "key")

exists, err := cache.Exists(ctx, "key")
```

## 📁 Project Structure

```
core-go/
├── auth/
│   ├── jwt.go
│   └── helper.go
├── cache/
│   └── redis.go
├── middleware/
│   ├── auth.go
│   ├── logger.go
│   └── recovery.go
├── response/
│   └── response.go
├── logger/
│   └── logger.go
└── config/
    └── config.go
```

## 🔐 Environment Variables

| Key | Description | Default |
|---|---|---|
| `JWT_SECRET` | Secret key for JWT signing | `secret` |
| `REDIS_URL` | Redis connection URL | `redis://localhost:6379/0` |
