package mns

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func StartHttpServer(addr string) {
	log.Println("start http at:", addr)
	e := gin.Default()
	e.Use(gin.Recovery())
	e.HEAD("/", func(c *gin.Context) { c.AbortWithStatus(200) })
	e.GET("/", func(c *gin.Context) { c.String(200, "OK") })
	e.GET("/state", stateHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println("listen error:", err)
			os.Exit(-1)
		}
	}()
}
func stateHandler(g *gin.Context) {
	g.JSON(200, State)
}
