package repositories

import (
    "github.com/jinzhu/gorm"
    "untitled3/model"
)

type DeviceAlert interface {
    Insert(tx Transaction, alert *model.DeviceAlert)
}

type deviceAlert struct {}

func NewDeviceAlert() DeviceAlert {
    return &deviceAlert {}
}

func (a deviceAlert) Insert(tx Transaction, alert *model.DeviceAlert) {
    tx.DataSource().(*gorm.DB).Create(alert)
}