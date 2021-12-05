package serror

import "fmt"

var (
	Success           = &BaseError{0, "success"}
	ErrorParseRequest = &BaseError{400, "parse request error"}
)

type Error interface {
	Code() uint32
	Msg() string
	error
}
type BaseError struct {
	code uint32
	msg  string
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("msg=%s, code=%d", e.msg, e.code)
}
func (e *BaseError) Code() uint32 {
	return e.code
}
func (e *BaseError) Msg() string {
	return e.msg
}
func (e *BaseError) SetCode(code uint32) {
	e.code = code
}
func (e *BaseError) SetMsg(msg string) *BaseError {
	err := *e
	err.msg = msg
	return &err
}
func New(code uint32, format string, a ...interface{}) *BaseError {
	message := fmt.Sprintf(format, a...)
	return &BaseError{
		code: code,
		msg:  message,
	}
}
