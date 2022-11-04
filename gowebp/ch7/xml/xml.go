package main

import (
	`encoding/xml`
	`fmt`
	`io`
	`log`
	`os`
)

type Post struct { // definition a Struct, expression data
	XMLName  string    `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	// 1、打开文件
	xmlFile, err := os.Open("./gowp/ch7/post.xml")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer xmlFile.Close()
	// 2、读取文件,读取到xmlData中
	// xmlDate, err := io.ReadAll(xmlFile)
	// if err != nil {
	// 	fmt.Println("Failed to reading xml data:", err)
	// 	return
	// }
	// var post Post
	// 3、将 xmlData存储到结构体post中（解封XML数据）
	// GO语言"encoding/json"包中"Unmarshal"函数
	// Unmarshal 解析 JSON-encoded 数据并将结果存储在 v 指向的值中。
	// 如果 v 为 nil 或不是指针，Unmarshal 返回一个 InvalidUnmarshalError
	// Unmarshal 不适合体积较大文件，无法处理流方式传输（使用Decoder结构来代替Unmarshal）
	// xml.Unmarshal(xmlDate, &post)
	// fmt.Println("post:", post)

	// 3、使用Decoder结构来代替Unmarshal
	decoder := xml.NewDecoder(xmlFile)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed to decodeing XML into token:", err)
			return
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
			}
		}
	}

}
