package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientShutdownMethod(L *lua.LState) int {
	L.RaiseError("client:shutdown() not implemented yet")
	return 0
}
