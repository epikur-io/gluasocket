package socketcore_test

import (
	"testing"

	"github.com/epikur-io/gluasocket"
	lua "github.com/epikur-io/gopher-lua"
	"github.com/stretchr/testify/assert"
)

func TestTcp(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString(`return require 'socket.core'.tcp()`))
	master := L.Get(-1)
	assert.Equal(lua.LTUserData, master.Type())
}
