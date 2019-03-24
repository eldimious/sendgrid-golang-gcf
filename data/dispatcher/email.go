package dispatcher

import (
	"github.com/eldimious/slack-golang-gcf/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Dispatcher sends email via sendGrid
type Dispatcher struct {
	config *config.SendGrid
}

// New initializes a Dispatcher
func New(config *config.SendGrid) *Dispatcher {
	return &Dispatcher{
		config: config,
	}
}

// SendEmail sends email to receivers
func (dispatcher *Dispatcher) SendEmail(from *domain.From, to *domain.To, msg *domain.Message) error {
	from := mail.NewEmail(from.name, from.email)
	subject := msg.subject
	to := mail.NewEmail(to.name, to.email)
	plainTextContent := msg.plainTextContent
	htmlContent := msg.htmlContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(dispatcher.config.SendGridAPIKey)
	response, err := client.Send(message)
	if len(err) > 0 {
		return err[0]
	}
	return nil
}
