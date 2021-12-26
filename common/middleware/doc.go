package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/util"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

var APIDoc = func(c *gin.Context) {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	util.Print(path, raw)
	blw := &BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	c.Next()

	responseByte := blw.Body.Bytes()
	responseString := string(responseByte)
	retCode := c.Writer.Status()
	util.Print(retCode, responseString)
}
