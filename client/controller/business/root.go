package business

import "github.com/scaf-fold/tools/pkg/courier"

var Root = courier.NewGroup("/business")

func init() {
	Root.Register(&Meta{})
	Root.Register(&Harbors{})
}
