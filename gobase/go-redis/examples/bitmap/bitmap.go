package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	rdb "go-redis/examples/config"
)

type UserSign struct {
}

func main() {
	var user UserSign
	userID := 123
	err := user.DoSign(userID)
	if err != nil {
		panic(err)
	}
	count, err := user.GetSignCount(userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("count: ", count)

	exist, err := user.CheckSign(userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("check: ", exist)

}

// 用户签到
func (s *UserSign) DoSign(uid int) error {
	// 计算offect
	var offset = 4 // time.Now().Local().Day() - 1
	var keys = s.buildSignKey(uid)
	resid := rdb.GetRedis()
	_, err := resid.SetBit(context.Background(), keys, int64(offset), 1).Result()
	if err != nil {
		return err
	}
	defer resid.Close()
	return nil
}

// 判断用户是都已经签到了
func (s UserSign) CheckSign(uid int) (int64, error) {
	var keys string = s.buildSignKey(uid)
	var offset int = time.Now().Local().Day() - 1
	redisclinet := rdb.GetRedis()
	defer redisclinet.Close()
	return redisclinet.GetBit(context.Background(), keys, int64(offset)).Result()
}

// 获取用户签到的次数
func (s UserSign) GetSignCount(uid int) (int64, error) {
	var keys string = s.buildSignKey(uid)
	redisclinet := rdb.GetRedis()
	defer redisclinet.Close()
	count := redis.BitCount{Start: 0, End: 31}
	return redisclinet.BitCount(context.Background(), keys, &count).Result()
}

// 获取用户首次签到的日期
func (s UserSign) GetFirstSignDate(uid int) (string, error) {
	var keys string = s.buildSignKey(uid)
	redisclinet := rdb.GetRedis()
	defer redisclinet.Close()
	pos, err := redisclinet.BitPos(context.Background(), keys, 1).Result() // 获取第一位为1 的位置
	if err != nil {
		return "", err
	}
	pos = pos + 1

	var day int = time.Now().Local().Day()

	var offsetDay int = (day - int(pos)) * -1

	return time.Now().AddDate(0, 0, offsetDay).Format("2006-01-02"), nil
}

// 获取当月签到情况
// 根据需要自己实现返回
func (s UserSign) GetSignInfo(uid int) (interface{}, error) {
	var keys string = s.buildSignKey(uid)
	redisclinet := rdb.GetRedis()
	defer redisclinet.Close()
	var day int = time.Now().Local().Day()
	var dddd string = fmt.Sprintf("u%d", day)
	st, _ := redisclinet.Do(context.Background(), "BITFIELD", keys, "GET", dddd, 0).Result()
	f := st.([]interface{})
	var res []bool = make([]bool, 0)
	var days []string = make([]string, 0)
	var v int64 = f[0].(int64)
	fmt.Println(v)
	for i := day; i > 0; i-- {
		var pos int = (day - i) * -1
		var keys = time.Now().Local().AddDate(0, 0, pos).Format("2006-01-02")
		days = append(days, keys)
		var value = v>>1<<1 != v
		res = append(res, value)
		v >>= 1
	}
	fmt.Println(res)
	fmt.Println(days)
	return nil, nil
}

func (s UserSign) buildSignKey(uid int) string {
	var nowDate string = s.formatDate()
	return fmt.Sprintf("u:sign:%d:%s", uid, nowDate)
}

// 获取当前的日期
func (s UserSign) formatDate() string {
	return time.Now().Format("2006-01")
}
