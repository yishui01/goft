package goft

import (
	"github.com/gin-gonic/gin"
)

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Ignite() *Goft {
	return &Goft{Engine: gin.New()}
}

// 启动
func (this *Goft) Launch() {
	this.Run(":8080")
}

// 加入中间件
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		if err := f.OnRequest(); err != nil {
			context.AbortWithStatusJSON(400, gin.H{"err": err.Error()})
		} else {
			context.Next()
		}
	})
	return this
}

// 加入路由、路由组
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}

// 路由配合group
func (this *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Goft {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}
