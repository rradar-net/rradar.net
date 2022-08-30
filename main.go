package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rradar-net/rradar.net/internal/auth"
	"github.com/rradar-net/rradar.net/internal/verrors"
)

func main() {
	router := gin.Default()

	verrors.Register()

	rAuth := router.Group("/auth")
	{
		rAuth.POST("/login", auth.LoginHandler)
		rAuth.POST("/register", auth.RegisterHandler)
	}

	router.Run()
}
