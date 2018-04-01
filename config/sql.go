package config

import (
    "github.com/spf13/viper"
)

type SqlConfig interface {
    Host()        string
    Name()        string
    Password()    string
    User()        string
    SSL()         string
}

var Sql sql

type sql struct {
    host		string
    name		string
    user		string
    password	string
    ssl		    string
}

func loadSql() sql {
    s := viper.Sub("sql")
    return sql{
        host: 		s.GetString("host"),
        name: 		s.GetString("name"),
        password: 	s.GetString("password"),
        user: 		s.GetString("user"),
        ssl: 		s.GetString("ssl"),
    }
}


func(s sql) Host() string {
    return s.host
}

func(s sql) Name() string {
    return s.name
}

func(s sql) Password() string {
    return s.password
}

func(s sql) User() string {
    return s.user
}

func(s sql) SSL() string {
    return s.ssl
}