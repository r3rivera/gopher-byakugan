package main

import "fmt"

import "gopher-byakugan/sms"

func sampleHello() string {

	actual := "Hello Again!"
	fmt.Println("Message is ", actual)

	return actual
}

func sendMessage() string {
	actual := sms.SendSMS()
	return actual
}
