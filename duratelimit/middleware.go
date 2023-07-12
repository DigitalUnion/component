package duratelimit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RateLimit(l *DuRateLimit) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
		c.Next()
	}
}

func MultiRateLimit(key string, lMap map[string]*DuRateLimit) gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.Request.Header.Get(key)
		l, ok := lMap[cid]
		if ok && !l.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
		c.Next()
	}
}
