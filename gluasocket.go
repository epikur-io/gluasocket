package gluasocket

import (
	"github.com/yuin/gopher-lua"
	"gitlab.com/megalithic-llc/gluasocket/ltn12"
	"gitlab.com/megalithic-llc/gluasocket/mime"
	"gitlab.com/megalithic-llc/gluasocket/mimecore"
	"gitlab.com/megalithic-llc/gluasocket/socket"
	"gitlab.com/megalithic-llc/gluasocket/socketcore"
	"gitlab.com/megalithic-llc/gluasocket/socketexcept"
	"gitlab.com/megalithic-llc/gluasocket/socketftp"
	"gitlab.com/megalithic-llc/gluasocket/socketheaders"
	"gitlab.com/megalithic-llc/gluasocket/sockethttp"
	"gitlab.com/megalithic-llc/gluasocket/socketsmtp"
	"gitlab.com/megalithic-llc/gluasocket/sockettp"
	"gitlab.com/megalithic-llc/gluasocket/socketurl"
)

func Preload(L *lua.LState) {
	L.PreloadModule("ltn12", ltn12.Loader)
	L.PreloadModule("mime.core", mimecore.Loader)
	L.PreloadModule("mime", mime.Loader)
	L.PreloadModule("socket", socket.Loader)
	L.PreloadModule("socket.core", socketcore.Loader)
	L.PreloadModule("socket.except", socketexcept.Loader)
	L.PreloadModule("socket.ftp", socketftp.Loader)
	L.PreloadModule("socket.headers", socketheaders.Loader)
	L.PreloadModule("socket.http", sockethttp.Loader)
	L.PreloadModule("socket.smtp", socketsmtp.Loader)
	L.PreloadModule("socket.tp", sockettp.Loader)
	L.PreloadModule("socket.url", socketurl.Loader)
}
