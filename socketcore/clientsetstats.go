package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientSetStatsMethod(L *lua.LState) int {
	L.RaiseError("client:setstats() not implemented yet")
	return 0
}
