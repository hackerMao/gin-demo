package routers

import "github.com/gin-gonic/gin"

type Option func(engine *gin.Engine)

var options []Option

// 注册子app路由
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	engine := gin.New()
	for _, opt := range options {
		opt(engine)
	}
	return engine
}
