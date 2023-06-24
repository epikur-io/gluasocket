package socketcore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
	"gitlab.com/megalithic-llc/gluasocket/socketcore"
	"testing"
	"time"
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
