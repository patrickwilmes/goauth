/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

package common

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
)

const (
	mailSmtp     = "mail.smtp"
	mailFrom     = "mail.from"
	mailPassword = "mail.password"
	mailPort     = "mail.port"
)

type Message struct {
	Token    string
	Subject  string
	Receiver string
}

func renderTemplateForToken(message Message) (string, error) {
	tmpl, err := template.ParseFiles("config/verification-mail.html")
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, message); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SentMessage(message Message) error {
	body, err := renderTemplateForToken(message)
	if err != nil {
		return err
	}
	fromAddress := GetStringValue(mailFrom)
	smtpPort := GetIntValue(mailPort)
	mail := gomail.NewMessage()
	mail.SetHeader("From", fromAddress)
	mail.SetHeader("To", message.Receiver)
	mail.SetHeader("Subject", message.Subject)
	mail.SetBody("text/html", body)

	d := gomail.NewDialer(GetStringValue(mailSmtp), smtpPort, fromAddress, GetStringValue(mailPassword))
	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
