package factorymethod

func TestFactoryMethod() {
	// If we need to send an Email
	emailFactory := &EmailCreator{}
	SendNotification(emailFactory, "Welcome to the platform!")

	// If we need to send an SMS
	smsFactory := &SMSCreator{}
	SendNotification(smsFactory, "Your OTP is 1234")
}
