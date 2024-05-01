package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/briannbig/afya-village/internal/domain/queue"
	"github.com/nats-io/nats.go"
	"gopkg.in/gomail.v2"
)

type Notifier interface {
	Notify(msg *nats.Msg)
	notify(recipient string, subject string, message string)
}

type emailNotifier struct {
	dialer *gomail.Dialer
}

func Emailnotifier() Notifier {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return nil
	}
	
	dialer := gomail.NewDialer(smtpHost, port, smtpUsername, smtpPassword)
	
	return emailNotifier{
		dialer: dialer,
	}
}

func (e emailNotifier) Notify(msg *nats.Msg) {
	var payload any

	err := json.Unmarshal(msg.Data, &payload)
	if err != nil {
		log.Printf("Error unmarshalling payload --- %s", err.Error())
		return
	}

	data := payload.(map[string]any)
	recipient := fmt.Sprintf("%s", data["email"])

	var subject, message string

	switch msg.Subject {
	case string(queue.EventUserCreated):
		username := data["username"]
		message = fmt.Sprintf("Hello %s, welcome to afya village", username)
		subject = "Welcome"
	}

	e.notify(recipient, subject, message)
}

func (e emailNotifier) notify(recipient string, subject string, message string) {
	m := gomail.NewMessage()

	m.SetHeader("From", "sender@example.com")
	m.SetHeader("To", recipient)

	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	if err := e.dialer.DialAndSend(m); err != nil {
		fmt.Println("Could not send email:", err)
		return
	}

}
