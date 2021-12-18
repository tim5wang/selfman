package controller

import (
	"github.com/tim5wang/selfman/common/web"
)

type DocModule struct{}

func NewDocModule() web.Module {
	return &DocModule{}
}

func (m *DocModule) Init(r web.Router) {
	r.Group("api/doc")
}
