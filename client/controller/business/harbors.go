package business

import (
	"github.com/dk-sirius/scaf-fold-demo/client/pkg/business"
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/tools/pkg/courier"
	"net/http"
)

type Harbors struct {
	courier.MethodGet
}

func (m *Harbors) Path() string {
	return "/harbors"
}

func (m *Harbors) Context(ctx *gin.Context) {
	config := business.NewConfigClient()
	list, err := config.ListHarbors()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, list)
}
