package dip

type EmailService struct{}

func (e *EmailService) Send(to string, message string) {
	// Send email
}

type NotificationService struct {
	emailService *EmailService
}

func (n *NotificationService) Notify(to string, message string) {
	n.emailService.Send(to, message)
}
