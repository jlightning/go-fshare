package fshare

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func getLoginServer() *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if _, err := res.Write([]byte(`{"code":200,"msg":"Login successfully!","token":"f2e287e308193fe02f9be2da8ff260ba1b7abf52","session_id":"e9ot9rpbdlcjim48s4q3c01u4s"}`)); err != nil {
			panic(err)
		}
	}))
	return testServer
}

// FAIL: {"code":405,"msg":"Authenticate fail!"}
// PASS: {"code":200,"msg":"Login successfully!","token":"f2e287e308193fe02f9be2da8ff260ba1b7abf52","session_id":"e9ot9rpbdlcjim48s4q3c01u4s"}
func TestClient_Login(t *testing.T) {
	testServer := getLoginServer()

	client := NewClient(&Config{
		Username:   PointerString("test@test.com"),
		Password:   PointerString("123456"),
		HttpClient: http.DefaultClient,
		loginUrl:   PointerString(testServer.URL),
	})
	if err := client.Login(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if client.token != "f2e287e308193fe02f9be2da8ff260ba1b7abf52" {
		t.Errorf("client.token must be: %s", "f2e287e308193fe02f9be2da8ff260ba1b7abf52")
	}

	if client.sessionID != "e9ot9rpbdlcjim48s4q3c01u4s" {
		t.Errorf("client.sessionID must be: %s", "e9ot9rpbdlcjim48s4q3c01u4s")
	}
}

// FAIL: {"code":404,"msg":"T\u1eadp tin kh\u00f4ng t\u1ed3n t\u1ea1i"}
// PASS: {"location":"http://test.com/test-file.zip"}
func TestClient_GetDownloadURL(t *testing.T) {
	loginTestServer := getLoginServer()

	downloadTestServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if _, err := res.Write([]byte(`{"location":"http://test.com/test-file.zip"}`)); err != nil {
			panic(err)
		}
	}))

	client := NewClient(&Config{
		Username:    PointerString("test@test.com"),
		Password:    PointerString("123456"),
		HttpClient:  http.DefaultClient,
		loginUrl:    PointerString(loginTestServer.URL),
		downloadUrl: PointerString(downloadTestServer.URL),
	})
	if err := client.Login(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	url, err := client.GetDownloadURL("https://www.fshare.vn/file/8CQFOSNLT3LJ")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if url != "http://test.com/test-file.zip" {
		t.Errorf("url must be: %s", "http://test.com/test-file.zip")
	}
}
