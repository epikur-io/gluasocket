package socketcore

import lua "github.com/epikur-io/gopher-lua"

func clientDirtyMethod(L *lua.LState) int {
	L.Push(lua.LBool(false))
	return 1
}
