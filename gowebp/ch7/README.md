分析JSON步骤和XML基本类型：
分析程序首先要做的是把JSON的分析结果存储到一些结构里面，然后通过反问这些结构来提取数据。
步骤：
* 创建一些用于包含JSON数据的结构
* 通过json.Unmarshal 函数，把 JSON 数据解封到结构体里面

本次所用到知识点：
* Marshal：
  * func Marshal(v any)([]byte, error)
  * Marshal 返回 v 的 JSON 编码
  * 

* Unmarshal：
  * Marshal 返回 v 的 JSON 编码。
  * Unmarshal 解析 JSON-encoded 数据并将结果存储在v指向的值中。
  * 如果 v 为 nil 或不是指针，Unmarshal 返回一个 InvalidUnmarshalError