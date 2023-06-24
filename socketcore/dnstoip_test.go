package socketcore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
	"gitlab.com/megalithic-llc/gluasocket"
	"testing"
)

func TestDnsToIpLocalhost(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString("return require 'socket.core'.dns.toip('localhost')"))
	assert.Equal("127.0.0.1", L.ToString(-1))
}
