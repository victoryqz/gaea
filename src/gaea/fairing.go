package gaea

import "github.com/gin-gonic/gin"

type Fairing interface {
	OnRequest(c *gin.Context) error
}
