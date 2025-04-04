package socketcore

import (
	"time"

	lua "github.com/epikur-io/gopher-lua"
)

func masterSetTimeoutMethod(L *lua.LState) int {
	master, _ := checkMaster(L)
	timeout := L.CheckNumber(2)
	master.Timeout = time.Duration(timeout * 1.0e9)
	L.Push(lua.LNumber(1))
	return 1
}
