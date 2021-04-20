package middlewares

import "fmt"

type UserMid struct {
}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (this *UserMid) OnRequest() error {
	fmt.Println("用户中间件")
	return nil
}
