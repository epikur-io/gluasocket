package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientSetOptionMethod(L *lua.LState) int {
	L.RaiseError("client:setoption() not implemented yet")
	return 0
}
