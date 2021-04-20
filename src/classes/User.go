package classes

import (
	"github.com/gin-gonic/gin"
	"mygin/src/goft"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "user list", // 换行别忘了 逗号
		})
	}
}

func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.UserList())
}
