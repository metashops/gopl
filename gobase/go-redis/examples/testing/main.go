package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	// 连接 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Redis无密码，可忽略
		DB:       0,  // 使用默认的DB
	})

	// 用户ID
	userId := 123

	// 获取当天日期字符串，格式为：20230216
	now := time.Now()
	today := now.Unix()

	// 设置用户在当天的签到状态
	rdb.SetBit(ctx, "user:"+strconv.Itoa(userId)+":signin", today, 1)

	// 获取用户的签到状态
	signinStatus, err := rdb.Get(ctx, "user:"+strconv.Itoa(userId)+":signin").Result()
	if err != nil {
		panic(err)
	}

	// 统计用户的签到天数
	signinCount := 0
	for i := 0; i < 365; i++ {
		if rdb.GetBit(ctx, "user:"+strconv.Itoa(userId)+":signin", int64(i)).Val() > 0 {
			signinCount++
		}
	}

	fmt.Printf("用户 %d 当天签到状态：%v，已连续签到 %d 天\n", userId, signinStatus, signinCount)
}
