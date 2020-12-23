package sms

type SMS interface {
	SendSMS(number string, message string) error
}

func ValidateNumber(number string) error {
	return nil
}
