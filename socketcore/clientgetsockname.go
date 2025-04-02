package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientGetSockNameMethod(L *lua.LState) int {
	L.RaiseError("client:getsockname() not implemented yet")
	return 0
}
