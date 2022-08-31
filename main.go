package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rradar-net/rradar.net/internal/auth"
	"github.com/rradar-net/rradar.net/internal/env"
	"github.com/rradar-net/rradar.net/internal/verrors"
)

func main() {
	router := gin.Default()
	verrors.Register()
	env := env.Init()

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
