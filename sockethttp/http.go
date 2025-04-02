package sockethttp

import lua "github.com/epikur-io/gopher-lua"

// ----------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	"request": requestFn,
}

// ----------------------------------------------------------------------------

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}
