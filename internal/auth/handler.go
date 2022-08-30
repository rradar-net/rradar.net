package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rradar-net/rradar.net/internal/verrors"
)

func LoginHandler(c *gin.Context) {

}

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=256"`
	Email    string `json:"email"`
}

func RegisterHandler(c *gin.Context) {
	var request registerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": verrors.Format(err)})
		return
	}

	c.JSON(http.StatusOK, request)
}
