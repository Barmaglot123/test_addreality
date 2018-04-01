package services

import (
    "untitled3/model"
    "untitled3/repositories"
    "reflect"
    "fmt"
    "strings"
    "untitled3/datasource"
    "untitled3/config"
    . "untitled3/constants"
    "strconv"
    "untitled3/utils"
)

type Metric interface {
    Create(metric *model.DeviceMetric)
}

type metric struct {
    repo repositories.Metric
    deviceRepo repositories.Device
    deviceAlertRepo repositories.DeviceAlert
    txFactory repositories.TransactionFactory
}

func NewMetric(repo repositories.Metric, deviceAlertRepo repositories.DeviceAlert, deviceRepo repositories.Device, txFactory repositories.TransactionFactory) Metric {
    return &metric{
        repo: repo,
        deviceRepo: deviceRepo,
        deviceAlertRepo: deviceAlertRepo,
        txFactory: txFactory,
    }
}

func (m metric)Create(metric *model.DeviceMetric) {
    s := reflect.ValueOf(metric).Elem()
    typeOfT := s.Type()
    var sendMail bool
    var mailBody string = AlertMailBody

    for i := 0; i < s.NumField(); i++ {
        if !strings.HasPrefix(typeOfT.Field(i).Name, "Metric") {
            continue
        }

        fieldVal := s.Field(i).FieldByName("Int64").Int()
        fieldName := typeOfT.Field(i).Name

        if fieldVal == 0 {
            continue
        }

        s.Field(i).FieldByName("Valid").SetBool(true)

        if fieldVal > config.Metric.MaxLim() || fieldVal < config.Metric.MinLim(){
            sendMail = true
            mailBody = mailBody + fieldName + " "
            tx := m.txFactory.BeginNewTransaction()
            da := model.DeviceAlert{
                DeviceID: metric.DeviceID,
                Message: fmt.Sprintf(OutOfRangeAlert, fieldName),
            }

            m.deviceAlertRepo.Insert(tx, &da)
            tx.Commit()
            setAlertToRedis(metric.DeviceID, fieldName)
        }
    }

    if sendMail {
        go m.sendAlertMail(mailBody, metric.DeviceID)
    }

    tx := m.txFactory.BeginNewTransaction()
    defer tx.Commit()
    m.repo.Insert(tx, metric)
}

func setAlertToRedis(deviceID int, fieldName string){
    strID := strconv.FormatUint(uint64(deviceID), 10)
    alert := fmt.Sprintf(OutOfRangeAlert, fieldName)
    datasource.Redis.Set(strID + RedisAlertKey, alert)
}

func (m metric) sendAlertMail(b string, deviceID int) {
    tx := m.txFactory.BeginNewTransaction()
    defer tx.Commit()
    d := m.deviceRepo.First(tx, deviceID)
    utils.SendEmail(d.User.Email, b)
}