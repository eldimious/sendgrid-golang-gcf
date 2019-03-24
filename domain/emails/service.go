package email

// EmailService defines email service behavior.
type EmailService interface {
	SendEmail(*From, *To, *Message) error
}

// Service struct handles email business logic tasks.
type Service struct {
	repository EmailService
}

func (svc *Service) SendEmail(from *From, to *To, message *Message) error {
	return svc.repository.SendEmail(from, to, message)
}

// NewService creates a new service struct
func NewService(repository EmailRepository) *Service {
	return &Service{repository: repository}
}
