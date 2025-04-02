package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientGetStatsMethod(L *lua.LState) int {
	L.RaiseError("client:getstats() not implemented yet")
	return 0
}
