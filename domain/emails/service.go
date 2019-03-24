package email

// EmailService defines email service behavior.
type EmailService interface {
	SendEmail(*Message) error
}

// Service struct handles email business logic tasks.
type Service struct {
	repository EmailService
}

func (svc *Service) SendEmail(message *Message) error {
	return svc.repository.SendEmail(message)
}

// NewService creates a new service struct
func NewService(repository EmailRepository) *Service {
	return &Service{repository: repository}
}
