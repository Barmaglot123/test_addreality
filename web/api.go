package web

import (
    "github.com/gin-gonic/gin"
    "untitled3/repositories"
    "untitled3/services"
    "untitled3/web/controllers"
    "untitled3/datasource"
    "untitled3/web/binders"
)

func Run() {
    r := gin.Default()

    metricCon := metricController()

    r.POST("/metric", binders.Metric, metricCon.Create)

    r.Run(":9669")
}

func NewServer() *gin.Engine {
    r := gin.Default()

    metricCon := metricController()

    r.POST("/metric", binders.Metric, metricCon.Create)

    return r
}

func metricController() controllers.Metric {
    factory := repositories.NewTransactionFactory(datasource.Sql)
    repo := repositories.NewMetric()
    deviceRepo := repositories.NewDevice()
    deviceAlertRepo := repositories.NewDeviceAlert()
    service := services.NewMetric(repo, deviceAlertRepo, deviceRepo, factory)
    return controllers.NewMetric(service)
}