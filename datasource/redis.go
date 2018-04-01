package datasource

import (
    "github.com/go-redis/redis"
    "log"
    "reflect"
    "untitled3/config"
)

var Redis redisClient

type redisClient struct {
    client *redis.Client
}

const redisConfigName string = "redis"

func SetupRedis() {
    connectToRedis()
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

func(r redisClient) Set(key string, value interface{}) error {
    err := r.client.Set(key, value, 0).Err()
    if err != nil { return err }
    return nil
}

func(r redisClient) Get(s interface{}) error {
    val := reflect.ValueOf(s)
    t := reflect.Indirect(val).Type()

    for i := 0; i < t.NumField(); i++ {
        key := t.Field(i).Tag.Get(redisConfigName)
        if val.Elem().Field(i).Type().Name() == "string"{
            d, _ := r.client.Get(key).Result()
            reflect.Indirect(val.Elem().Field(i)).SetString(d)
        } else {
            d, _ := r.client.Get(key).Int64()
            reflect.Indirect(val.Elem().Field(i)).SetInt(d)
        }
    }
    return nil
}

func(r redisClient)Close(){
    r.client.Close()
}