package main

func searchByCityOrState(addressBook map[string]*Contact, city string, state string) []*Contact {
	filteredContacts := filterContacts(addressBook, func(c *Contact) bool {
		return c.City == city || c.State == state
	})
	return filteredContacts
}
