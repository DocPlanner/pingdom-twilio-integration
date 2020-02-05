package main

import (
	"DocPlanner/pingdom-twilio-integration/contacts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/subosito/twilio"
	"net/http"
)

type pingdomPayload struct {
	CheckId             int      `json:"check_id"`
	CheckName           string   `json:"check_name"`
	CheckType           string   `json:"check_type"`
	Tags                []string `json:"tags"`
	PreviousState       string   `json:"previous_state"`
	CurrentState        string   `json:"current_state"`
	StateChangedUtcTime string   `json:"state_changed_utc_time"`
	Description         string   `json:"description"`
}

func pingdomHandler(c *gin.Context) {
	var contactGroupsMap contacts.ContactsMap
	contactGroupsMap = *c.MustGet("contacts_map").(*contacts.ContactsMap)

	var twilioClient *twilio.Twilio
	twilioClient = c.MustGet("twilio_client").(*twilio.Twilio)

	fromNumber := c.MustGet("twilio_phone_form").(string)

	var pingdomPayload pingdomPayload
	if err := c.ShouldBindJSON(&pingdomPayload); err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactGroupName := c.Query("contact_group")
	if len(contactGroupName) == 0 || len(contactGroupsMap[contactGroupName]) == 0 {
		err := fmt.Errorf("contact_group empty or parameter not set!")
		_ = c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var errors []error
	for _, contact := range contactGroupsMap[contactGroupName] {
		_, err := twilioClient.SimpleSendSMS(fromNumber, contact, pingdomPayload.Description)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors})
	}

	c.Status(http.StatusOK)
}
