package main

import (
	`encoding/json`
	`fmt`
	`io`
	`os`
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

func NewPost() *Post {
	return &Post{}
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// func main() {
// 	// 1、open file
// 	jsonFile, err := os.Open("./gowp/ch7/json/post.json")
// 	if err != nil {
// 		fmt.Printf("Error opening", err)
// 		return
// 	}
// 	defer jsonFile.Close()
//
// 	// 2、Read jsonFile
// 	jsonData, err := io.ReadAll(jsonFile)
// 	if err != nil {
// 		fmt.Println("Failed to read error:", err)
// 		return
// 	}
// 	// 3、will read json store to Struct
// 	var post Post
// 	json.Unmarshal(jsonData, &post) // store into struct
//
// 	// 把结构体封装为由字节切片组成的 JSON 数据
// 	res, err := json.MarshalIndent(&post, "", "\t\t") // Marshal 返回 v 的 JSON 编码
// 	if err != nil {
// 		fmt.Println("Failed to MarshalIndent error:", err)
// 		return
// 	}
// 	err = os.WriteFile("./gowp/ch7/json/post1.json", res, 0644)
// 	if err != nil {
// 		fmt.Println("Failed to os.WriteFile error:", err)
// 		return
// 	}
//
// 	// 3、use decoder
// 	// decoder := json.NewDecoder(jsonFile)
// 	// var post Post
// 	// for {
// 	// 	err := decoder.Decode(&post) // 将 JOSN 数据解码至结构体
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		fmt.Println("Failed to decode json:", err)
// 	// 		return
// 	// 	}
// 	// }
// 	// res, err := json.Marshal(post)
// 	// if err != nil {
// 	// 	fmt.Println("Failed to marshal json:", err)
// 	// 	return
// 	// }
// 	// var out bytes.Buffer
// 	// err = json.Indent(&out, res, "", "\t")
// 	// out.WriteTo(os.Stdout)
// }

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open json file", err)
		return
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Failed to decode json file", err)
		return
	}
	return
}

func unmarshal(filename string) (post Post, err error) {
	jsonfile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open json file", err)
		return
	}
	defer jsonfile.Close()
	jsonData, err := io.ReadAll(jsonfile)
	if err != nil {
		fmt.Println("Failed to read json data", err)
		return
	}
	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	_, err := decode("./gowp/ch7/json/post.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
