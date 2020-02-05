package main

import (
	"DocPlanner/pingdom-twilio-integration/auth"
	"DocPlanner/pingdom-twilio-integration/contacts"
	"DocPlanner/pingdom-twilio-integration/twilio"
	"flag"
	"github.com/gin-gonic/gin"
	"go.uber.org/config"
	"net/http"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "config file full path")
	flag.Parse()

	yaml, err := config.NewYAML(config.File(*configFile))
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(auth.BananaAuthMiddleware(yaml))
	r.Use(contacts.ContactsMapMiddleware(yaml))
	r.Use(twilio.TwilioMiddleware(yaml))

	r.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
		return
	})

	r.POST("/", pingdomHandler)

	_ = r.Run(":80")
}
