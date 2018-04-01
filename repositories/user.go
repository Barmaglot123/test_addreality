package repositories

import (
    "github.com/jinzhu/gorm"
    "untitled3/model"
)

type User interface {
    Insert(tx Transaction, metric *model.DeviceMetric)
    First(tx Transaction, ID int) *model.User
}

type user struct {}

func NewUser() Metric {
    return &metric {}
}

func (u user) Insert(tx Transaction, user *model.User) {
    tx.DataSource().(*gorm.DB).Create(user)
}

func (u user) First(tx Transaction, ID int) *model.User {
    m := model.User{
        ID: ID,
    }

    tx.DataSource().(*gorm.DB).First(&m)

    return &m
}