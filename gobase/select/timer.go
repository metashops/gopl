package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 3)
	select {
	case <-t.C:
		fmt.Println("1 s time", time.Now().Unix())
	case <-time.After(1 * time.Second):
		// 获取时间戳
		timestamp := time.Now().Unix()
		// //时间戳转Time 再转 string
		timeNow := time.Unix(timestamp, 0)
		timeString := timeNow.Format("2006-01-02 15:04:05")
		fmt.Println("2 s time", timeString)
	case <-time.After(2 * time.Second):
		// string 转 时间戳
		stringTime := "2017-08-30 16:40:41"
		loc, _ := time.LoadLocation("Local")
		the_time, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
		if err != nil {
			unix_time := the_time.Unix()
			fmt.Println("unix_time:", unix_time)
		}
	}
}
