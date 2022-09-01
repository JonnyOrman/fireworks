package fireworks

import "github.com/gin-gonic/gin"

type GinRouterBuilder struct {
	router *gin.Engine
}

func NewGinRouterBuilder() *GinRouterBuilder {
	this := new(GinRouterBuilder)

	this.router = gin.Default()

	return this
}

func (this GinRouterBuilder) AddGet(routePath string, handler gin.HandlerFunc) {
	this.router.GET(routePath, handler)
}

func (this GinRouterBuilder) AddPost(routePath string, handler gin.HandlerFunc) {
	this.router.POST(routePath, handler)
}

func (this GinRouterBuilder) Build() *gin.Engine {
	return this.router
}
