package main

func checkForDuplicateName(addressBook map[string]*Contact, contact *Contact) []*Contact {
	duplicateContacts := filterContacts(addressBook, func(c *Contact) bool {
		return c.FirstName == contact.FirstName && c.LastName == contact.LastName
	})
	return duplicateContacts

}

func filterContacts(addressBook map[string]*Contact, filter func(*Contact) bool) []*Contact {
	filteredContacts := []*Contact{}
	for _, contact := range addressBook {
		if filter(contact) {
			filteredContacts = append(filteredContacts, contact)
		}
	}
	return filteredContacts
}
