package user

import (
	"gin-demo/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(c *gin.Context) {
	var loginDto Login
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if loginDto.Username != "root" || loginDto.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	token := jwtAuth.GenerateToken(c, 1)
	c.JSON(http.StatusOK, gin.H{"code": "200", "accessToken": token})
}
