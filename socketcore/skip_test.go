package socketcore_test

import (
	"testing"

	"github.com/epikur-io/gluasocket"
	lua "github.com/epikur-io/gopher-lua"
	"github.com/stretchr/testify/assert"
)

func TestSkip0(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString(`return require 'socket.core'.skip(0, 1, 2, 3)`))
	assert.Equal(3, L.GetTop())
	assert.Equal(1, L.ToInt(-3))
	assert.Equal(2, L.ToInt(-2))
	assert.Equal(3, L.ToInt(-1))
}

func TestSkip1(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString(`return require 'socket.core'.skip(1, 1, 2, 3)`))
	assert.Equal(2, L.GetTop())
	assert.Equal(2, L.ToInt(-2))
	assert.Equal(3, L.ToInt(-1))
}

func TestSkip2(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString(`return require 'socket.core'.skip(2, 1, 2, 3)`))
	assert.Equal(1, L.GetTop())
	assert.Equal(3, L.ToInt(-1))
}

func TestSkipPast(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	assert.NoError(L.DoString(`return require 'socket.core'.skip(9, 1, 2, 3)`))
	assert.Equal(0, L.GetTop())
}
