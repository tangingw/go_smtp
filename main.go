package main

import (
	"fmt"

	"github.com/darkfoxs96/go_smtp/gomail"
)

func main() {

	sender := gomail.NewSender("dark@yandex.com", "123456qwe", gomail.SMTP["yandex.com"].Host, gomail.SMTP["yandex.com"].Port)

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"dark@yandex.com"}

	Subject := "Testing HTLML Email from golang"
	message := `
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>This is the body<br>
	<div class="moz-signature"><i><br>
	<br>
	Regards<br>
	Alex<br>
	<i></div>
	</body>
	</html>
	`
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)

	err := sender.SendMail(Receiver, Subject, bodyMessage)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
