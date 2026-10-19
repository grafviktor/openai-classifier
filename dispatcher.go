package main

import "log"

type Dispatcher struct{}

func NewDispatcher() Dispatcher {
	return Dispatcher{}
}

func (d Dispatcher) Dispatch(message Message) error {
	switch message.classification.Service {
	case Payment:
		log.Printf("Send message to Payment Support: %q with priority %q", message.Text, message.classification.Priority)
	case Authentication:
		log.Printf("Send message to Account Administrators: %q with priority %q", message.Text, message.classification.Priority)
	case Storage:
		log.Printf("Send message to DBA: %q with priority %q", message.Text, message.classification.Priority)
	default:
		log.Printf("Send message to Help Desk: %q with priority %q", message.Text, message.classification.Priority)
	}
	return nil
}
