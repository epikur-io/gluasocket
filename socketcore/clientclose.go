package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientCloseMethod(L *lua.LState) int {
	client := checkClient(L)
	if err := client.Conn.Close(); err != nil {
		L.RaiseError(err.Error())
	}
	return 0
}
