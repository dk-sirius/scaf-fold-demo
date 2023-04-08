package controller

import (
	"github.com/dk-sirius/scaf-fold-demo/client/controller/business"
	"github.com/scaf-fold/tools/pkg/courier"
)

var Root = courier.NewGroup("/srv-demo")

func init() {
	V1 := Root.Group("/v1")
	{
		V1.Register(business.Root)
	}
}
