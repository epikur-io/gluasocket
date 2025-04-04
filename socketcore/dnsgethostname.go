package socketcore

import (
	"fmt"
	"os"

	lua "github.com/epikur-io/gopher-lua"
)

func dnsGetHostName(L *lua.LState) int {
	hostname, err := os.Hostname()
	if err != nil {
		L.RaiseError(fmt.Sprintf("Failure detecting hostname: %v", err))
		return 0
	}

	L.Push(lua.LString(hostname))
	return 1
}
