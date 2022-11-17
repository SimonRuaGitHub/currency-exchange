package notification

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendNotification(message string, fullPhoneNum string) {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:" + fullPhoneNum)
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Message was successfully sent!")
	}
}
