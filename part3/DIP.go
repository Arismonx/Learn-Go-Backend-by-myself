package main

import "fmt"

type Notifier interface {
	SendNotification(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) SendNotification(message string) {
	fmt.Println("Sending Email:", message)
}

type NotificationService struct {
	notifier Notifier
}

func (ns NotificationService) NotifyUser(message string) {
	ns.notifier.SendNotification(message)
}
