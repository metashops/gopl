package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 清空排行榜
	client.Del(ctx, "leaderboard")

	// 生成一些随机得分
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 10; i++ {
		score := rand.Intn(100)
		// 将得分添加到有序集合中
		client.ZAdd(ctx, "leaderboard", redis.Z{
			Score:  float64(score),
			Member: "player" + strconv.Itoa(i),
		})
	}

	// 获取排行榜前三名
	result, err := client.ZRevRangeWithScores(ctx, "leaderboard", 0, 2).Result()
	if err != nil {
		log.Fatal(err)
	}

	// 打印排行榜前三名
	fmt.Println("Leaderboard =>")
	for i, v := range result {
		fmt.Printf("%d. %s (score: %v)\n", i+1, v.Member, v.Score)
	}

	// 获取玩家 player3 的排名和得分
	rank, err := client.ZRevRank(ctx, "leaderboard", "player3").Result()
	if err != nil {
		log.Fatal(err)
	}
	score, err := client.ZScore(ctx, "leaderboard", "player3").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("player3: rank=%v, score=%v\n", rank+1, score)

	// 将玩家 player3 的得分加 10 分
	client.ZIncrBy(ctx, "leaderboard", 10, "player3")

	// 获取新的排行榜前三名
	result, err = client.ZRevRangeWithScores(ctx, "leaderboard", 0, 2).Result()
	if err != nil {
		log.Fatal(err)
	}

	// 打印新的排行榜前三名
	fmt.Println("Updated leaderboard:")
	for i, v := range result {
		fmt.Printf("%d. %s (score: %v)\n", i+1, v.Member, v.Score)
	}
}
