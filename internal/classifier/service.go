package classifier

import (
	"log"
)

type Executor interface {
	Execute(message Classification) error
}

var _ Executor = (*JiraService)(nil)

type JiraService struct{}

func (js JiraService) Execute(message Classification) error {
	log.Println("Raised a new ticket in Jira")
	return nil
}

type SupportEmailService struct{}

func (js SupportEmailService) Execute(message Classification) error {
	log.Println("Sent an email to Production Support Team")
	return nil
}
