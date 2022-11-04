package example7_6

import (
	`encoding/json`
	`net/http`
	`path`
	`strconv`
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePOST(w, r)
	case "PUT":
		err = handlePUT(w, r)
	case "DELETE":
		err = handleDELETE(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// 从CRL路径提取字符串格式的帖子ID
	// strconv.Atoi 是把这个ID转换为整数格式
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id) // 从数据库获取数据,并将其填充到Post结构中
	if err != nil {
		return
	}
	// 把Post结构体封装为JSON格式的字节切片,
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	// 把 JSON 数据写入 ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func handlePOST(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len) // 创建一个字节切片,用于存放请求体
	r.Body.Read(body)         // 读取到的请求体,并将存在在字节切片中
	var post Post
	json.Unmarshal(body, &post) // 将body切片存解封至post结构体
	err = post.create()         // 创建数据库记录
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePUT(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body) // 从请求主体读取JSON数据
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return nil
}

func handleDELETE(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
