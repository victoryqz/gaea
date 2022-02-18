package gaea

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type G struct {
	*gin.Engine
	routergroup *gin.RouterGroup
}

func Ignite() *G {
	g := &G{Engine: gin.Default()}
	g.Use(gin.Recovery(), gin.Logger())
	return g
}

func (g *G) Attach(fairing Fairing) *G {
	g.Use(func(c *gin.Context) {
		err := fairing.OnRequest(c)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
		} else {
			c.Next()
		}
	})

	return g
}

func (g *G) Mount(group string, classes ...IClass) *G {
	g.routergroup = g.Group(group)
	for _, clazz := range classes {
		clazz.Register(g)
	}
	return g
}

func (g *G) Launch() {
	g.Run(":80")
}

func (g *G) Register(httpMethod, relativePath string, handler interface{}) {
	if f, ok := handler.(func(c *gin.Context) string); ok {
		g.routergroup.Handle(httpMethod, relativePath, func(c *gin.Context) {
			c.String(
				http.StatusOK,
				f(c),
			)
		})
		return
	}

	if f, ok := handler.(func(c *gin.Context) interface{}); ok {
		g.routergroup.Handle(httpMethod, relativePath, func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				f(c),
			)
		})
		return
	}

	if f, ok := handler.(func(c *gin.Context)); ok {
		g.routergroup.Handle(httpMethod, relativePath, func(c *gin.Context) {
			f(c)
		})
	}

}
