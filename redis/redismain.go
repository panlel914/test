package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "47.100.245.167:6379",
		Password: "feg@125800", // no password set
		DB:       0,  // use default DB
		PoolSize: 1,
		PoolTimeout: time.Second * 10,
	})

	for i := 0;i<10;i++{
		go func(i int){
			fmt.Println(i)
			d:=redisdb.Get("aaa")
			fmt.Println(d.Result())
		}(i)
	}

	time.Sleep(time.Second*5)
}
