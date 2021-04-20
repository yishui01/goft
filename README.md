# goft (gin框架的强化版本)
-  在gin的基础上再进行封装，提高编写流畅度，提倡链式调用
```go
goft.Ignite().Attach(middlewares.NewUserMid()).
		Mount("v1", classes.NewUserClass()).
		Mount("v2", classes.NewIndexClass()).
		Launch()
```