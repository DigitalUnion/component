package duinflux

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type PathListType int

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if writeAPI == nil {
			return
		}
		path := c.Request.URL.Path
		if path == "/" || path == "/favicon.ico" {
			return
		}
		start := time.Now()
		c.Next()
		if pathsType == -1 {
			if inPathList(path) {
				return
			}
		} else if pathsType == 1 {
			if !inPathList(path) {
				return
			}
		}
		if len(path) > 12 {
			path = path[:12]
		}
		Done(fmt.Sprintf("gin_%d_%s_%s", c.Writer.Status(), c.Request.Method, path), start)
	}
}

func inPathList(path string) bool {
	if len(paths) == 0 {
		return false
	}
	for _, e := range paths {
		if e == path {
			return true
		}
	}
	return false
}
