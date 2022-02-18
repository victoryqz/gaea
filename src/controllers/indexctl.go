package controllers

import (
	"gaea.olympus.io/src/gaea"
	"github.com/gin-gonic/gin"
)

type IndexCtl struct {
}

func NewIndex() *IndexCtl {
	return &IndexCtl{}
}

func (i *IndexCtl) ShowIndex01(c *gin.Context) string {
	return "hi, yqzhang"
}

func (i *IndexCtl) ShowIndex02(c *gin.Context) interface{} {
	return nil
}

/*
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
*/

func (i *IndexCtl) ShowIndex03(c *gin.Context) {

}

func (i *IndexCtl) Register(g *gaea.G) {
	//	g.Register(http.MethodGet, "/1", i.ShowIndex01)
	//	g.Register(http.MethodGet, "/2", i.ShowIndex02)
	//  g.Register(http.MethodGet, "/3", i.ShowIndex03)
}
