package socketcore

import lua "github.com/epikur-io/gopher-lua"

func dnsToHostName(l *lua.LState) int {
	l.RaiseError("socket.dns.tohostname(address) not implemented yet")
	return 0
}
