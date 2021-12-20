package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/serror"
)

const (
	StatusOK    uint32 = 0
	StatusError uint32 = 400
)

type Response struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, data ...interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, &Response{
		Code: StatusOK,
		Msg:  "",
		Data: data,
	})
}

func Error(ctx *gin.Context, err error) {
	code := StatusError
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	if se, ok := err.(serror.Error); ok {
		code = se.Code()
	}
	ctx.AbortWithStatusJSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
	})
}

func GeneralResponse(ctx *gin.Context, err error, data ...interface{}) {
	if err != nil {
		Error(ctx, err)
		return
	}
	Success(ctx, data)
}
