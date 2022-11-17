package main

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:+573105103968")
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello from Golang!")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Te amo mucho mi amor - esto es lo que me falta programar!")
	}
}
