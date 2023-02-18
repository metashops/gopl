package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

const COUNTER = 15

func main() {
	ctx := context.Background()
	// 初始化 Redis 客户端
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 文章ID
	articleID := "1001"

	err := ArticleCount(ctx, articleID)
	if err != nil {
		panic(err)
	}

	viewCount, err := GetArticleCount(ctx, articleID)
	if err != nil {
		panic(err)
	}

	fmt.Println("文章ID为:", articleID, "的阅读量是 ", viewCount)

}

// ArticleCount 增加阅读量计数器
func ArticleCount(ctx context.Context, articleID string) error {
	count, err := GetArticleCount(ctx, articleID)
	if err != nil {
		panic(err)
	}
	if count > COUNTER {
		return nil
	} else {
		err1 := rdb.Incr(ctx, articleID).Err()
		if err1 != nil {
			panic(err1)
		}

	}
	return nil
}

// GetArticleCount 获取阅读量计数器的值
func GetArticleCount(ctx context.Context, articleID string) (int, error) {
	viewCount, err := rdb.Get(ctx, articleID).Int()
	if err != nil {
		panic(err)
	}
	if viewCount > COUNTER {
		return COUNTER, nil
	} else {
		return viewCount, nil
	}
}
