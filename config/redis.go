package config

import "github.com/spf13/viper"

var Redis redis

type RedisConfig interface {
    Address()  string
    Password() string
    Db()       int
}

type redis struct {
    address     string
    testArrdess string
    password    string
    db          int
}

func loadRedis() redis{
    s := viper.Sub("redis")
    return redis{
        address:     s.GetString("address"),
        testArrdess: s.GetString("test_address"),
        db:          s.GetInt("db"),
        password:    s.GetString("password"),
    }
}

func (r redis)Address() string{
    return r.address
}

func (r redis)Password() string{
    return r.password
}

func (r redis)Db() int{
    return r.db
}

func (r redis)TestAddress() string {
    return r.testArrdess
}