package mimecore

import lua "github.com/epikur-io/gopher-lua"

// ----------------------------------------------------------------------------

var exports = map[string]lua.LGFunction{
	"b64":   b64Fn,
	"dot":   dotFn,
	"eol":   eolFn,
	"qp":    qpFn,
	"qpwrp": qpwrpFn,
	"unb64": unb64Fn,
	"unqp":  unqpFn,
	"wrp":   wrpFn,
}

// ----------------------------------------------------------------------------

func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)

	qpsetup()
	b64setup()

	return 1
}
