package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "hello world"})

}
