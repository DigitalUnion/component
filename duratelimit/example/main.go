package main

import (
	"context"
	"fmt"
	"git.du.com/cloud/du_component/duratelimit"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	// 基本使用
	//baseUse()
	// 在 gin 中使用中间件
	//ginMiddleware()
	//// gin 中间件使用多个限流器
	//multiMiddleware()
	c := redis.NewClient(&redis.Options{
		Addr:         "r-2zeb410hhylkwsko48.redis.rds.aliyuncs.com:6379",
		Password:     "v2wyewfBJ0bD22GZvmrule",
		PoolSize:     100,
		MinIdleConns: 2,
		//MaxConnAge:   30 * time.Minute,
	})
	// 172.17.146.221
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c.Get(context.Background(), "1")
				time.Sleep(time.Millisecond * 10)
			}
		}()
	}
	fmt.Println(c)
	select {}
}

func baseUse() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	limit := duratelimit.NewDuRateLimit(rdb, duratelimit.Resource{Name: "test", RatePerMinute: 20, Allow: 19})
	var success, fail int
	go func() {
		for {
			if limit.Allow() {
				success++
				//fmt.Println("正常处理请求")
			} else {
				fail++
				//fmt.Println("请求被限流")
			}
			time.Sleep(time.Millisecond * 2)
		}
	}()
	time.Sleep(time.Second * 500)
	fmt.Printf("正常处理请求: %d, 请求被限流: %d", success, fail)
}

func ginMiddleware() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	limit := duratelimit.NewDuRateLimit(rdb, duratelimit.Resource{Name: "test", RatePerSecond: 10, Allow: 1})
	e := gin.Default()
	e.Use(duratelimit.RateLimit(limit))
}

func multiMiddleware() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	limitName1 := "test1"
	limitName2 := "test2"
	limitName3 := "test3"
	limit1 := duratelimit.NewDuRateLimit(rdb, duratelimit.Resource{Name: limitName1, RatePerSecond: 10, Allow: 1})
	limit2 := duratelimit.NewDuRateLimit(rdb, duratelimit.Resource{Name: limitName2, RatePerSecond: 10, Allow: 1})
	limit3 := duratelimit.NewDuRateLimit(rdb, duratelimit.Resource{Name: limitName3, RatePerSecond: 10, Allow: 1})
	e := gin.Default()
	e.Use(duratelimit.MultiRateLimit("client_id", map[string]*duratelimit.DuRateLimit{limitName1: limit1, limitName2: limit2, limitName3: limit3}))
}
