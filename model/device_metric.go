package model

import (
    "time"
    "database/sql"
)

type DeviceMetric struct {
    ID        int           `json:"id"`
    DeviceID  int           `json:"device_id"`
    Metric1   sql.NullInt64 `json:"metric1" gorm:"column:metric_1"`
    Metric2   sql.NullInt64 `json:"metric2" gorm:"column:metric_2"`
    Metric3   sql.NullInt64 `json:"metric3" gorm:"column:metric_3"`
    Metric4   sql.NullInt64 `json:"metric4" gorm:"column:metric_4"`
    Metric5   sql.NullInt64 `json:"metric5" gorm:"column:metric_5"`
    LocalTime time.Time     `json:"local_time"`
    CreatedAt time.Time     `json:"-" gorm:"column:server_time"`
}

