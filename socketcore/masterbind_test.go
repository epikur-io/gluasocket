package socketcore_test

import (
	"fmt"
	"testing"

	"github.com/epikur-io/gluasocket"
	lua "github.com/epikur-io/gopher-lua"
	"github.com/stretchr/testify/assert"
)

func TestMasterBind(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	script := fmt.Sprintf(`require 'socket'.bind('%s', '%s')`, "localhost", "8383")
	assert.NoError(L.DoString(script))
}
