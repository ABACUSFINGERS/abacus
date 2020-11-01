

package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "abacusfingers@gmail.com", "TestMDP1234", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"abacusfingers@gmail.com",}
	msg := []byte("From: xyz@gmail.com\r\n" +"To: abc@gmail.com\r\n" +
		"Subject: Go language email send using smpt package\r\n" +
		"\r\n" +
		"Hi Team,\nThis is sample email using Go language executable file.\nIn this we will be learning how to send email using go lang.\n\nRegards,\nEducateSpace support\r\n")
	err := smtp.SendMail("smtp.gmail.com:25", auth, "abacusfingers@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
