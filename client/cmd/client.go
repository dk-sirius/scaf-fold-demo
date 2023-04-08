package main

import (
	"github.com/dk-sirius/scaf-fold-demo/client/controller"
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/tools/pkg/courier"
)

func main() {
	e := gin.Default()
	courier.App(e, controller.Root, func(e *gin.Engine) {
		e.Run(":8081")
	})
}
