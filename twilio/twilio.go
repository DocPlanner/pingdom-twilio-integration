package twilio

import (
	"github.com/gin-gonic/gin"
	"github.com/subosito/twilio"
	"go.uber.org/config"
)

type twilioSettings struct {
	PhoneFrom  string `yaml:"phone_from"`
	AccountSid string `yaml:"account_sid"`
	AuthToken  string `yaml:"auth_token"`
}

func TwilioMiddleware(yaml *config.YAML) gin.HandlerFunc {
	var twilioSettings twilioSettings
	err := yaml.Get("twilio").Populate(&twilioSettings)
	if err != nil {
		panic(err)
	}

	twilioClient := twilio.NewTwilio(twilioSettings.AccountSid, twilioSettings.AuthToken)

	return func(context *gin.Context) {
		context.Set("twilio_client", twilioClient)
		context.Set("twilio_phone_form", twilioSettings.PhoneFrom)

		context.Next()
	}
}
