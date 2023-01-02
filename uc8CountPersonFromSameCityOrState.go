package main

func filterContact(addressBook map[string]*Contact, filter func(*Contact) bool) map[string]int {
	filteredContacts := make(map[string]int)
	for _, contact := range addressBook {
		if filter(contact) {
			if contact.City != "" {
				filteredContacts[contact.City]++
			}
			if contact.State != "" {
				filteredContacts[contact.State]++
			}
		}
	}
	return filteredContacts
}
