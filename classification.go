package main

import "fmt"

type Service string

const (
	Payment        Service = "payment"
	Authentication Service = "authentication"
	Storage        Service = "storage"
	UnknownService Service = "unknown"
)

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Category string

const (
	Incident        Category = "incident"
	Request         Category = "request"
	Question        Category = "question"
	UnknownCategory Category = "unknown"
)

type Classification struct {
	Service  `json:"service"`
	Priority `json:"priority"`
	Category `json:"category"`
}

func (c Classification) Validate() error {
	switch c.Service {
	case "payment", "authentication", "storage", "unknown":
	default:
		return fmt.Errorf("Not valid service: %s", c.Service)
	}

	switch c.Priority {
	case "low", "medium", "high":
	default:
		return fmt.Errorf("Not valid priority: %s", c.Priority)
	}

	switch c.Category {
	case "incident", "request", "question", "unknown":
	default:
		return fmt.Errorf("Not valid category: %s", c.Category)
	}

	return nil
}
