package web

import (
    "net/http"
    "testing"
    "untitled3/datasource"
    . "github.com/smartystreets/goconvey/convey"
    "encoding/json"
    . "untitled3/constants"
    "untitled3/config"
    "time"
    . "untitled3/dockertest"
    "untitled3/utils"
    "fmt"
)

func SetupDatasourcesForTest() {
    datasource.SetupTestSql()
    datasource.SetupTestRedis()
}

func TestMerticAdding (t *testing.T) {
    config.Load()
    client, c := StartContainer("marsianin/test_db_1",t)
    defer RemoveContainer(client, c.ID, t)

    WaitStarted(client, c.ID, 5*time.Second, t)
    WaitReachable(10*time.Second, t)

    SetupDatasourcesForTest()
    defer datasource.Sql.Close()

    s := NewServer()
    url := "http://localhost:9669/metric"

    Convey("Test method for adding a metric", t, func() {
        Convey("When all params are correct", func() {
            m := initMetric(1, 6, 6, 6, 6, 6, 1500000000)
            body, err := json.Marshal(m)
            So(err, ShouldBeNil)
            w := utils.SendPostRequest(url, body, s)
            So(w.Code, ShouldEqual, http.StatusOK)
        })

        Convey("When missed required params", func() {
            m := struct {}{}
            body, err := json.Marshal(m)
            So(err, ShouldBeNil)
            w := utils.SendPostRequest(url, body, s)
            So(w.Code, ShouldEqual, http.StatusBadRequest)
        })

        Convey("When metric value out of limits", func() {
            m := initMetric(1, 100, 6, 6, 6, 6, 1500000000)
            body, err := json.Marshal(m)
            So(err, ShouldBeNil)
            w := utils.SendPostRequest(url, body, s)
            So(w.Code, ShouldEqual, http.StatusOK)
            a, _ := datasource.Redis.Get("1" + RedisAlertKey)
            e := fmt.Sprintf(OutOfRangeAlert, "Metric1")
            So(a, ShouldEqual, e)
        })
    })
}

func initMetric(deviceID, m1, m2, m3, m4, m5, localTime int) interface{} {
    m := struct {
        DeviceID  int       `json:"device_id"`
        Metric1   int       `json:"metric_1"`
        Metric2   int       `json:"metric_2"`
        Metric3   int       `json:"metric_3"`
        Metric4   int       `json:"metric_4"`
        Metric5   int       `json:"metric_5"`
        LocalTime int       `json:"local_time"`
    }{
        DeviceID:  deviceID,
        Metric1:   m1,
        Metric2:   m2,
        Metric3:   m3,
        Metric4:   m4,
        Metric5:   m5,
        LocalTime: localTime,
    }
    return m
}