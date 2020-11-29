package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.POST("/login", loginHandler)
	}
}
