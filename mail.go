package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendMail(from, authCode string, to []string, subject, body string) error {
	contentType := "Content-Type: text/plain" + "; charset=UTF-8"
	auth := smtp.PlainAuth("", from, authCode, "smtp.qq.com")
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n%s\r\n\r\n %s",
		from, strings.Join(to, ","), subject, contentType, body)

	// fmt.Println(msg)

	err := smtp.SendMail("smtp.qq.com:587", auth, from, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

// func SendGoMail(from, auth string, to []string, subject, body string) error {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", from)
// 	m.SetHeader("To", to...)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", body)

// 	d := gomail.NewDialer("smtp.qq.com", 587, from, auth)

// 	if err := d.DialAndSend(m); err != nil {
// 		return err
// 	}
// 	return nil
// }
