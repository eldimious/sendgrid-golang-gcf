package email

// EmailRepository provides an abstraction on top of the dispatcher data source
type EmailRepository interface {
	SendEmail(*Sender, *Receiver, *Message) error
}
