package main

import (
	"fmt"
	"gin-demo/app/blog"
	"gin-demo/app/shop"
	"gin-demo/app/user"
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 加载多个app路由
	routers.Include(shop.Routers, blog.Routers, user.Routers)
	// 初始化路由
	engine := routers.Init()

	//engine := gin.Default()

	//engine.GET("/", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "hello World!")
	//})
	// 获取api参数
	engine.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		action = strings.Trim(action, "/")
		ctx.String(http.StatusOK, name+" is "+action)
	})
	// 查询字符串
	engine.GET("/user", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "stranger")
		ctx.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	engine.POST("/commit", func(c *gin.Context) {
		err := c.Bind(&user.Login{})
		if err != nil {
			fmt.Println("bind data error, err: ", err.Error())
			c.String(http.StatusInternalServerError, "bind data error")
		}
	})
	// form表单参数
	//engine.POST("/login", func(ctx *gin.Context) {
	//	types := ctx.DefaultPostForm("type", "post")
	//	username := ctx.PostForm("username")
	//	password := ctx.PostForm("password")
	//	ctx.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s, type:%s", username, password, types))
	//})
	// 上传单个文件
	//engine.MaxMultipartMemory = 8 << 20 // 限制最大上传尺寸
	//fmt.Println(engine.MaxMultipartMemory)
	//engine.POST("/upload", func(ctx *gin.Context) {
	//	file, err := ctx.FormFile("test")
	//	if err != nil {
	//		ctx.String(500, "上传图片失败")
	//	}
	//	err = ctx.SaveUploadedFile(file, file.Filename)
	//	if err != nil {
	//		ctx.String(500, "保存文件失败")
	//	}
	//	ctx.String(http.StatusOK, file.Filename)
	//})

	// 分组路由
	//v1 := engine.Group("/v1")
	//{
	//	v1.GET("/login", login)
	//	v1.GET("/submit", submit)
	//}
	//v2 := engine.Group("/v2")
	//{
	//	v2.POST("/login", login)
	//	v2.POST("/submit", submit)
	//}

	err := engine.Run("0.0.0.0:8000")
	if err != nil {
		fmt.Println("Failed to run server")
		panic(err)
	}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "stranger")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}
