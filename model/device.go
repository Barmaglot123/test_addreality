package model

type Device struct {
    ID     int   `json:"id"`
    Name   string `json:"name"`
    UserID uint   `json:"user_id"`
    User   User   `json:"setting" gorm:"ForeignKey:UserID"`
    Metrics []DeviceMetric
}
