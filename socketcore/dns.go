package socketcore

import lua "github.com/epikur-io/gopher-lua"

type DNS struct {
}

var dnsMethods = map[string]lua.LGFunction{
	"getaddrinfo": dnsGetAddrInfo,
	"gethostname": dnsGetHostName,
	"tohostname":  dnsToHostName,
	"toip":        dnsToIp,
}
