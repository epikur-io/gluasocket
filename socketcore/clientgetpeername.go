package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientGetPeerNameMethod(L *lua.LState) int {
	L.RaiseError("client:getpeername() not implemented yet")
	return 0
}
