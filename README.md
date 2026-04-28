# core-go

Reusable core library for Go backend services.

## 🚀 Features

- JWT Authentication (generate & validate)
- Auth Middleware (Gin)
- Logger Middleware (Zap structured logging)
- Recovery Middleware (panic handler)
- Rate Limiter Middleware (per-IP, token bucket)
- Standard API Response (success, error, helpers)
- Structured Logger (Zap)
- Redis Cache (Set, Get, Delete, Exists)
- Storage (S3 / MinIO — Upload, GetURL, Delete)
- Mailer (SMTP transactional email)
- Pagination (Params, Meta, Response)
- Environment Config Helper

---

## 📦 Installation

```bash
go get github.com/codelogydev/core-go@v1.0.5
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

### Storage (S3 / MinIO)

```go
import (
    "context"
    "time"
    "github.com/codelogydev/core-go/storage"
)

storage.Init(
    os.Getenv("STORAGE_ENDPOINT"),
    os.Getenv("STORAGE_ACCESS_KEY"),
    os.Getenv("STORAGE_SECRET_KEY"),
    os.Getenv("STORAGE_USE_SSL") == "true",
)

storage.EnsureBucket(ctx, "my-bucket", "ap-southeast-1")

storage.Upload(ctx, "my-bucket", "uploads/photo.jpg", "/tmp/photo.jpg", "image/jpeg")

storage.UploadReader(ctx, "my-bucket", "uploads/photo.jpg", r.Body, r.ContentLength, "image/jpeg")

url, err := storage.GetURL(ctx, "my-bucket", "uploads/photo.jpg", 24*time.Hour)

storage.Delete(ctx, "my-bucket", "uploads/photo.jpg")

exists, err := storage.Exists(ctx, "my-bucket", "uploads/photo.jpg")
```

### Rate Limiter

```go
import "github.com/codelogydev/core-go/ratelimit"

r := gin.New()
r.Use(ratelimit.New(10, 20))

api := r.Group("/api")
api.Use(ratelimit.New(5, 10))
```

### Pagination

```go
import "github.com/codelogydev/core-go/pagination"

func GetUsers(c *gin.Context) {
    var params pagination.Params
    c.ShouldBindQuery(&params)
    params.Normalize()

    users, total, _ := service.GetUsers(params.Offset(), params.Limit)
    response.Success(c, pagination.NewResponse(users, total, params))
}
```

### Mailer

```go
import "github.com/codelogydev/core-go/mailer"

mailer.Init(mailer.Config{
    Host:     os.Getenv("MAIL_HOST"),
    Port:     587,
    Username: os.Getenv("MAIL_USERNAME"),
    Password: os.Getenv("MAIL_PASSWORD"),
    From:     os.Getenv("MAIL_FROM"),
})

mailer.Send("user@example.com", "OTP Code", "<h1>Your OTP: 123456</h1>")
```

## 📁 Project Structure

```
core-go/
├── auth/
│   ├── jwt.go
│   └── helper.go
├── cache/
│   └── redis.go
├── storage/
│   └── minio.go
├── mailer/
│   └── mailer.go
├── ratelimit/
│   └── ratelimit.go
├── pagination/
│   └── pagination.go
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
| `STORAGE_ENDPOINT` | S3/MinIO endpoint | `localhost:9000` |
| `STORAGE_ACCESS_KEY` | Access key | - |
| `STORAGE_SECRET_KEY` | Secret key | - |
| `STORAGE_USE_SSL` | Use HTTPS | `false` |
| `MAIL_HOST` | SMTP host | - |
| `MAIL_PORT` | SMTP port | `587` |
| `MAIL_USERNAME` | SMTP username | - |
| `MAIL_PASSWORD` | SMTP password | - |
| `MAIL_FROM` | Sender email address | - |
