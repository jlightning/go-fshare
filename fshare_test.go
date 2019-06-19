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

// PASS: [{"id":"1017114","linkcode":"T9W1XXFV0T","name":"Tap 83 84.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2324","size":"145626234","downloadcount":"352","deleted":"0","mimetype":"application\/x-rar","created":"1287782157","modified2":"1287782157","modified":"1287782157","file_type":"1","pwd":null,"crc32":"1812127515","folder_path":"201010\/23","storage_id":"18","realname":"b357a0043a82a8c98e9b8d353a533a290ee841ae_tap-83-84.rar","lastdownload":"1550380704"},{"id":"1017128","linkcode":"TH3FM1ZBBT","name":"Tap 85 86.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2325","size":"145640625","downloadcount":"262","deleted":"0","mimetype":"application\/x-rar","created":"1287784927","modified2":"1287784927","modified":"1287784927","file_type":"1","pwd":null,"crc32":"3335621943","folder_path":"201010\/23","storage_id":"23","realname":"134a5b777c443ec436a577cd4e4ad50990b7cf4e_tap-85-86.rar","lastdownload":"1550380746"},{"id":"1017151","linkcode":"TCVHF15BST","name":"Tap 87 88.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2326","size":"132638738","downloadcount":"275","deleted":"0","mimetype":"application\/x-rar","created":"1287787452","modified2":"1287787452","modified":"1287787452","file_type":"1","pwd":null,"crc32":"859100688","folder_path":"201010\/23","storage_id":"24","realname":"9c97a71bf507c21d01e53e6f27affba8c3cf2755_87_88.rar","lastdownload":"1550380784"},{"id":"1017174","linkcode":"T0Y8RV66GT","name":"Tap 89 90.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2327","size":"133767377","downloadcount":"358","deleted":"0","mimetype":"application\/x-rar","created":"1287793872","modified2":"1287793872","modified":"1287793872","file_type":"1","pwd":null,"crc32":"4063582691","folder_path":"201010\/23","storage_id":"17","realname":"e11e40af30cf7c9356dafde8492f1455cdff7eb2_tap-89-90.rar","lastdownload":"1550380827"},{"id":"976117","linkcode":"T2AC89FQPT","name":"Tap 9 10.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2106","size":"159215805","downloadcount":"234","deleted":"0","mimetype":"application\/x-rar","created":"1285045770","modified2":"1285045770","modified":"1285045770","file_type":"1","pwd":null,"crc32":"1388354608","folder_path":"201009\/21","storage_id":"23","realname":"217364882674649279b3de68ac48b8d54c845959_tap-9-10.rar","lastdownload":"1550380873"},{"id":"1017198","linkcode":"TQFNK2X75T","name":"Tap 91 92.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2328","size":"142358546","downloadcount":"272","deleted":"0","mimetype":"application\/x-rar","created":"1287796647","modified2":"1287796647","modified":"1287796647","file_type":"1","pwd":null,"crc32":"3538464102","folder_path":"201010\/23","storage_id":"20","realname":"191ad0a7c7650a64f79a4d8e3c2e172012da5110_tap-91-92.rar","lastdownload":"1550380922"},{"id":"1020645","linkcode":"THFZ72W9KT","name":"Tap 93 94.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2332","size":"135898203","downloadcount":"365","deleted":"0","mimetype":"application\/x-rar","created":"1287976304","modified2":"1287976304","modified":"1287976304","file_type":"1","pwd":null,"crc32":"1250507545","folder_path":"201010\/25","storage_id":"24","realname":"f756b2b92557a2a25dc95faf3f772aeabee40ad9_tap-93-94.rar","lastdownload":"1550380961"},{"id":"1020731","linkcode":"TFASJBP16T","name":"Tap 95 96.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2333","size":"140219481","downloadcount":"407","deleted":"0","mimetype":"application\/x-rar","created":"1287979227","modified2":"1287979227","modified":"1287979227","file_type":"1","pwd":null,"crc32":"1344708659","folder_path":"201010\/25","storage_id":"18","realname":"1f94e0db8ae0c4c574bc544798db854bf38212a9_tap-95-96.rar","lastdownload":"1550381001"},{"id":"1017336","linkcode":"TX2VJ1SX4T","name":"Tap 97 98.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2329","size":"130544912","downloadcount":"346","deleted":"0","mimetype":"application\/x-rar","created":"1287804008","modified2":"1287804008","modified":"1287804008","file_type":"1","pwd":null,"crc32":"4218091492","folder_path":"201010\/23","storage_id":"17","realname":"2c8fe238837d77f6aefbea952144cd9f959da190_tap-97-98.rar","lastdownload":"1550381038"},{"id":"1020576","linkcode":"TTTMMV9MTT","name":"Tap 99 100.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2331","size":"138212640","downloadcount":"479","deleted":"0","mimetype":"application\/x-rar","created":"1287973643","modified2":"1287973643","modified":"1287973643","file_type":"1","pwd":null,"crc32":"4128546228","folder_path":"201010\/25","storage_id":"19","realname":"cb0f0a159eaf4a92a1a38ee9b1ee199415aaafee_tap-99-100.rar","lastdownload":"1550381079"}]
func TestClient_GetFolderURLs(t *testing.T) {
	loginTestServer := getLoginServer()
	folderListTestServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if _, err := res.Write([]byte(`[{"id":"1017114","linkcode":"T9W1XXFV0T","name":"Tap 83 84.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2324","size":"145626234","downloadcount":"352","deleted":"0","mimetype":"application\/x-rar","created":"1287782157","modified2":"1287782157","modified":"1287782157","file_type":"1","pwd":null,"crc32":"1812127515","folder_path":"201010\/23","storage_id":"18","realname":"b357a0043a82a8c98e9b8d353a533a290ee841ae_tap-83-84.rar","lastdownload":"1550380704"},{"id":"1017128","linkcode":"TH3FM1ZBBT","name":"Tap 85 86.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2325","size":"145640625","downloadcount":"262","deleted":"0","mimetype":"application\/x-rar","created":"1287784927","modified2":"1287784927","modified":"1287784927","file_type":"1","pwd":null,"crc32":"3335621943","folder_path":"201010\/23","storage_id":"23","realname":"134a5b777c443ec436a577cd4e4ad50990b7cf4e_tap-85-86.rar","lastdownload":"1550380746"},{"id":"1017151","linkcode":"TCVHF15BST","name":"Tap 87 88.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2326","size":"132638738","downloadcount":"275","deleted":"0","mimetype":"application\/x-rar","created":"1287787452","modified2":"1287787452","modified":"1287787452","file_type":"1","pwd":null,"crc32":"859100688","folder_path":"201010\/23","storage_id":"24","realname":"9c97a71bf507c21d01e53e6f27affba8c3cf2755_87_88.rar","lastdownload":"1550380784"},{"id":"1017174","linkcode":"T0Y8RV66GT","name":"Tap 89 90.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2327","size":"133767377","downloadcount":"358","deleted":"0","mimetype":"application\/x-rar","created":"1287793872","modified2":"1287793872","modified":"1287793872","file_type":"1","pwd":null,"crc32":"4063582691","folder_path":"201010\/23","storage_id":"17","realname":"e11e40af30cf7c9356dafde8492f1455cdff7eb2_tap-89-90.rar","lastdownload":"1550380827"},{"id":"976117","linkcode":"T2AC89FQPT","name":"Tap 9 10.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2106","size":"159215805","downloadcount":"234","deleted":"0","mimetype":"application\/x-rar","created":"1285045770","modified2":"1285045770","modified":"1285045770","file_type":"1","pwd":null,"crc32":"1388354608","folder_path":"201009\/21","storage_id":"23","realname":"217364882674649279b3de68ac48b8d54c845959_tap-9-10.rar","lastdownload":"1550380873"},{"id":"1017198","linkcode":"TQFNK2X75T","name":"Tap 91 92.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2328","size":"142358546","downloadcount":"272","deleted":"0","mimetype":"application\/x-rar","created":"1287796647","modified2":"1287796647","modified":"1287796647","file_type":"1","pwd":null,"crc32":"3538464102","folder_path":"201010\/23","storage_id":"20","realname":"191ad0a7c7650a64f79a4d8e3c2e172012da5110_tap-91-92.rar","lastdownload":"1550380922"},{"id":"1020645","linkcode":"THFZ72W9KT","name":"Tap 93 94.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2332","size":"135898203","downloadcount":"365","deleted":"0","mimetype":"application\/x-rar","created":"1287976304","modified2":"1287976304","modified":"1287976304","file_type":"1","pwd":null,"crc32":"1250507545","folder_path":"201010\/25","storage_id":"24","realname":"f756b2b92557a2a25dc95faf3f772aeabee40ad9_tap-93-94.rar","lastdownload":"1550380961"},{"id":"1020731","linkcode":"TFASJBP16T","name":"Tap 95 96.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2333","size":"140219481","downloadcount":"407","deleted":"0","mimetype":"application\/x-rar","created":"1287979227","modified2":"1287979227","modified":"1287979227","file_type":"1","pwd":null,"crc32":"1344708659","folder_path":"201010\/25","storage_id":"18","realname":"1f94e0db8ae0c4c574bc544798db854bf38212a9_tap-95-96.rar","lastdownload":"1550381001"},{"id":"1017336","linkcode":"TX2VJ1SX4T","name":"Tap 97 98.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2329","size":"130544912","downloadcount":"346","deleted":"0","mimetype":"application\/x-rar","created":"1287804008","modified2":"1287804008","modified":"1287804008","file_type":"1","pwd":null,"crc32":"4218091492","folder_path":"201010\/23","storage_id":"17","realname":"2c8fe238837d77f6aefbea952144cd9f959da190_tap-97-98.rar","lastdownload":"1550381038"},{"id":"1020576","linkcode":"TTTMMV9MTT","name":"Tap 99 100.rar","secure":"1","public":"1","shared":"0","directlink":"0","type":"1","path":"\/Tom and Jerry","owner_id":"90082","pid":"2331","size":"138212640","downloadcount":"479","deleted":"0","mimetype":"application\/x-rar","created":"1287973643","modified2":"1287973643","modified":"1287973643","file_type":"1","pwd":null,"crc32":"4128546228","folder_path":"201010\/25","storage_id":"19","realname":"cb0f0a159eaf4a92a1a38ee9b1ee199415aaafee_tap-99-100.rar","lastdownload":"1550381079"}]`)); err != nil {
			panic(err)
		}
	}))

	client := NewClient(&Config{
		Username:      PointerString("test@test.com"),
		Password:      PointerString("123456"),
		HttpClient:    http.DefaultClient,
		loginUrl:      PointerString(loginTestServer.URL),
		folderListUrl: PointerString(folderListTestServer.URL),
	})
	if err := client.Login(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	res, err := client.GetFolderURLs("https://www.fshare.vn/folder/TPSSYSPYFT", 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(res) != 10 {
		t.Errorf("len should be %d", 10)
	}
}
