package repositories

import (
    "github.com/jinzhu/gorm"
    "untitled3/model"
)

type Device interface {
    First(tx Transaction, id int) *model.Device
}

type device struct {}

func NewDevice() Device {
    return &device {}
}

func (d device) First(tx Transaction, id int) *model.Device {
    m := model.Device{ID: id}
    tx.DataSource().(*gorm.DB).Preload("User").First(&m)

    return &m
}