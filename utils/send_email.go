package utils

import (
    "net/smtp"
    . "untitled3/config"
    "strconv"
)

func SendEmail (to string, mailBody string) error {
    auth := smtp.PlainAuth("", Smtp.Username(), Smtp.Password(), Smtp.Host())
    err := smtp.SendMail(
        Smtp.Host() + ":" + strconv.Itoa(Smtp.Port()),
        auth,
        Smtp.Username(),
        []string{to},
        []byte(mailBody),
    )
    return err
}
