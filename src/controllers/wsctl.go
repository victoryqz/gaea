package controllers

import (
	"log"
	"net/http"
	"sync"

	"gaea.olympus.io/src/core"
	"gaea.olympus.io/src/gaea"
	"github.com/gin-gonic/gin"
)

type WsCtl struct {
}

var wsCtlInstance *WsCtl
var wsCtlOnce sync.Once

func GetWsCtl() *WsCtl {
	wsCtlOnce.Do(func() {
		wsCtlInstance = &WsCtl{}
	})

	return wsCtlInstance
}

func (wc *WsCtl) WsRegister(c *gin.Context) {

	conn, err := core.NewWsConnection(c)

	if err != nil {
		log.Println(err)
		return
	}

	core.GetWsManage().
		Join(conn, func(wm *core.WsMessage) interface{} {
			return wm
		}).
		Subscribe(
			"rook-ceph",
			"deployment",
			conn.GetID(),
		)
}

func (wc *WsCtl) Register(g *gaea.G) {
	g.Register(http.MethodGet, "/ws/register", wc.WsRegister)
}
