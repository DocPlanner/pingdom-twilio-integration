package contacts

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/config"
)

type contacts map[string]string
type contactGroups map[string][]string

type ContactsMap map[string][]string

func ContactsMapMiddleware(yaml *config.YAML) gin.HandlerFunc {
	var contacts contacts
	var contactGroups contactGroups

	err := yaml.Get("contacts").Populate(&contacts)
	if err != nil {
		panic(err)
	}

	err = yaml.Get("contact_groups").Populate(&contactGroups)
	if err != nil {
		panic(err)
	}

	contactsMap := buildContactsMap(contacts, contactGroups)

	return func(context *gin.Context) {
		context.Set("contacts_map", &contactsMap)

		context.Next()
	}
}

func buildContactsMap(contacts contacts, contactGroups contactGroups) ContactsMap {
	contactsMap := ContactsMap{}
	for groupName, contactNames := range contactGroups {
		var phoneNumbers []string
		for _, contactName := range contactNames {
			phone := contacts[contactName]
			phoneNumbers = append(phoneNumbers, phone)
		}
		contactsMap[groupName] = phoneNumbers
	}

	return contactsMap
}
