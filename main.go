package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello World!")
	})
	err := engine.Run("0.0.0.0:8000")
	if err != nil {
		fmt.Println("Failed to run server")
		panic(err)
	}
}
