package main

import (
	"fmt"

	gcLogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rradar-net/rradar.net/internal/auth"
	"github.com/rradar-net/rradar.net/internal/env"
	"github.com/rradar-net/rradar.net/internal/errors"
	"github.com/rradar-net/rradar.net/internal/logger"
)

func main() {
	env := env.Init()
	errors.Init()

	router := gin.New()
	loggerMw := gcLogger.SetLogger(gcLogger.WithLogger(logger.Init()))
	router.Use(loggerMw, gin.Recovery())

	v1 := router.Group("/v1")
	{
		rAuth := v1.Group("/auth")
		{
			rAuth.POST("/login", auth.LoginHandler(env))
			rAuth.POST("/register", auth.RegisterHandler(env))
		}
	}

	router.Run(fmt.Sprintf(":%d", env.Cfg.Port))
}
