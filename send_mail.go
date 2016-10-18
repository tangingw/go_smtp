package main

import (
	"net/smtp"
	"fmt"
	"strings"
)

/**
	Modified from https://gist.github.com/jpillora/cb46d183eca0710d909a
	Thank you very much.
**/

const (
	SMTP_SERVER = "smtp.gmail.com"
)


type Sender struct {

	User		string
	Password	string
}


func NewSender(Username, Password string) Sender {

	return Sender{Username, Password}
}


func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) {

	msg := 	"From: " + sender.User + "\n" +
	       	"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTP_SERVER + ":587", 
		smtp.PlainAuth("", sender.User, sender.Password, SMTP_SERVER),
		sender.User, Dest, []byte(msg))	
	
	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Mail sent successfully!")
}
