package sockethttp_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
	"gitlab.com/megalithic-llc/gluasocket/sockethttp"
	"net"
	"net/http"
	"testing"
)

func TestWebHdfsUrlScheme(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()

	// Start an HTTP endpoint for LuaSocket to connect to
	listener, err := net.Listen("tcp", "localhost:0")
	assert.NoError(err)
	port := listener.Addr().(*net.TCPAddr).Port
	go func() {
		http.Serve(listener, nil)
	}()

	// Make a webdav request
	L.PreloadModule("socket.http", sockethttp.Loader)
	url := fmt.Sprintf("webhdfs://localhost:%d", port)
	assert.NoError(L.DoString(fmt.Sprintf(`require 'socket.http'.request('%s')`, url)))
}
