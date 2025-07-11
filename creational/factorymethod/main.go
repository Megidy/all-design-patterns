package main

import (
	"fmt"
	"log"
)

const (
	SmsNotifier   = "sms_notifier"
	EMailNotifier = "email_notifier"
)

func main() {
	notifier, err := NewNotifier(SmsNotifier)
	if err != nil {
		log.Fatal(err)
	}
	notifier.SendMessage("Hello")
}

func NewNotifier(notifier string) (Notifier, error) {
	switch notifier {
	case SmsNotifier:
		return NewSMSNotifier(), nil
	case EMailNotifier:
		return NewEmailNotifier(), nil
	default:
		return nil, fmt.Errorf("invalid notifier")
	}
}

type Notifier interface {
	SendMessage(string)
}

type SMSNotifier struct{}

func NewSMSNotifier() *SMSNotifier {
	return &SMSNotifier{}
}

func (s *SMSNotifier) SendMessage(string) {
	log.Printf("sent message from sms notifier")
}

type EmailNotifier struct{}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (s *EmailNotifier) SendMessage(string) {
	log.Printf("sent message from email notifier")
}
