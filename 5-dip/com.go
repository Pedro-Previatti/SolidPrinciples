package dip

type MessageSender interface {
	Send(to string, message string)
}

type EmailServiceDIP struct{}

func (e *EmailServiceDIP) Send(to string, message string) {
	// Send email
}

type SMSService struct{}

func (s *SMSService) Send(to string, message string) {
	// Send SMS
}

type NotificationServiceDIP struct {
	messageSender MessageSender
}

func (n *NotificationServiceDIP) Notify(to string, message string) {
	n.messageSender.Send(to, message)
}
