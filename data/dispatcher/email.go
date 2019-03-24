package dispatcher

import (
	"github.com/eldimious/sendgrid-golang-gcf/config"
	domain "github.com/eldimious/sendgrid-golang-gcf/domain/emails"
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
func (dispatcher *Dispatcher) SendEmail(sender *domain.Sender, receiver *domain.Receiver, msg *domain.Message) error {
	from := mail.NewEmail(sender.Name, sender.Email)
	subject := msg.Subject
	to := mail.NewEmail(receiver.Name, receiver.Email)
	plainTextContent := msg.PlainTextContent
	htmlContent := msg.HtmlContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(dispatcher.config.APIKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
