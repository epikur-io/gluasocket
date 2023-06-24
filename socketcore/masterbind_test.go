package socketcore_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
	"gitlab.com/megalithic-llc/gluasocket"
	"testing"
)

func TestMasterBind(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	script := fmt.Sprintf(`require 'socket'.bind('%s', '%s')`, "localhost", "8383")
	assert.NoError(L.DoString(script))
}
