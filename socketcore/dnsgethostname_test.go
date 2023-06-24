package socketcore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
	"gitlab.com/megalithic-llc/gluasocket/socketcore"
	"os"
	"testing"
)

func TestDnsGetHostName(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("socket.core", socketcore.Loader)

	expected, err := os.Hostname()
	assert.NoError(err)

	assert.NoError(L.DoString(`return require 'socket.core'.dns.gethostname()`))
	actual := L.Get(-1)

	assert.Equal(expected, actual.String())
}
