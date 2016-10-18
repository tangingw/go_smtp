package main

func main() {

	sender := NewSender("<YOUR EMAIL ADDRESS>", "<YOUR EMAIL PASSWORD>")

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"abc@gmail.com", "xyz@gmail.com", "larrypage@googlemail.com"} 

	Subject := "Testing email from golang"
	bodyMessage := "Sending email using Golang. Yeah"

	sender.SendMail(Dest, Subject, bodyMessage)
}
