package business

import (
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/tools/pkg/courier"
	"net/http"
)

type Meta struct {
	courier.MethodGet
}

func (m *Meta) Path() string {
	return "/meta"
}

func (m *Meta) Context(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
}
