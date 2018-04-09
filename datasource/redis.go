package datasource

import (
    "github.com/go-redis/redis"
    "log"
    "untitled3/config"
)

var Redis redisClient

type redisClient struct {
    client *redis.Client
}

func SetupRedis() {
    connectToRedis()
}

func SetupTestRedis() {
    connectToTestRedis()
}

func connectToRedis() {
    Redis.client = redis.NewClient(&redis.Options{
        Addr:     config.Redis.Address(),
        Password: config.Redis.Password(),
        DB:       config.Redis.Db(),
    })
    _, err := Redis.client.Ping().Result()

    if err != nil {
        log.Fatal(err)
    }
}

func connectToTestRedis() {
    Redis.client = redis.NewClient(&redis.Options{
        Addr:     config.Redis.TestAddress(),
        Password: config.Redis.Password(),
        DB:       config.Redis.Db(),
    })
    _, err := Redis.client.Ping().Result()

    if err != nil {
        log.Fatal(err)
    }
}

func(r redisClient) Set(key string, value interface{}) error {
    err := r.client.Set(key, value, 0).Err()
    if err != nil { return err }
    return nil
}

func(r redisClient) Get(key string) (string, error) {
    v , err := r.client.Get(key).Result()
    if err != nil { return "", err }
    return v, err
}

func(r redisClient)Close(){
    r.client.Close()
}