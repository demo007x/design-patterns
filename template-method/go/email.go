package main

import (
	"fmt"
)

type Email struct {
	Otp
}

func (s *Email) genRandomOPT(len int) string {
	randomOTP := "2345"
	fmt.Printf("Email: generating random otp %s \n", randomOTP)
	return randomOTP
}

func (s *Email) saveOPTCache(otp string) {
	fmt.Printf("Email: saving otp %s to cache", otp)
}

func (s *Email) getMessage(otp string) string {
	return "Email otp for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("Email: sending email %s \n", message)
	return nil
}
