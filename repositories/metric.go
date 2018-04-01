package repositories

import (
    "github.com/jinzhu/gorm"
    "untitled3/model"
)

type Metric interface {
    Insert(tx Transaction, metric *model.DeviceMetric)
}

type metric struct {}

func NewMetric() Metric {
    return &metric {}
}

func (m metric) Insert(tx Transaction, metric *model.DeviceMetric) {
    tx.DataSource().(*gorm.DB).Create(metric)
}