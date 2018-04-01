package binders

import (
    "github.com/gin-gonic/gin"
    "time"
    . "untitled3/constants"
    "untitled3/model"
    "database/sql"
    "net/http"
)

func Metric(c *gin.Context) {
    b := struct {
        DeviceID  int       `json:"device_id" binding:"required"`
        Metric1   int       `json:"metric_1"`
        Metric2   int       `json:"metric_2"`
        Metric3   int       `json:"metric_3"`
        Metric4   int       `json:"metric_4"`
        Metric5   int       `json:"metric_5"`
        LocalTime int64     `json:"local_time" binding:"required"`
    }{}
    err := c.Bind(&b)

    if err != nil {
        c.Status(http.StatusBadRequest)
        c.Abort()
        return
    }

    dm := model.DeviceMetric{
        DeviceID: b.DeviceID,
        Metric1: sql.NullInt64{Int64: int64(b.Metric1)},
        Metric2: sql.NullInt64{Int64: int64(b.Metric2)},
        Metric3: sql.NullInt64{Int64: int64(b.Metric3)},
        Metric4: sql.NullInt64{Int64: int64(b.Metric4)},
        Metric5: sql.NullInt64{Int64: int64(b.Metric5)},
        LocalTime: time.Unix(b.LocalTime, 0),
    }
    c.Set(BindingDeviceMetric, &dm)
}