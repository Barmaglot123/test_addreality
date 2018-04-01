package controllers

import (
    "github.com/gin-gonic/gin"
    "untitled3/model"
    . "untitled3/constants"
    "untitled3/services"
)

type Metric interface {
    Create(c *gin.Context)
}

type metric struct {
    service services.Metric
}

func NewMetric (service services.Metric) Metric {
    return &metric{ service: service }
}

func (m metric) Create(c *gin.Context) {
    metric := c.MustGet(BindingDeviceMetric).(*model.DeviceMetric)
    m.service.Create(metric)
}