package gluasocket_test

import (
	"os"
	"testing"

	"github.com/epikur-io/gluasocket"
	lua "github.com/epikur-io/gopher-lua"
	"github.com/stretchr/testify/assert"
)

func TestExcept(t *testing.T) {
	doTest("excepttest.lua", t)
}

// func TestFtp(t *testing.T) {
// 	doTest("ftptest.lua", t)
// }

// func TestHttp(t *testing.T) {
// 	doTest("httptest.lua", t)
// }

func TestLtn12(t *testing.T) {
	doTest("ltn12test.lua", t)
}

// func TestMime(t *testing.T) {
// 	doTest("mimetest.lua", t)
// }

func TestMimeDot(t *testing.T) {
	doTest("stufftest.lua", t)
}

// func TestSmtp(t *testing.T) {
// 	doTest("smtptest.lua", t)
// }

func TestSmtpMessage(t *testing.T) {
	doTest("testmesg.lua", t)
}

func TestSocketGetAddrInfo(t *testing.T) {
	doTest("test_getaddrinfo.lua", t)
}

//func TestSocketError(t *testing.T) {
//	doTest("test_socket_error.lua", t)
//}

func TestUrl(t *testing.T) {
	doTest("urltest.lua", t)
}

// ----------------------------------------------------------------------------
func doTest(testScript string, t *testing.T) {
	assert := assert.New(t)

	// Bring up a GopherLua VM
	L := lua.NewState()
	defer L.Close()

	// Register the Gluasocket module
	gluasocket.Preload(L)

	// Change working directory to where scripts are, so that nested scripts are found
	os.Chdir("testdata/luasocket-test")

	// Run Lua test script
	assert.NoError(L.DoFile(testScript))
}

// func TestAll(t *testing.T) {
// 	os.Chdir("testdata/luasocket-test")
// 	pages, err := filepath.Glob("./*.lua")
// 	if err != nil {
// 		t.Errorf("failed to get test files: %s", err)
// 		return
// 	}
// 	for _, p := range pages {
// 		if !strings.Contains(p, "test") {
// 			continue
// 		}
// 		t.Run(p, func(t *testing.T) {
// 			doTest(p, t)
// 		})
// 	}
// }
