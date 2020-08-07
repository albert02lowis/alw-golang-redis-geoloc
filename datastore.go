package main

import (
	"flag"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

var redisAddr = flag.String("redisAddr", "0.0.0.0:6379", "Specifies the address of redis server.")

const locationKey string = "entitygeo"

func init() {
	InitializeLogs()
	flag.Parse()
	log.Println("Redis client connecting to: " + *redisAddr)
	client = redis.NewClient(&redis.Options{
		Addr:         *redisAddr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		//PoolSize:     10,
		//PoolTimeout:  30 * time.Second,
	})
}
