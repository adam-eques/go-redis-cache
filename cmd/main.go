package main

import (
	"fmt"

	"github.com/acentior/go-redis-cache/cache"
)

func main() {
	cache := cache.NewCache("127.0.0.1:6379", "default", "123456", 0)
	cache.Add("con1", "har1", 1000)
	cache.Add("con1", "har2", 1000)
	cache.Add("con1", "har3", 1000)
	cache.Add("con1", "har4", 1000)
	cache.Add("con2", "har1", 1000)
	cache.Add("con2", "har2", 1000)
	err := cache.All()
	fmt.Println(err)
}
