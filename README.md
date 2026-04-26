# core-go

Reusable core library for Go backend services.

## 🚀 Features

- JWT Authentication (generate & validate)
- Auth Middleware (Gin)
- Standard API Response (success & error)
- Structured Logger (Zap)
- Environment Config Helper

---

## 📦 Installation

```bash
go get github.com/codelogydev/core-go@v1.0.0
```

## 🛠 Usage

### Initialize Logger

```go
import (
	"github.com/codelogydev/core-go/logger"
)

func main() {
	logger.Init()
	defer logger.Log.Sync()

	logger.Log.Info("logger initialized")
}
```

### JWT Generate & Verify

```go
import (
	"github.com/codelogydev/core-go/auth"
)

// Generate token
token, err := auth.GenerateToken(1)

// Verify token
claims, err := auth.VerifyToken(token)
```

### Auth Middleware

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/codelogydev/core-go/auth"
	"github.com/codelogydev/core-go/response"
)

func ProtectedRoute(c *gin.Context) {
	userID := auth.GetUserID(c)
	if userID == 0 {
		response.Unauthorized(c, "invalid token")
		return
	}
	// ...
}
```

### Standard Response

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/codelogydev/core-go/response"
)

// Success
response.Success(c, gin.H{"message": "ok"}, 200)

// Error
response.Error(c, 500, "database error")
response.BadRequest(c, "invalid request")
```
