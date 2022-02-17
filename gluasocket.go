package gluasocket

import (
	"github.com/yuin/gopher-lua"
	"gluasocket/ltn12"
	"gluasocket/mime"
	"gluasocket/mimecore"
	"gluasocket/socket"
	"gluasocket/socketcore"
	"gluasocket/socketexcept"
	"gluasocket/socketftp"
	"gluasocket/socketheaders"
	"gluasocket/sockethttp"
	"gluasocket/socketsmtp"
	"gluasocket/sockettp"
	"gluasocket/socketurl"
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
