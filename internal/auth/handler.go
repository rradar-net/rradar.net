package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rradar-net/rradar.net/internal/verrors"
	"github.com/rradar-net/rradar.net/pkg/proto"
)

func LoginHandler(c *gin.Context) {

}

func RegisterHandler(c *gin.Context) {
	var request proto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": verrors.Format(err)})
		return
	}

	c.JSON(http.StatusOK, &request)
}
