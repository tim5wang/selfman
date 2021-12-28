package idgen

import (
	"testing"

	"github.com/tim5wang/selfman/common/util"
)

func TestSnowflakeID(t *testing.T) {
	n, _ := NewNode(1)
	id := n.Generate()
	util.Print(id.Base64())
	util.Print(id.String())
}
