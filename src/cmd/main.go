package main

import (
	"mygin/src/classes"
	"mygin/src/goft"
	"mygin/src/middlewares"
)

func main() {
	goft.Ignite().Attach(middlewares.NewUserMid()).
		Mount("v1", classes.NewUserClass()).
		Mount("v2", classes.NewIndexClass()).
		Launch()
}
