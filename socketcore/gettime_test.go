package socketcore_test

import (
	"testing"
	"time"

	"github.com/epikur-io/gluasocket/socketcore"
	lua "github.com/epikur-io/gopher-lua"
	"github.com/stretchr/testify/assert"
)

func TestGettime(t *testing.T) {
	assert := assert.New(t)

	luaState := lua.NewState()
	defer luaState.Close()

	luaState.PreloadModule("socket.core", socketcore.Loader)

	now := time.Now()
	assert.NoError(luaState.DoString("return require 'socket.core'.gettime()"))

	lv := luaState.Get(-1)
	retval, ok := lv.(lua.LNumber)

	assert.True(ok)
	expectedMin := float64(now.UnixNano()) / 1e9
	assert.True(float64(retval) >= expectedMin)
}
