package db

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/zyfdegh/redis-bench/entity"

	redis "gopkg.in/redis.v5"
)

// Ping connects to redius
func Ping() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

// Bench set and get KV
func Bench() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	for i := 0; i < 10000; i++ {
		key := "k123"
		value := entity.RedisTest{
			Int:    rand.Int(),
			String: "abcd",
			Bool:   true,
			Array:  []int{4, 5, 6},
			Struct: entity.Struct{
				Name: "tom",
				Age:  20,
			},
		}
		data, err := json.Marshal(value)
		if err != nil {
			log.Printf("marshal error: %v\n", err)
			return
		}

		err = client.Set(key, data, 0).Err()
		if err != nil {
			log.Printf("set error: %v\n", err)
			return
		}

		result, err := client.Get(key).Bytes()
		if err != nil {
			log.Printf("get error: %v\n", err)
			return
		}
		redisTest2 := &entity.RedisTest{}
		err = json.Unmarshal(result, redisTest2)
		if err != nil {
			log.Printf("unmarshal error: %v", err)
			return
		}
	}
	return
}
