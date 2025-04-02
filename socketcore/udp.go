package socketcore

import lua "github.com/epikur-io/gopher-lua"

func udpFn(L *lua.LState) int {
	L.RaiseError("socket.udp() not implemented yet")
	return 0
}
