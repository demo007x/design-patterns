package main

type IOtp interface {
	genRandomOPT(int) string
	saveOPTCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOPT(optLength int) error {
	opt := o.iOtp.genRandomOPT(optLength)
	o.iOtp.saveOPTCache(opt)
	message := o.iOtp.getMessage(opt)
	if err := o.iOtp.sendNotification(message); err != nil {
		return err
	}
	return nil
}
