package main

import (
	"context"
	"fmt"
	"git.du.com/cloud/du_component/dubreaker"
	"github.com/go-redis/redis/v8"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	// 慢请求策略配置
	slowReqRule := &dubreaker.BreakerRule{
		Strategy:         0,
		RetryTimeoutMs:   10000,
		MinRequestAmount: 1,
		StatIntervalMs:   100,
		Threshold:        0.1,
		MaxAllowedRtMs:   1,
	}
	// 初始化一个熔断器
	breaker, err := dubreaker.InitBreaker("./", "example", slowReqRule, nil, false)
	if err != nil {
		panic(err)
	}
	key := "22"
	fn := func() (interface{}, error) {
		return redisGet(key)
	}
	for i := 0; i < 10; i++ {
		// 保护一个使用 redis 有返回值的函数
		ret, err := breaker.Entry(fn, "", redis.Nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(ret)
	}

	value := "11"
	fn1 := func() error {
		return redisSet(key, value)
	}
	for i := 0; i < 10; i++ {
		// 保护一个使用 redis 无返回值的函数
		err := breaker.EntryWithoutResult(fn1, redis.Nil)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func redisGet(key string) (interface{}, error) {
	time.Sleep(time.Millisecond * 100)
	result, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func redisSet(key, value string) error {
	time.Sleep(time.Millisecond * 100)
	_, err := client.Set(context.Background(), key, value, 0).Result()
	if err != nil {
		return err
	}
	return nil
}
