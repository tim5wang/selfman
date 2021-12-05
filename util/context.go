package util

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var contextPool = sync.Pool{
	New: func() interface{} {
		return &Context{}
	},
}

type Context struct {
	ctx *gin.Context
}
