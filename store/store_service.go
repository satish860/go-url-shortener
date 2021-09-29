package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)


type StorageService struct{
	redis *redis.Client
}

var (
	storeService = &StorageService{}
)

func InitializeStore() *StorageService{
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong,err := redisClient.Ping().Result()
	if err !=nil{
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redis=redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string){ 
	err := storeService.redis.Set(shortUrl,originalUrl,time.Duration(6*time.Hour)).Err()
	if err!=nil{
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redis.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}