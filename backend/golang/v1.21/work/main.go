package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	gin_zap "github.com/gin-contrib/zap"

	"idv/chris/component"
)

func main() {
	comps = component.New(component.DefaultConfig())

	g := gin.New()

	g.Use(MiddlewareLogger())
	g.Use(MiddlewareToken())
	g.Use(gin_zap.RecoveryWithZap(comps.Logger(), true)) // Recovery error
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// config.AllowHeaders = append(config.AllowHeaders, []string{"token"}...)
	g.Use(cors.New(config))

	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server v0.0.0")
	})

	router_api(g)

	s := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	go func() {
		comps.Logger().Debug("localhost:8080")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	q := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
	comps.Logger().Debug("Shutting down server...")

}
