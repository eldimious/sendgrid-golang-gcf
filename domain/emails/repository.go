package email

// EmailRepository provides an abstraction on top of the email data source
type EmailRepository interface {
	SendEmail(*From, *To, *Message) error
}
