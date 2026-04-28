package ratelimit

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type entry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type errBody struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

var (
	mu      sync.Mutex
	clients = make(map[string]*entry)
)

func init() {
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, e := range clients {
				if time.Since(e.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()
}

func get(ip string, r rate.Limit, burst int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if e, ok := clients[ip]; ok {
		e.lastSeen = time.Now()
		return e.limiter
	}

	l := rate.NewLimiter(r, burst)
	clients[ip] = &entry{limiter: l, lastSeen: time.Now()}
	return l
}

func New(rps float64, burst int) gin.HandlerFunc {
	r := rate.Limit(rps)
	return func(c *gin.Context) {
		if !get(c.ClientIP(), r, burst).Allow() {
			c.JSON(http.StatusTooManyRequests, errBody{Success: false, Error: "too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
