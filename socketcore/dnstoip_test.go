package socketcore_test

import (
	"testing"

	"github.com/epikur-io/gluasocket"
	lua "github.com/epikur-io/gopher-lua"
	"github.com/stretchr/testify/assert"
)

func TestDnsToIpLocalhost(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString("return require 'socket.core'.dns.toip('localhost')"))
	assert.Equal("127.0.0.1", L.ToString(-1))
}
