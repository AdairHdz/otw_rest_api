package utility

import (
	"net/smtp"
	"os"
)

func SendToEmail(email string, verificationCode string) (err error) {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	to := []string{
		email,
	}
	
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	
	message := []byte("Subject: Codigo de verificacion\r\n" +
	"\r\n" +
	"Su codigo de verificacion es: "+ verificationCode)
	
	auth := smtp.PlainAuth("", from, password, smtpHost)
	
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	
	return
}

