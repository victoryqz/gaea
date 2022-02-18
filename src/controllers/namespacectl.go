package controllers

import (
	"net/http"
	"sync"

	"gaea.olympus.io/src/gaea"
	"gaea.olympus.io/src/services"
	"github.com/gin-gonic/gin"
)

type NamespaceCtl struct {
	service services.NamespaceService
}

var namespaceCtlInstance *NamespaceCtl
var namespaceCtlOnce sync.Once

func GetNamespaceCtl() *NamespaceCtl {
	namespaceCtlOnce.Do(func() {
		namespaceCtlInstance = &NamespaceCtl{service: *services.GetNamespaceService()}
	})

	return namespaceCtlInstance
}

func (nc *NamespaceCtl) ShowAllNamespace(c *gin.Context) interface{} {
	return nc.service.ShowAllNamespace()
}

func (nc *NamespaceCtl) Register(g *gaea.G) {
	g.Register(http.MethodGet, "/namespaces", nc.ShowAllNamespace)
}
