package main

import (
    "untitled3/datasource"
    "untitled3/config"
    "untitled3/web"
)

func main() {
    config.Load()
    datasource.SetupSql()
    defer datasource.Sql.Close()
    datasource.SetupRedis()
    defer datasource.Redis.Close()
    web.Run()
}