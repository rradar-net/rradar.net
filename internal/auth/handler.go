package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rradar-net/rradar.net/internal/env"
	"github.com/rradar-net/rradar.net/internal/errs"
	"github.com/rradar-net/rradar.net/pkg/proto"
)

func LoginHandler(env env.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request proto.LoginRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errs.Format(err))
			return
		}

		user, err := loginUser(env, &request)

		if err != nil {
			c.JSON(err.HttpStatus, err.JSON())
			return
		}

		c.JSON(http.StatusCreated, &proto.RegisterResponse{
			Status: proto.Status_success,
			Data: &proto.RegisterResponse_Data{
				Username: user.Username,
			},
		})
	}
}

func RegisterHandler(env env.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request proto.RegisterRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errs.Format(err))
			return
		}

		user, err := registerUser(env, &request)
		if err != nil {
			c.JSON(err.HttpStatus, err.JSON())
			return
		}

		c.JSON(http.StatusCreated, &proto.RegisterResponse{
			Status: proto.Status_success,
			Data: &proto.RegisterResponse_Data{
				Username: user.Username,
			},
		})
	}
}
