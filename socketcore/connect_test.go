package socketcore_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/epikur-io/gluasocket"
	lua "github.com/epikur-io/gopher-lua"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	listener, err := net.Listen("tcp", "localhost:0")
	assert.NoError(err)
	port := listener.Addr().(*net.TCPAddr).Port

	accepted := false
	go func() {
		if _, err := listener.Accept(); err == nil {
			accepted = true
		}
	}()

	script := fmt.Sprintf(`c=require 'socket.core'.connect('%s', %d); c:close()`, "127.0.0.1", port)
	assert.NoError(L.DoString(script))

	time.Sleep(20 * time.Millisecond)
	assert.True(accepted)

	assert.NoError(listener.Close())
}
