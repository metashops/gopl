package example7_6

import (
	`encoding/json`
	`net/http`
	`net/http/httptest`
	`testing`
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux() // 创建一个用于运行测试的多路复用
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}
