package twilio_app

import (
	"time"

	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type twilioRepo struct {
	callStatus   bool
	twilioClient *twilioClient
}

func NewTwilioRepo(twilioClient *twilioClient) *twilioRepo {
	return &twilioRepo{
		callStatus:   true,
		twilioClient: twilioClient,
	}
}

func (repo *twilioRepo) MakeAlertCall(callAnswerApiUrl, callFrom, callTo string) error {
	if repo.callStatus {
		repo.callStatus = false

		go func() {
			time.Sleep(5 * time.Minute)
			repo.callStatus = true
		}()

		params := new(api.CreateCallParams)
		params.SetUrl(callAnswerApiUrl)
		params.SetTo(callTo)
		params.SetFrom(callFrom)
		_, err := repo.twilioClient.client.Api.CreateCall(params)

		return err
	}

	return nil
}
