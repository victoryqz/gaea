package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

type RBACMid struct {
}

func NewRBACMid() *RBACMid {
	return &RBACMid{}
}

func (r *RBACMid) OnRequest(c *gin.Context) error {

	log.Println("rbac middleware")

	return nil
}
