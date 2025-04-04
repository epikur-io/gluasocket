package socketcore

import (
	"net"

	lua "github.com/epikur-io/gopher-lua"
)

func clientGetFdMethod(L *lua.LState) int {
	client := checkClient(L)
	if file, err := client.Conn.(*net.TCPConn).File(); err != nil {
		L.RaiseError(err.Error())
		return 0
	} else {
		L.Push(lua.LNumber(file.Fd()))
		return 1
	}
}
